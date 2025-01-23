package main

import (
	"fmt"
)

type Animal struct {
	food       string
	locomotion string
	sound      string
}

func (a Animal) Eat() string {
	return a.food
}

func (a Animal) Move() string {
	return a.locomotion
}

func (a Animal) Speak() string {
	return a.sound
}

func main() {
	animals := map[string]Animal{
		"cow":   {"grass", "walk", "moo"},
		"bird":  {"worms", "fly", "peep"},
		"snake": {"mice", "slither", "hsss"},
	}

	for {
		var animalName, action string

		fmt.Print("Enter animal name (cow, bird, snake): ")
		fmt.Scan(&animalName)
		fmt.Print("Enter action they can perform (eat, move, speak): ")
		fmt.Scan(&action)

		animal, exists := animals[animalName]
		if !exists {
			fmt.Println("Invalid animal name.")
			continue
		}

		switch action {
		case "eat":
			fmt.Println(animal.Eat())
		case "move":
			fmt.Println(animal.Move())
		case "speak":
			fmt.Println(animal.Speak())
		default:
			fmt.Println("Invalid action.")
		}
	}
}
