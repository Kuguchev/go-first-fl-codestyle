package character

import "fmt"

type Mage struct {
	BaseCharacter
}

func (m *Mage) UseSpecialAbility() string {
	return m.BaseUseSpecialAbility(m.attack)
}

func (m *Mage) PersonalDescription() string {
	return fmt.Sprintf("%s, ты %s - превосходный укротитель стихий.", m.name, m.className)
}

func (m *Mage) Clone(name string) Character {
	newMage := &Mage{
		BaseCharacter: BaseCharacter{
			name:        name,
			className:   m.className,
			description: m.description,
			Stats:       m.Stats,
		},
	}
	return newMage
}
