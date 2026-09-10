package character

import (
	"encoding/json"
)

type JSONCharacterBuilder struct {
	data characterData
}

func NewJSONCharacterBuilder() *JSONCharacterBuilder {
	return &JSONCharacterBuilder{
		data: characterData{
			abilities: make([]string, 0),
		},
	}
}

func (b *JSONCharacterBuilder) SetName(name string) CharacterBuilder {
	b.data.name = name
	return b
}

func (b *JSONCharacterBuilder) SetClass(charClass CharacterClass) CharacterBuilder {
	b.data.charClass = charClass
	return b
}

func (b *JSONCharacterBuilder) SetLevel(level int) CharacterBuilder {
	b.data.level = level
	return b
}

func (b *JSONCharacterBuilder) SetWeapon(weapon string) CharacterBuilder {
	b.data.weapon = weapon
	return b
}

func (b *JSONCharacterBuilder) SetArmor(armor string) CharacterBuilder {
	b.data.armor = armor
	return b
}

func (b *JSONCharacterBuilder) AddAbility(ability string) CharacterBuilder {
	b.data.abilities = append(b.data.abilities, ability)
	return b
}

func (b *JSONCharacterBuilder) GetResult() (string, error) {
	if err := validateCharacter(b.data); err != nil {
		return "", err
	}

	result, err := json.MarshalIndent(struct {
		Name      string         `json:"name"`
		Class     CharacterClass `json:"class"`
		Level     int            `json:"level"`
		Weapon    string         `json:"weapon"`
		Armor     string         `json:"armor"`
		Abilities []string       `json:"abilities"`
	}{
		Name:      b.data.name,
		Class:     b.data.charClass,
		Level:     b.data.level,
		Weapon:    b.data.weapon,
		Armor:     b.data.armor,
		Abilities: b.data.abilities,
	}, "", "  ")

	if err != nil {
		return "", err
	}

	return string(result), nil
}