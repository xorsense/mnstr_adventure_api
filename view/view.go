package view

import "io"

type View interface {
	Name() string
	Description() string
	Render(io.Writer, ...any) (int, error)
}
