package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a Animal) Eat() string {
	return a.food
}

func (a Animal) Move() string {
	return a.locomotion
}

func (a Animal) Speak() string {
	return a.noise
}

func Exec(a Animal, cmd string) {
	switch cmd {
	case "eat":
		fmt.Println(a.Eat())
	case "move":
		fmt.Println(a.Move())
	case "speak":
		fmt.Println(a.Speak())
	default:
		fmt.Printf("Can not identify %s.\n", cmd)
	}
}

func main() {
	// create new reader
	reader := bufio.NewReader(os.Stdin)
	var input string

	// init 3 animals
	cow := Animal{"grass", "walk", "moo"}
	bird := Animal{"worms", "fly", "peep"}
	snake := Animal{"mice", "slither", "hsss"}

	for {
		fmt.Printf("> ")
		input, _ = reader.ReadString('\n')
		input = strings.TrimSuffix(input, "\n")

		inputList := strings.Split(input, " ")

		if len(inputList) < 2 {
			fmt.Printf("Invaild input!")
			continue
		}

		animal, cmd := strings.ToLower(inputList[0]), strings.ToLower(inputList[1])

		switch animal {
		case "cow":
			Exec(cow, cmd)
		case "bird":
			Exec(bird, cmd)
		case "snake":
			Exec(snake, cmd)
		default:
			fmt.Printf("Can not identify %s.\n", inputList[0])
		}
	}
}