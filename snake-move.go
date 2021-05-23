func (s *Snake) Move(duration time.Duration, processor elements.EventProcessor) error {
	playerOrb := s.Orbs[0]

	if s.Dead && !s.Lost {
		s.Lost = true
		return processor.ProcessEvent(event.Event{Type: "lose", From: s.GetID()})
	}
	if s.I.GenNew && time.Since(s.LastCreated) > time.Second {
		// creating new orb, behind the last one
	}
	if s.I.Moving {
		// moving code ommited, every next orb move towards previous
	}
	return nil
}
