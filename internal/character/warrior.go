package character

import "fmt"

type Warrior struct {
	BaseCharacter
}

func (w *Warrior) UseSpecialAbility() string {
	return w.BaseUseSpecialAbility(w.endurance)
}

func (w *Warrior) PersonalDescription() string {
	return fmt.Sprintf("%s, ты %s - отличный боец ближнего боя.", w.name, w.className)
}

func (w *Warrior) Clone(name string) Character {
	newWarrior := &Warrior{
		BaseCharacter: BaseCharacter{
			name:        name,
			className:   w.className,
			description: w.description,
			Stats:       w.Stats,
		},
	}

	return newWarrior
}
