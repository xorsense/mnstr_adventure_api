package mapping

import (
	"mnstr.at/mud/character"
	"mnstr.at/mud/item"
)

type AreaOption func(Area)

type Area interface {
	Exits() []Exit
	Items() []item.Item
	Characters() []character.Character
	Description() string
}
