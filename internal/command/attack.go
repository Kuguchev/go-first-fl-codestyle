package command

import "github.com/Kuguchev/go-first-fl-codestyle/internal/character"

type AttackCommand struct {
	BaseCommand
}

func (c *AttackCommand) Execute(char character.Character) string {
	return char.Attack()
}

func (c *AttackCommand) Description() string {
	return c.description
}
