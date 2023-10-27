package mapping

import (
	"mnstr.at/mud/character"
	"mnstr.at/mud/item"
)

type Ground struct {
	exits       []Exit
	items       []item.Item
	characters  []character.Character
	description string
}

func NewGround(options ...AreaOption) *Ground {
	ground := &Ground{}
	for _, option := range options {
		option(ground)
	}
	return ground
}

func (g Ground) Exits() []Exit {
	return g.exits
}

func (g Ground) Items() []item.Item {
	return g.items
}

func (g Ground) Characters() []character.Character {
	return g.characters
}

func (g Ground) Description() string {
	return g.description
}
