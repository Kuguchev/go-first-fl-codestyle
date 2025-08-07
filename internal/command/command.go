package command

import (
	"github.com/Kuguchev/go-first-fl-codestyle/internal/character"
	"github.com/Kuguchev/go-first-fl-codestyle/internal/utils"
)

type Command interface {
	utils.Identifiable
	Execute(character.Character) string
	Description() string
}

type BaseCommand struct {
	id          uint64
	name        string
	description string
}

func (c *BaseCommand) Id() uint64 {
	return c.id
}
