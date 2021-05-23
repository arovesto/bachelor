type Playable interface {
	Element
	Input() ([]byte, error) // result used in "input" event
	SetInput([]byte) error // parameter passed from "input" event
}

type Collidable interface {
	Element
	Collide(other Collidable) error // collision should be checked inside
	Collider() math.Shape           // useful in Collide
}

type Movable interface {
	Element
	Move(duration time.Duration, processor EventProcessor) error // all changes in object
}
