package mediators

import "fmt"

type AirTrafficControl struct {
	landingQueue []Plane
	isFree       bool
}

func newAirCraftControl() *AirTrafficControl {
	return &AirTrafficControl{
		isFree: true,
	}
}

func (atc *AirTrafficControl) canLanding(p Plane) bool {
	if atc.isFree {
		fmt.Println("AirTrafficControl: is free ...")
		return true
	}
	fmt.Println("AirTrafficControl: is not free ...")
	atc.landingQueue = append(atc.landingQueue, p)
	return false
}

func (atc *AirTrafficControl) notifyTakeOff() {
	if !atc.isFree {
		atc.isFree = true
	}

	if len(atc.landingQueue) > 0 {
		p := atc.landingQueue[0]
		atc.landingQueue = atc.landingQueue[1:]
		p.permitLanding()
	}
}
