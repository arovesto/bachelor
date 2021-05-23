type StaticBackground struct {
	Where     math.Box
	TextureID string
	ID        int
}

func (s *StaticBackground) Draw(c canvas.Canvas) {
	c.DrawShape(s.TextureID, s.Where, math.Box{Size: s.Where.Size})
}