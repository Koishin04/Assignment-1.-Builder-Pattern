package main

import (
	"fmt"
	"log"

	"Assignment-1-Builder-Pattern/character"
)

func main() {
	director := character.NewCharacterDirector()

	objectBuilder := character.NewGameCharacterBuilder()

	director.MakeWarrior(objectBuilder)

	warrior, err := objectBuilder.GetResult()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("=== GameCharacter Object ===")
	fmt.Println("Name:", warrior.Name())
	fmt.Println("Class:", warrior.Class())
	fmt.Println("Level:", warrior.Level())
	fmt.Println("Weapon:", warrior.Weapon())
	fmt.Println("Armor:", warrior.Armor())
	fmt.Println("Abilities:", warrior.Abilities())

	fmt.Println()

	
	jsonBuilder := character.NewJSONCharacterBuilder()

	director.MakeWarrior(jsonBuilder)

	warriorJSON, err := jsonBuilder.GetResult()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("=== JSON Representation ===")
	fmt.Println(warriorJSON)

	fmt.Println()

	mageBuilder := character.NewGameCharacterBuilder()

	director.MakeMage(mageBuilder)

	mage, err := mageBuilder.GetResult()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("=== Mage ===")
	fmt.Println("Name:", mage.Name())
	fmt.Println("Class:", mage.Class())
	fmt.Println("Level:", mage.Level())
	fmt.Println("Weapon:", mage.Weapon())
	fmt.Println("Armor:", mage.Armor())
	fmt.Println("Abilities:", mage.Abilities())

	fmt.Println()

	
	villainBuilder := character.NewJSONCharacterBuilder()

	director.MakeVillain(villainBuilder)

	villainJSON, err := villainBuilder.GetResult()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("=== Villain JSON ===")
	fmt.Println(villainJSON)
}