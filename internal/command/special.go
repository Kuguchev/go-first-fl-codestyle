package command

import "github.com/Kuguchev/go-first-fl-codestyle/internal/character"

type SpecialCommand struct {
	BaseCommand
}

func (s *SpecialCommand) Execute(char character.Character) string {
	return char.UseSpecialAbility()
}

func (s *SpecialCommand) Description() string {
	return s.description
}
