package mapping

import (
	"github.com/xorsense/mnstr_adventure_api/character"
	"github.com/xorsense/mnstr_adventure_api/item"
)

type AreaOption func(Area)

type Area interface {
	Exits() []Exit
	Items() []item.Item
	Characters() []character.Character
	Description() string
}
