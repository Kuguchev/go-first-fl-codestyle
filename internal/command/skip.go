package command

import "github.com/Kuguchev/go-first-fl-codestyle/internal/character"

type SkipCommand struct {
	BaseCommand
}

func (s *SkipCommand) Execute(char character.Character) string {
	return ""
}

func (s *SkipCommand) Description() string {
	return s.description
}
