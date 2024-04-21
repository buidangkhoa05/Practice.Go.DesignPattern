package mediators

import "fmt"

type MediatorDemo struct {
}

func (d MediatorDemo) Run() {
	atc := newAirCraftControl()
	cPlane := &CombatPlane{atc: atc}
	cmPlane := &CommercialPlane{atc: atc}

	cPlane.Landing()
	cmPlane.Landing()
	cPlane.TakeOff()
	cmPlane.TakeOff()

	fmt.Println("=====================================")
}
