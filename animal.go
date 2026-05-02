package main

import (
	"fmt"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	sound      string
}

func (animals *Animal) Eat() {
	fmt.Println(animals.food)
}
func (animals *Animal) Move() {
	fmt.Println(animals.locomotion)
}
func (animals *Animal) Speak() {
	fmt.Println(animals.sound)
}

func main() {
	bird := Animal{"worms", "fly", "peep"}
	cow := Animal{"grass", "walk", "moo"}
	snake := Animal{"mice", "slither", "hiss"}

	var animal string
	var action string

	for {
		fmt.Println("Please enter the animal and action you want to check in the prompt below(in lower case). ")
		fmt.Printf("> ")
		count, err := fmt.Scan(&animal, &action)

		var animals = map[string]Animal{
			"bird":  bird,
			"cow":   cow,
			"snake": snake,
		}

		if err != nil {
			fmt.Println("An unexpected error occured!")
		}
		if count != 2 {
			fmt.Println("Submit two spaced arguments.")
		}
		animal = strings.ToLower(animal)
		mammal := animals[animal]

		action = strings.ToLower(action)

		animalAction(action, mammal)
	}
}

func animalAction(action string, mammal Animal) {
	if action == "eat" {
		mammal.Eat()
	}
	if action == "move" {
		mammal.Move()
	}
	if action == "speak" {
		mammal.Speak()
	}
}
