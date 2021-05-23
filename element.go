type Element interface {
	Draw(c canvas.Canvas) // all actions to draw the object
	GetID() int           // for cross element actions
	GetType() int         // for cross element actions
	GetState() ([]byte, error)  // additional method, implement if required                
	SetState([]byte) error // additional method, implement if required
}