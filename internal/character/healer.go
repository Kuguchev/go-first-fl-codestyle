package character

import "fmt"

type Healer struct {
	BaseCharacter
}

func (h *Healer) UseSpecialAbility() string {
	return h.BaseUseSpecialAbility(h.defence)
}

func (h *Healer) PersonalDescription() string {
	return fmt.Sprintf("%s, ты %s - чародей, способный исцелять раны.", h.name, h.className)
}

func (h *Healer) Clone(name string) Character {
	newHealer := &Healer{
		BaseCharacter: BaseCharacter{
			name:        name,
			className:   h.className,
			description: h.description,
			Stats:       h.Stats,
		},
	}
	return newHealer
}
