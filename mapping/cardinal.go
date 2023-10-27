package mapping

type Cardinal struct {
	position    Position
	shortCode   string
	description string
}

func (c Cardinal) Position() Position {
	return c.position
}

func (c Cardinal) ShortCode() string {
	return c.shortCode
}

func (c Cardinal) Description() string {
	return c.description
}
