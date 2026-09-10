package main

import (
	"fmt"
	"log"

	"Assignment-1-Builder-Pattern/character"
)

func main() {
	director := character.NewCharacterDirector()

	fmt.Println("=== BUILDER PATTERN: GAME CHARACTER ===")
	fmt.Println()

	demonstrateObjectBuilder(director)
	fmt.Println()

	demonstrateJSONBuilder(director)
	fmt.Println()

	demonstrateDifferentConfigurations(director)
}

func demonstrateObjectBuilder(director *character.CharacterDirector) {
	builder := character.NewGameCharacterBuilder()

	director.MakeWarrior(builder)

	warrior, err := builder.GetResult()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("--- GameCharacter Object ---")
	fmt.Println("Name:", warrior.Name())
	fmt.Println("Class:", warrior.Class())
	fmt.Println("Level:", warrior.Level())
	fmt.Println("Weapon:", warrior.Weapon())
	fmt.Println("Armor:", warrior.Armor())
	fmt.Println("Abilities:", warrior.Abilities())
}

func demonstrateJSONBuilder(director *character.CharacterDirector) {
	builder := character.NewJSONCharacterBuilder()

	director.MakeWarrior(builder)

	warriorJSON, err := builder.GetResult()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("--- JSON Representation ---")
	fmt.Println(warriorJSON)
}

func demonstrateDifferentConfigurations(director *character.CharacterDirector) {
	objectBuilder := character.NewGameCharacterBuilder()
	director.MakeMage(objectBuilder)

	mage, err := objectBuilder.GetResult()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("--- Mage Configuration ---")
	fmt.Println("Name:", mage.Name())
	fmt.Println("Class:", mage.Class())
	fmt.Println("Level:", mage.Level())
	fmt.Println("Weapon:", mage.Weapon())
	fmt.Println("Armor:", mage.Armor())
	fmt.Println("Abilities:", mage.Abilities())

	jsonBuilder := character.NewJSONCharacterBuilder()
	director.MakeVillain(jsonBuilder)

	villainJSON, err := jsonBuilder.GetResult()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println()
	fmt.Println("--- Villain JSON Representation ---")
	fmt.Println(villainJSON)
}