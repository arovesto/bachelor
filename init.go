func init() {
	elements.GenElements[SnakeType] = func() elements.Element {
		return &Snake{}
	}
}