package character

type CharacterBuilder interface {
	SetName(name string) CharacterBuilder
	SetClass(charClass CharacterClass) CharacterBuilder
	SetLevel(level int) CharacterBuilder
	SetWeapon(weapon string) CharacterBuilder
	SetArmor(armor string) CharacterBuilder
	AddAbility(ability string) CharacterBuilder
}