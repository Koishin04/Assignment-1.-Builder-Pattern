package character

type GameCharacterBuilder struct {
	data characterData
}

func NewGameCharacterBuilder() *GameCharacterBuilder {
	return &GameCharacterBuilder{
		data: characterData{
			abilities: make([]string, 0),
		},
	}
}

func (b *GameCharacterBuilder) SetName(name string) CharacterBuilder {
	b.data.name = name
	return b
}

func (b *GameCharacterBuilder) SetClass(charClass CharacterClass) CharacterBuilder {
	b.data.charClass = charClass
	return b
}

func (b *GameCharacterBuilder) SetLevel(level int) CharacterBuilder {
	b.data.level = level
	return b
}

func (b *GameCharacterBuilder) SetWeapon(weapon string) CharacterBuilder {
	b.data.weapon = weapon
	return b
}

func (b *GameCharacterBuilder) SetArmor(armor string) CharacterBuilder {
	b.data.armor = armor
	return b
}

func (b *GameCharacterBuilder) AddAbility(ability string) CharacterBuilder {
	b.data.abilities = append(b.data.abilities, ability)
	return b
}

func (b *GameCharacterBuilder) GetResult() (GameCharacter, error) {
	if err := validateCharacter(b.data); err != nil {
		return GameCharacter{}, err
	}

	return newGameCharacter(
		b.data.name,
		b.data.charClass,
		b.data.level,
		b.data.weapon,
		b.data.armor,
		b.data.abilities,
	), nil
}