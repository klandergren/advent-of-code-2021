package day11

import "fmt"

type Octopus struct {
	EnergyLevel int
	HasFlashed  bool
}

func (o *Octopus) CanFlash() bool {
	return (9 < o.EnergyLevel) && !o.HasFlashed
}

func (o *Octopus) Increment() {
	o.EnergyLevel += 1
}

func (o *Octopus) Flash() {
	o.HasFlashed = true
}

func (o *Octopus) SmartReset() {
	if o.HasFlashed {
		o.HasFlashed = false
		o.EnergyLevel = 0
	}
}

func (o *Octopus) Reset() {
	o.ResetFlash()
	o.ResetEnergyLevel()
}

func (o *Octopus) ResetFlash() {
	o.HasFlashed = false
}

func (o *Octopus) ResetEnergyLevel() {
	o.EnergyLevel = 0
}

func (o *Octopus) String() string {
	return fmt.Sprintf("%d", o.EnergyLevel)
}
