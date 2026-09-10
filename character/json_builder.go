package character

import (
	"encoding/json"
	"errors"
)

type JSONCharacterBuilder struct {
	name      string
	charClass CharacterClass
	level     int
	weapon    string
	armor     string
	abilities []string
}

func NewJSONCharacterBuilder() *JSONCharacterBuilder {
	return &JSONCharacterBuilder{
		abilities: make([]string, 0),
	}
}

func (b *JSONCharacterBuilder) SetName(name string) CharacterBuilder {
	b.name = name
	return b
}

func (b *JSONCharacterBuilder) SetClass(charClass CharacterClass) CharacterBuilder {
	b.charClass = charClass
	return b
}

func (b *JSONCharacterBuilder) SetLevel(level int) CharacterBuilder {
	b.level = level
	return b
}

func (b *JSONCharacterBuilder) SetWeapon(weapon string) CharacterBuilder {
	b.weapon = weapon
	return b
}

func (b *JSONCharacterBuilder) SetArmor(armor string) CharacterBuilder {
	b.armor = armor
	return b
}

func (b *JSONCharacterBuilder) AddAbility(ability string) CharacterBuilder {
	b.abilities = append(b.abilities, ability)
	return b
}

func (b *JSONCharacterBuilder) GetResult() (string, error) {
	if err := b.validate(); err != nil {
		return "", err
	}

	data := struct {
		Name      string         `json:"name"`
		Class     CharacterClass `json:"class"`
		Level     int            `json:"level"`
		Weapon    string         `json:"weapon"`
		Armor     string         `json:"armor"`
		Abilities []string       `json:"abilities"`
	}{
		Name:      b.name,
		Class:     b.charClass,
		Level:     b.level,
		Weapon:    b.weapon,
		Armor:     b.armor,
		Abilities: b.abilities,
	}

	result, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", err
	}

	return string(result), nil
}

func (b *JSONCharacterBuilder) validate() error {
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