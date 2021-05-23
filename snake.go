type Snake struct {
	Orbs []math.Sphere
	ID   int

	Vel  math.Vector
	Dead bool

	I           PlayerInput
	Lost        bool
	LastCreated time.Time
}
