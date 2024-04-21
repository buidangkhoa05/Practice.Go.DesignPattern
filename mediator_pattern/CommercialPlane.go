package mediators

import "fmt"

type CommercialPlane struct {
	atc Mediator
}

func (c *CommercialPlane) Landing() {
	if c.atc.canLanding(c) {
		fmt.Println("CommercialPlane: Landed ...")
		return
	}
	fmt.Println("CombatPlane: Waiting ...")
}

func (c *CommercialPlane) TakeOff() {
	fmt.Println("CommercialPlane: TakeOff ...")
	c.atc.notifyTakeOff()
}

func (c *CommercialPlane) permitLanding() {
	fmt.Println("CombatPlane: Is Permitted Landing ...")
	c.Landing()
}
