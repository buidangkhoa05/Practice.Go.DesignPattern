package mediators

import "fmt"

type CombatPlane struct {
	atc Mediator
}

func (c *CombatPlane) Landing() {
	if c.atc.canLanding(c) {
		fmt.Println("CombatPlane: Landed ...")
		return
	}
	fmt.Println("CombatPlane: Waiting ...")
}

func (c *CombatPlane) TakeOff() {
	fmt.Println("CombatPlane: TakeOff ...")
	c.atc.notifyTakeOff()
}

func (c *CombatPlane) permitLanding() {
	fmt.Println("CombatPlane: Is Permitted Landing ...")
	c.Landing()
}
