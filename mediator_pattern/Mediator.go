package mediators

type Mediator interface {
	canLanding(a Plane) bool
	notifyTakeOff()
}
