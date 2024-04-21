package mediators

type Plane interface {
	Landing()
	TakeOff()
	permitLanding()
}
