package exercices

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

var AnimalsAvailable = make(map[string]Animal)

type Cow struct {
	food, locomotion, noise string
}

func (a *Cow) Eat() {
	fmt.Println(a.food)
}

func (a *Cow) Move() {
	fmt.Println(a.locomotion)
}

func (a *Cow) Speak() {
	fmt.Println(a.noise)
}

type Bird struct {
	food, locomotion, noise string
}

func (a *Bird) Eat() {
	fmt.Println(a.food)
}

func (a *Bird) Move() {
	fmt.Println(a.locomotion)
}

func (a *Bird) Speak() {
	fmt.Println(a.noise)
}

type Snake struct {
	food, locomotion, noise string
}

func (a *Snake) Eat() {
	fmt.Println(a.food)
}

func (a *Snake) Move() {
	fmt.Println(a.locomotion)
}

func (a *Snake) Speak() {
	fmt.Println(a.noise)
}

func PrintAnimalInfo(a Animal, info string) {
	switch info {
	case "eat":
		a.Eat()
	case "move":
		a.Move()
	case "speak":
		a.Speak()
	default:
		fmt.Println("Input a valid argument.\nAvailable information: 'eat', 'move' and 'speak'")
	}
}

func CreateAnimal(animalType string) (Animal, error) {
	switch animalType {
	case "cow":
		return &Cow{food: "grass", locomotion: "walk", noise: "moo"}, nil
	case "bird":
		return &Bird{food: "worms", locomotion: "fly", noise: "peep"}, nil
	case "snake":
		return &Snake{food: "mice", locomotion: "slither", noise: "hss"}, nil
	default:
		return nil, fmt.Errorf("invalid animal: %s", animalType)
	}
}

func CmdCreate(name, animalType string) {
	if _, found := AnimalsAvailable[name]; found {
		fmt.Println("Animal already created!")
	}

	animal, err := CreateAnimal(animalType)
	if err != nil {
		fmt.Println(err, "\nType of animals avalable: 'cow', 'bird' and 'snake'")
		return
	}

	AnimalsAvailable[name] = animal
	fmt.Println("Created it!")
}

func CmdQuery(name, info string) {
	animal, found := AnimalsAvailable[name];
	if !found {
		fmt.Println("Animal was not found!")
		return
	}

	PrintAnimalInfo(animal, info)
}

func CmdHandler(cmd string, args []string) {
	switch cmd {
	case "newanimal":
		CmdCreate(args[0], args[1])
	case "query":
		CmdQuery(args[0], args[1])
	default:
		fmt.Println("Invalid cmd:", cmd, "\nAvailable cmd: 'newanimal' and 'query'")
	}
}

func ExecuteAnimals() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(">")
		scanner.Scan()
		input := scanner.Text()
		strs := strings.Split(input, " ")

		if len(strs) < 3 {
			fmt.Println("Missing arguments")
			continue
		}

		cmd := strs[0]
		CmdHandler(cmd, strs[1:3])
	}
}
