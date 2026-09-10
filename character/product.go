package character

type CharacterClass string

const (
	Warrior CharacterClass = "Warrior"
	Mage    CharacterClass = "Mage"
	Villain CharacterClass = "Villain"
)

const (
	MinLevel = 1
	MaxLevel = 100
)

type GameCharacter struct {
	name          string
	charClass     CharacterClass
	level         int
	weapon        string
	armor         string
	abilities     []string
}

func newGameCharacter(
	name string,
	charClass CharacterClass,
	level int,
	weapon string,
	armor string,
	abilities []string,
) GameCharacter {
	return GameCharacter{
		name:      name,
		charClass: charClass,
		level:     level,
		weapon:    weapon,
		armor:     armor,
		abilities: append([]string(nil), abilities...),
	}
}



func (c GameCharacter) Name() string {
	return c.name
}

func (c GameCharacter) Class() CharacterClass {
	return c.charClass
}

func (c GameCharacter) Level() int {
	return c.level
}

func (c GameCharacter) Weapon() string {
	return c.weapon
}

func (c GameCharacter) Armor() string {
	return c.armor
}

func (c GameCharacter) Abilities() []string {
	return append([]string(nil), c.abilities...)
}