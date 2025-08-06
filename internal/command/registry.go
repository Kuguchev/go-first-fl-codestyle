package command

import (
	"fmt"
	"github.com/Kuguchev/go-first-fl-codestyle/internal/utils"
)

const (
	AttackCommandName  = "attack"
	DefenceCommandName = "defence"
	SpecialCommandName = "special"
	SkipCommandName    = "skip"
)

type Registry struct {
	commands map[string]Command
}

func NewRegistry() *Registry {
	registry := Registry{
		commands: make(map[string]Command),
	}
	registry.initCommands()
	return &registry
}

func (r *Registry) initCommands() {
	r.commands[AttackCommandName] = &AttackCommand{
		BaseCommand{
			id:          1,
			name:        AttackCommandName,
			description: AttackCommandName + " — чтобы атаковать противника,",
		},
	}
	r.commands[DefenceCommandName] = &DefenceCommand{
		BaseCommand{
			id:          2,
			name:        DefenceCommandName,
			description: DefenceCommandName + " — чтобы блокировать атаку противника,",
		},
	}
	r.commands[SpecialCommandName] = &SpecialCommand{
		BaseCommand{
			id:          3,
			name:        SpecialCommandName,
			description: SpecialCommandName + " — чтобы использовать свою суперсилу,",
		},
	}
	r.commands[SkipCommandName] = &SkipCommand{
		BaseCommand{
			id:          4,
			name:        SkipCommandName,
			description: "Если не хочешь тренироваться, введи команду " + SkipCommandName + ".",
		},
	}
}

func (r *Registry) Command(name string) (Command, error) {
	command, ok := r.commands[name]
	if !ok {
		return nil, fmt.Errorf("неизвестная команда: %s", name)
	}

	return command, nil
}

func (r *Registry) CommandsDescription() []string {
	commands := make([]Command, 0, len(r.commands))

	for _, command := range r.commands {
		commands = append(commands, command)
	}

	utils.SortById(commands)

	descriptions := make([]string, 0, len(commands))
	for _, command := range commands {
		descriptions = append(descriptions, command.Description())
	}

	return descriptions
}
