func (s *Snake) Draw(c canvas.Canvas) {
	for i, o := range s.Orbs {
		if i == 0 {
			c.DrawColor(color.RGBA{G: 255, A: 255}, o, o)
		} else {
			c.DrawColor(color.RGBA{R: 255, A: 255}, o, o)
		}
	}
}

func (s *Snake) Input() ([]byte, error) {
	s.I.Target = input.MousePosition
	//o.I.Moving = input.MousePressed
	s.I.Moving = true
	s.I.GenNew = input.Pressed[input.KEY_SPACE]
	return json.Marshal(s.I)
}


func (s *Snake) Collide(other elements.Collidable) error {
	info := math.Collide(s.Orbs[0], other.Collider())
	if info.Collided {
		s.Dead = true
	}
	return nil
}