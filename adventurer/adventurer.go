package adventurer

import (
	"context"
	"time"

	"github.com/uptrace/bun"
	"github.com/xorsense/mnstr_adventure_api/database"
)

type AdventurerOptions func(*Adventurer)

type Adventurer struct {
	bun.BaseModel `bun:"table:adventurers"`
	ID            int       `bun:"id,pk,autoincrement"`
	Name          string    `bun:"name,notnull"`
	PasswordHash  string    `bun:"password_hash,notnull"`
	CreatedAt     time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp"`
	UpdatedAt     time.Time `bun:"updated_at,nullzero,notnull,default:current_timestamp"`
}

func New(options ...AdventurerOptions) *Adventurer {
	a := Adventurer{}
	for _, option := range options {
		option(&a)
	}
	return &a
}

func FindWithName(name string) (*Adventurer, error) {
	db, err := database.Instance()
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	var a *Adventurer
	if err := db.NewSelect().Model(a).Column("id", "name", "password_hash").Scan(ctx); err != nil {

		return nil, err
	}
	return a, nil
}
