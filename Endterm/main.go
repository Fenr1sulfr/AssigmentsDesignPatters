package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"sync"
)

var validRaces = []string{
	"Blood Elf",
	"Draenei",
	"Dwarf",
	"Gnome",
	"Human",
	"Night Elf",
	"Orc",
	"Tauren",
	"Troll",
	"Undead",
}

var validClasses = []string{
	"Paladin",
	"Mage",
	"Warrior",
	"Priest",
	"Death Knight",
}

var raceOptions = map[int]string{
	1:  "Blood Elf",
	2:  "Draenei",
	3:  "Dwarf",
	4:  "Gnome",
	5:  "Human",
	6:  "Night Elf",
	7:  "Orc",
	8:  "Tauren",
	9:  "Troll",
	10: "Undead",
}

var classOptions = map[int]string{
	1: "Paladin",
	2: "Mage",
	3: "Warrior",
	4: "Priest",
	5: "Death Knight",
}

type Character struct {
	Name  string
	Race  string
	Class string
}

type CharacterBuilder struct {
	character *Character
}

var builderInstance *CharacterBuilder
var builderOnce sync.Once

func GetCharacterBuilder() *CharacterBuilder {
	builderOnce.Do(func() {
		builderInstance = &CharacterBuilder{character: &Character{}}
	})
	return builderInstance
}

func NewCharacterBuilder() *CharacterBuilder {
	return GetCharacterBuilder()
}

func (cb *CharacterBuilder) SetName(name string) *CharacterBuilder {
	cb.character.Name = name
	return cb
}

func (cb *CharacterBuilder) SetRace(race string) *CharacterBuilder {
	cb.character.Race = race
	return cb
}

func (cb *CharacterBuilder) SetClass(class string) *CharacterBuilder {
	cb.character.Class = class
	return cb
}

func (cb *CharacterBuilder) Build() *Character {
	return cb.character
}

func main() {
	fmt.Print("Do you want to create a new character (N) or load an existing character (L): ")
	choice := getUserInput()

	var character *Character

	if strings.ToLower(choice) == "n" {
		fmt.Print("Enter your character's name: ")
		name := getUserInput()

		var selectedRace string
		for {
			fmt.Println("Choose your character's race:")
			for i, validRace := range validRaces {
				fmt.Printf("%d. %s\n", i+1, validRace)
			}

			raceNumber := getUserInput()
			raceIndex, err := strconv.Atoi(raceNumber)
			if err != nil || raceIndex < 1 || raceIndex > len(raceOptions) {
				fmt.Println("Please enter a valid number.")
				continue
			}

			selectedRace = raceOptions[raceIndex]
			break
		}

		var selectedClass string
		for {
			fmt.Println("Choose your character's class:")
			for i, validClass := range validClasses {
				fmt.Printf("%d. %s\n", i+1, validClass)
			}

			classNumber := getUserInput()
			classIndex, err := strconv.Atoi(classNumber)
			if err != nil || classIndex < 1 || classIndex > len(classOptions) {
				fmt.Println("Please enter a valid number.")
				continue
			}

			selectedClass = classOptions[classIndex]
			break
		}

		character = NewCharacterBuilder().
			SetName(name).
			SetRace(selectedRace).
			SetClass(selectedClass).
			Build()

		saveCharacterData(character)
	} else if strings.ToLower(choice) == "l" {
		character = loadCharacterData()
		if character == nil {
			fmt.Println("No character data found.")
			return
		}
	} else {
		fmt.Println("Invalid choice. Please enter 'N' for a new character or 'L' for loading an existing character.")
		return
	}

	fmt.Printf("Character Name: %s\n", character.Name)
	fmt.Printf("Character Race: %s\n", character.Race)
	fmt.Printf("Character Class: %s\n", character.Class)
}

func getUserInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	return input
}

func saveCharacterData(character *Character) {
	data := fmt.Sprintf("Name: %s\nRace: %s\nClass: %s\n", character.Name, character.Race, character.Class)
	err := ioutil.WriteFile("data.txt", []byte(data), 0644)
	if err != nil {
		fmt.Println("Failed to save character data:", err)
	}
}

func loadCharacterData() *Character {
	data, err := ioutil.ReadFile("data.txt")
	if err != nil {
		return nil
	}

	lines := strings.Split(string(data), "\n")
	if len(lines) < 3 {
		return nil
	}

	name := strings.TrimPrefix(lines[0], "Name: ")
	race := strings.TrimPrefix(lines[1], "Race: ")
	class := strings.TrimPrefix(lines[2], "Class: ")

	return &Character{
		Name:  name,
		Race:  race,
		Class: class,
	}
}
