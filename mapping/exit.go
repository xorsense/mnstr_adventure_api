package mapping

type Exit interface {
	Position() int
	ShortCode() string
	Description() string
}
