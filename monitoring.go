/*
	This file permit to handle the monitoring state to send on nagios engine
*/

package nagiosPlugin

import (
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"os"
)

// Nagios status
const (
	STATUS_UNKNOWN  = 3
	STATUS_CRITICAL = 2
	STATUS_WARNING  = 1
	STATUS_OK       = 0
)

// Monitoring data
type Monitoring struct {
	messages  []string
	status    int
	perfdatas modelPerfdata.Perfdatas
}

// NewMonitoring permit to create monitoring object
// The status is inizialized to Ok
func NewMonitoring() *Monitoring {
	monitoringData := &Monitoring{
		status:    STATUS_OK,
		perfdatas: make(modelPerfdata.Perfdatas, 0, 1),
		messages:  make([]string, 0, 1),
	}

	return monitoringData
}

// SetStatus permit to set new monitoring status if the current status is more critical that the previous status
func (m *Monitoring) SetStatus(status int) error {
	log.Debugf("Status: %d", status)
	if status > 3 || status < 0 {
		return errors.New("Status can't be greater than 3")
	}

	if status > m.status {
		log.Debugf("New monitoring status is %d", status)
		m.status = status
	}

	return nil
}

// Status permit to get the monitoring status
func (m *Monitoring) Status() int {
	return m.status
}

// AddMessage permit to add message to display in monitoring tools
func (m *Monitoring) AddMessage(message string) {
	log.Debugf("Message: %s", message)

	m.messages = append(m.messages, message)
}

// Messages permit to get all messages
func (m *Monitoring) Messages() []string {
	return m.messages
}

// Message permit to get message on given index
func (m *Monitoring) Message(index int) (string, error) {
	if index >= len(m.messages) {
		return "", errors.New("Index is out of list messages")
	}

	return m.messages[index], nil
}

// AddPerfdata permit to add perfdata to display in monitoring tools
func (m *Monitoring) AddPerfdata(label string, value int, unit string) error {
	log.Debugf("Label: %s, Value: %d, Unit: %s", label, value, unit)

	perfdata, err := modelPerfdata.NewPerfdata(label, value, unit)
	if err != nil {
		return errors.Wrap(err, "Error appear when tryp to create new perfdata")
	}

	m.perfdatas = append(m.perfdatas, perfdata)

	return nil
}

// Perfdatas permit to get all perfdatas
func (m *Monitoring) Perfdatas() modelPerfdata.Perfdatas {
	return m.perfdatas
}

// Perfdata permit to get perfdata on given index
func (m *Monitoring) Perfdata(index int) (*modelPerfdata.Perfdata, error) {
	if index >= len(m.perfdatas) {
		return nil, errors.New("Index is out of list messages")
	}

	return m.perfdatas[index], nil
}

// ToString permit to get string from monitoring data
func (m *Monitoring) ToString() string {

	var buffer bytes.Buffer

	for idx, message := range m.messages {
		if idx == 0 {
			buffer.WriteString(fmt.Sprintf("%s", message))
		} else {
			buffer.WriteString(fmt.Sprintf("\n%s", message))
		}
	}

	if len(m.Perfdatas()) > 0 {
		buffer.WriteString("|")
		for _, perfdata := range m.Perfdatas() {
			buffer.WriteString(fmt.Sprintf("%s=%d%s;;;; ", perfdata.Label(), perfdata.Value(), perfdata.Unit()))
		}
	}

	return buffer.String()

}

// ToSdtOut permit to print the state on stdout and exit with the right status code
func (m *Monitoring) ToSdtOut() {
	var status string
	switch m.Status() {
	case STATUS_UNKNOWN:
		status = "UNKNOWN"
	case STATUS_CRITICAL:
		status = "CRITICAL"
	case STATUS_WARNING:
		status = "WARNING"
	case STATUS_OK:
		status = "OK"
	}

	fmt.Printf("%s - %s\n", status, m.ToString())
	os.Exit(int(m.Status()))
}
