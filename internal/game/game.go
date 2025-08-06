package game

import (
	"fmt"
	"github.com/Kuguchev/go-first-fl-codestyle/internal/input"
	"strings"

	"github.com/Kuguchev/go-first-fl-codestyle/internal/character"
	"github.com/Kuguchev/go-first-fl-codestyle/internal/command"
)

type Game struct {
	inputReader *input.Reader
	charFactory *character.Factory
	registry    *command.Registry
}

func NewGame() *Game {
	return &Game{
		inputReader: input.NewInputReader(),
		charFactory: character.NewFactory(),
		registry:    command.NewRegistry(),
	}
}

func (g *Game) Start() error {
	fmt.Println("Приветствую тебя, искатель приключений!")
	fmt.Println("Прежде чем начать игру...")

	var name string
	for {
		name = g.inputReader.ReadStringWithPrompt("...назови себя: ")
		if name == "" {
			fmt.Println("Имя не может быть пустым")
			continue
		}

		break
	}

	fmt.Printf("Здравствуй, %s!\n", name)
	fmt.Println(g.charFactory.BaseStatsDescription())

	class, err := g.chooseCharacterClass()

	if err != nil {
		return err
	}

	char, err := g.charFactory.CreateCharacter(class, name)

	if err != nil {
		return err
	}

	err = g.startTraining(char)

	if err != nil {
		return err
	}

	return nil
}

func (g *Game) chooseCharacterClass() (character.Class, error) {

	charTemplates := g.charFactory.Templates()

	if len(charTemplates) == 0 {
		return "", fmt.Errorf("нет доступных персонажей")
	}

	classNames := g.charFactory.ClassNames()

	fmt.Printf("Ты можешь выбрать один из доступных путей силы: ")
	fmt.Println(strings.Join(classNames, ", "))

	classExtendedNames := g.charFactory.ClassExtendedNames()

	for {
		prompt := fmt.Sprintf("Введи название персонажа, за которого хочешь играть: %s: ", strings.Join(classExtendedNames, ", "))
		class := character.Class(g.inputReader.ReadStringWithPrompt(prompt))

		template, exists := charTemplates[class]

		if !exists {
			fmt.Println("Такого персонажа нет.")
			continue
		}

		fmt.Println(template.ClassName(), template.Description())

		approveChoice := g.inputReader.ReadStringWithPrompt("Нажми (Y), чтобы подтвердить выбор, или любую другую кнопку, чтобы выбрать другого персонажа: ")

		if strings.ToLower(approveChoice) == "y" {
			return class, nil
		}
	}
}

func (g *Game) startTraining(c character.Character) error {
	descriptions := g.registry.CommandsDescription()

	if len(descriptions) == 0 {
		return fmt.Errorf("нет доступных команд")
	}

	fmt.Println(c.PersonalDescription())
	fmt.Println("Потренеруйся управлять своими навыками...")
	fmt.Print("Введи одну из команд: ")

	for i := range descriptions {
		fmt.Println(descriptions[i])
	}

	for {
		commandName := g.inputReader.ReadStringWithPrompt("Введи команду: ")
		cmd, err := g.registry.Command(commandName)

		if err != nil {
			fmt.Println(err)
			continue
		}

		if commandName == command.SkipCommandName {
			break
		}

		fmt.Println(cmd.Execute(c))
	}

	fmt.Println("тренировка окончена")
	return nil
}
