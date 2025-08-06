package command

import "github.com/Kuguchev/go-first-fl-codestyle/internal/character"

type Command interface {
	Id() uint64
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
