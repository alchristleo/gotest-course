package main

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

type Cow struct {
	name string
}

type Bird struct {
	name string
}

type Snake struct {
	name string
}

func (cow Cow) Eat() {
	fmt.Println("grass")
}

func (cow Cow) Move() {
	fmt.Println("walk")
}

func (cow Cow) Speak() {
	fmt.Println("moo")
}

func (bird Bird) Eat() {
	fmt.Println("worms")
}

func (bird Bird) Move() {
	fmt.Println("fly")
}

func (bird Bird) Speak() {
	fmt.Println("peep")
}

func (snake Snake) Eat() {
	fmt.Println("mice")
}

func (snake Snake) Move() {
	fmt.Println("slither")
}

func (snake Snake) Speak() {
	fmt.Println("hsss")
}


func addAnimal(m map[string]Animal, name string, t string) {
	switch t {
	case "cow":
		m[name] = Cow{name}
	case "bird":
		m[name] = Bird{name}
	case "snake":
		m[name] = Snake{name}
	default:
		fmt.Printf("Can not identify type: %s.\n", t)
	}
	fmt.Println("Created it!")
}

func printAnimal(m map[string]Animal, name string, cmd string) {
	switch cmd {
	case "eat":
		// handle error if name not in map
		if animal, ok := m[name]; ok {
			animal.Eat()
		} else {
			fmt.Printf("Animal named %s do not exist.", name)
		}
	case "move":
		// handle error if name not in map
		if animal, ok := m[name]; ok {
			animal.Move()
		} else {
			fmt.Printf("Animal named %s do not exist.", name)
		}
	case "speak":
		// handle error if name not in map
		if animal, ok := m[name]; ok {
			animal.Speak()
		} else {
			fmt.Printf("Animal named %s do not exist.", name)
		}
	default:
		fmt.Printf("Can not identify cmd: %s.\n", cmd)
	}
}

func main() {
	// init a map to save animals
	animalMap := make(map[string]Animal)

	// create new reader
	reader := bufio.NewReader(os.Stdin)
	var input string

	for true {
		fmt.Printf("> ")
		input, _ = reader.ReadString('\n')
		input = strings.TrimSuffix(input, "\n")

		inputList := strings.Split(input, " ")

		cmd := strings.ToLower(inputList[0])

		if cmd == "newanimal" {
			if len(inputList) < 3 {
				fmt.Println("Invalid input!")
				continue
			}
			addAnimal(animalMap, strings.ToLower(inputList[1]), strings.ToLower(inputList[2]))
		} else if cmd == "query" {
			if len(inputList) < 3 {
				fmt.Println("Invalid input!")
				continue
			}
			printAnimal(animalMap, strings.ToLower(inputList[1]), strings.ToLower(inputList[2]))
		} else {
			fmt.Println("Invalid input!")
		}
	}
}


