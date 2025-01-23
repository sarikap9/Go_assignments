package main

import (
	"testing"
)

func TestAnimalActions(t *testing.T) {
	cow := Animal{"grass", "walk", "moo"}
	bird := Animal{"worms", "fly", "peep"}
	snake := Animal{"mice", "slither", "hsss"}

	tests := []struct {
		animal   Animal
		action   string
		expected string
	}{
		{cow, "eat", "grass"},
		{cow, "move", "walk"},
		{cow, "speak", "moo"},
		{bird, "eat", "worms"},
		{bird, "move", "fly"},
		{bird, "speak", "peep"},
		{snake, "eat", "mice"},
		{snake, "move", "slither"},
		{snake, "speak", "hsss"},
	}

	for _, tt := range tests {
		var result string
		switch tt.action {
		case "eat":
			result = tt.animal.Eat()
		case "move":
			result = tt.animal.Move()
		case "speak":
			result = tt.animal.Speak()
		}

		if result != tt.expected {
			t.Errorf("For action %s, expected %s but got %s", tt.action, tt.expected, result)
		}
	}
}
