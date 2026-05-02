package main

import "fmt"

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Bird struct {
	food, locomotion, sound string
}

func (animals *Bird) Eat() {
	fmt.Println(animals.food)
}
func (animals *Bird) Move() {
	fmt.Println(animals.locomotion)
}
func (animals *Bird) Speak() {
	fmt.Println(animals.sound)
}

type Cow struct {
	food, locomotion, sound string
}

func (animals *Cow) Eat() {
	fmt.Println(animals.food)
}
func (animals *Cow) Move() {
	fmt.Println(animals.locomotion)
}
func (animals *Cow) Speak() {
	fmt.Println(animals.sound)
}

type Snake struct {
	food, locomotion, sound string
}

func (animals *Snake) Eat() {
	fmt.Println(animals.food)
}
func (animals *Snake) Move() {
	fmt.Println(animals.locomotion)
}
func (animals *Snake) Speak() {
	fmt.Println(animals.sound)
}

func main() {
	animals := make(map[string]Animal)

	var keyword, animal_name, animal_type string

	for {
		fmt.Print("> ")
		fmt.Scan(&keyword, &animal_name, &animal_type)

		if keyword == "newanimal" {
			switch animal_type {
			case "cow":
				animals[animal_name] = &Cow{"grass", "walk", "moo"}
			case "bird":
				animals[animal_name] = &Bird{"worms", "fly", "peep"}
			case "snake":
				animals[animal_name] = &Snake{"mice", "slither", "hiss"}
			default:
				fmt.Println("Unknown animal type")
				continue
			}
			fmt.Println("Created it!")

		} else if keyword == "query" {
			animal, ok := animals[animal_name]
			if !ok {
				fmt.Println("Animal not found")
				continue
			}

			switch animal_type {
			case "eat":
				animal.Eat()
			case "move":
				animal.Move()
			case "speak":
				animal.Speak()
			default:
				fmt.Println("Unknown query")
			}
		}
	}
}
