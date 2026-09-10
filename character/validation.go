package character

import "errors"

type characterData struct {
	name      string
	charClass CharacterClass
	level     int
	weapon    string
	armor     string
	abilities []string
}

func validateCharacter(data characterData) error {
	if data.name == "" {
		return errors.New("character name is required")
	}

	if !isValidClass(data.charClass) {
		return errors.New("character class must be Warrior, Mage, or Villain")
	}

	if data.level < MinLevel || data.level > MaxLevel {
		return errors.New("character level must be between 1 and 100")
	}

	if data.weapon == "" {
		return errors.New("weapon is required")
	}

	if data.armor == "" {
		return errors.New("armor is required")
	}

	if len(data.abilities) == 0 {
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