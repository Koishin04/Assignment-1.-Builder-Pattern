package character

import "errors"

type GameCharacterBuilder struct {
	name      string
	charClass CharacterClass
	level     int
	weapon    string
	armor     string
	abilities []string
}

func NewGameCharacterBuilder() *GameCharacterBuilder {
	return &GameCharacterBuilder{
		abilities: make([]string, 0),
	}
}

func (b *GameCharacterBuilder) SetName(name string) CharacterBuilder {
	b.name = name
	return b
}

func (b *GameCharacterBuilder) SetClass(charClass CharacterClass) CharacterBuilder {
	b.charClass = charClass
	return b
}

func (b *GameCharacterBuilder) SetLevel(level int) CharacterBuilder {
	b.level = level
	return b
}

func (b *GameCharacterBuilder) SetWeapon(weapon string) CharacterBuilder {
	b.weapon = weapon
	return b
}

func (b *GameCharacterBuilder) SetArmor(armor string) CharacterBuilder {
	b.armor = armor
	return b
}

func (b *GameCharacterBuilder) AddAbility(ability string) CharacterBuilder {
	b.abilities = append(b.abilities, ability)
	return b
}

func (b *GameCharacterBuilder) GetResult() (GameCharacter, error) {
	if err := b.validate(); err != nil {
		return GameCharacter{}, err
	}

	return newGameCharacter(
		b.name,
		b.charClass,
		b.level,
		b.weapon,
		b.armor,
		b.abilities,
	), nil
}

func (b *GameCharacterBuilder) validate() error {
	if b.name == "" {
		return errors.New("character name is required")
	}

	if !isValidClass(b.charClass) {
		return errors.New("character class must be Warrior, Mage, or Villain")
	}

	if b.level < MinLevel || b.level > MaxLevel {
		return errors.New("character level must be between 1 and 100")
	}

	if b.weapon == "" {
		return errors.New("weapon is required")
	}

	if b.armor == "" {
		return errors.New("armor is required")
	}

	if len(b.abilities) == 0 {
		return errors.New("at least one ability is required")
	}

	return nil
}

func isValidClass(charClass CharacterClass) bool {
	switch charClass {
	case Warrior, Mage, Villain:
		return true
	default:
		return false
	}
}