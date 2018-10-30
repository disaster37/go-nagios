/*
	This file permit to handle the perfdata
*/
package nagiosPlugin

import (
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// Perfdata struct
type Perfdata struct {
	unit  string
	value int
	label string
}

type Perfdatas []*Perfdata

// NewPerfdata permit to create new perfdata struct
// It 's' return Perfdata object
// It's return error if label is empty
func NewPerfdata(label string, value int, unit string) (*Perfdata, error) {

	log.Debugf("label: %s, value: %d, unit: %s", label, value, unit)

	if label == "" {
		return nil, errors.New("Label can't be empty")
	}

	perfdata := &Perfdata{
		label: label,
		value: value,
		unit:  unit,
	}

	return perfdata, nil
}

// Unit permit to get unit
// It's return unit as string
func (p *Perfdata) Unit() string {
	return p.unit
}

// SetUnit permit to set unit
func (p *Perfdata) SetUnit(unit string) {
	log.Debug("Unit: ", unit)

	p.unit = unit
}

// Label permit to get label
func (p *Perfdata) Label() string {
	return p.label
}

// SetLabel permit to set label
func (p *Perfdata) SetLabel(label string) error {
	log.Debug("Label: ", label)

	if label == "" {
		return errors.New("Label can't be empty")
	}

	p.label = label

	return nil
}

// Value permit to get value
func (p *Perfdata) Value() int {
	return p.value
}

// SetValue permit to set value
func (p *Perfdata) SetValue(value int) {
	log.Debug("Value: ", value)

	p.value = value
}
