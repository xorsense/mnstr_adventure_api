package mapping

type PositionTransformer func(*Position)

type Position [3]int

func NewPosition(point [3]int, transformers ...PositionTransformer) *Position {
	p := Position(point)
	for _, transformer := range transformers {
		transformer(&p)
	}
	return &p
}
