package character

import (
	"fmt"

	"github.com/Kuguchev/go-first-fl-codestyle/internal/utils"
)

type Character interface {
	utils.Identifiable
	Name() string
	Description() string
	PersonalDescription() string
	ClassName() string
	ClassKey() string
	Attack() string
	Defence() string
	UseSpecialAbility() string
	Clone(name string) Character
}

type BonusRange struct {
	min int
	max int
}

type Stats struct {
	attack              int
	defence             int
	endurance           int
	damageBonus         BonusRange
	blockBonus          BonusRange
	specialAbilityBonus int
	specialAbility      string
}

type BaseCharacter struct {
	id          uint64
	name        string
	className   string
	classKey    string
	description string
	Stats
}

func (c *BaseCharacter) Id() uint64 {
	return c.id
}

func (c *BaseCharacter) Name() string {
	return c.name
}

func (c *BaseCharacter) Description() string {
	return c.description
}

func (c *BaseCharacter) ClassName() string {
	return c.className
}

func (c *BaseCharacter) ClassKey() string {
	return c.classKey
}

func (c *BaseCharacter) Attack() string {
	damage := c.attack + utils.Random(c.damageBonus.min, c.damageBonus.max)
	return fmt.Sprintf("%s нанес урон противнику равный %d.", c.name, damage)
}

func (c *BaseCharacter) Defence() string {
	block := c.defence + utils.Random(c.blockBonus.min, c.blockBonus.max)
	return fmt.Sprintf("%s блокировал урона %d.", c.name, block)
}

func (c *BaseCharacter) BaseUseSpecialAbility(baseAbilityValue int) string {
	return fmt.Sprintf("%s применил специальное умение `%s %d`", c.name, c.specialAbility, baseAbilityValue+c.specialAbilityBonus)
}
