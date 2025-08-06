package command

import "github.com/Kuguchev/go-first-fl-codestyle/internal/character"

type DefenceCommand struct {
	BaseCommand
}

func (d *DefenceCommand) Execute(char character.Character) string {
	return char.Defence()
}

func (d *DefenceCommand) Description() string {
	return d.description
}
