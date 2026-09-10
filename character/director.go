package character

type CharacterDirector struct{}

func NewCharacterDirector() *CharacterDirector {
	return &CharacterDirector{}
}

func (d *CharacterDirector) MakeWarrior(b CharacterBuilder) {
	b.SetName("Arthur").
		SetClass(Warrior).
		SetLevel(10).
		SetWeapon("Long Sword").
		SetArmor("Knight Armor").
		AddAbility("Power Strike").
		AddAbility("Shield Wall")
}

func (d *CharacterDirector) MakeMage(b CharacterBuilder) {
	b.SetName("Merlin").
		SetClass(Mage).
		SetLevel(10).
		SetWeapon("Magic Staff").
		SetArmor("Wizard Robe").
		AddAbility("Fireball").
		AddAbility("Teleport").
		AddAbility("Ice Blast")
}

func (d *CharacterDirector) MakeVillain(b CharacterBuilder) {
	b.SetName("Dark Lord").
		SetClass(Villain).
		SetLevel(20).
		SetWeapon("Shadow Blade").
		SetArmor("Dark Armor").
		AddAbility("Dark Strike").
		AddAbility("Summon Minions").
		AddAbility("Shadow Step")
}