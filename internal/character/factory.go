package character

import (
	"fmt"
	"github.com/Kuguchev/go-first-fl-codestyle/internal/utils"
)

type Class string

const (
	WarriorClass Class = "warrior"
	MageClass    Class = "mage"
	HealerClass  Class = "healer"
)

type Factory struct {
	baseStats          Stats
	characterTemplates map[Class]Character
}

func NewFactory() *Factory {
	factory := Factory{
		baseStats: Stats{
			attack:    5,
			defence:   10,
			endurance: 80,
		},
		characterTemplates: make(map[Class]Character),
	}
	factory.initializeTemplates()
	return &factory
}

func (f *Factory) initializeTemplates() {
	f.characterTemplates[WarriorClass] = &Warrior{
		BaseCharacter{
			id:          1,
			name:        "",
			className:   "Воитель",
			classKey:    string(WarriorClass),
			description: "— дерзкий воин ближнего боя. Сильный, выносливый и отважный.",
			Stats: Stats{
				attack:              f.baseStats.attack,
				defence:             f.baseStats.defence,
				endurance:           f.baseStats.endurance,
				damageBonus:         BonusRange{min: 3, max: 5},
				blockBonus:          BonusRange{min: 5, max: 10},
				specialAbilityBonus: 25,
				specialAbility:      "Выносливость",
			},
		},
	}

	f.characterTemplates[MageClass] = &Mage{
		BaseCharacter{
			id:          2,
			name:        "",
			className:   "Маг",
			classKey:    string(MageClass),
			description: "— находчивый воин дальнего боя. Обладает высоким интеллектом.",
			Stats: Stats{
				attack:              f.baseStats.attack,
				defence:             f.baseStats.defence,
				endurance:           f.baseStats.endurance,
				damageBonus:         BonusRange{min: 5, max: 10},
				blockBonus:          BonusRange{min: -2, max: 2},
				specialAbilityBonus: 40,
				specialAbility:      "Атака",
			},
		},
	}

	f.characterTemplates[HealerClass] = &Healer{
		BaseCharacter{
			id:          3,
			name:        "",
			className:   "Лекарь",
			classKey:    string(HealerClass),
			description: "— могущественный заклинатель. Черпает силы из природы, веры и духов.",
			Stats: Stats{
				attack:              f.baseStats.attack,
				defence:             f.baseStats.defence,
				endurance:           f.baseStats.endurance,
				damageBonus:         BonusRange{min: -3, max: -1},
				blockBonus:          BonusRange{min: 2, max: 5},
				specialAbilityBonus: 30,
				specialAbility:      "Защита",
			},
		},
	}
}

func (f *Factory) CreateCharacter(class Class, name string) (Character, error) {
	template, exists := f.characterTemplates[class]

	if !exists {
		return nil, fmt.Errorf("неизвестный класс персонажа: %s", class)
	}

	return template.Clone(name), nil
}

func (f *Factory) BaseStatsDescription() string {
	return fmt.Sprintf("Сейчас твоя выносливость — %d, атака — %d и защита — %d.",
		f.baseStats.endurance, f.baseStats.attack, f.baseStats.defence)
}

func (f *Factory) Templates() map[Class]Character {
	return f.characterTemplates
}

func (f *Factory) ClassNames() []string {
	chars := f.sortIds()

	names := make([]string, 0, len(chars))
	for _, char := range chars {
		names = append(names, char.ClassName())
	}

	return names
}

func (f *Factory) ClassExtendedNames() []string {
	chars := f.sortIds()

	names := make([]string, 0, len(chars))
	for _, char := range chars {
		names = append(names, fmt.Sprintf("%s — %s", char.ClassName(), char.ClassKey()))
	}

	return names
}

func (f *Factory) sortIds() []Character {
	chars := make([]Character, 0, len(f.characterTemplates))

	for _, char := range f.characterTemplates {
		chars = append(chars, char)
	}

	return utils.SortById(chars)
}
