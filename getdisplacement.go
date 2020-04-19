package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
)

type Prompt int

const (
	RegexFloat          string = `([\d]+\.?\d*)`
	MaxLength int    = 50
)

const (
	ReadAcceleration Prompt = iota
	ReadVelocity
	ReadDisplacement
	ReadTime
	ShowStringValue
	RegexFound
	RegexNotFound
	MaxIntegerCount
)

func GetUserPrompt(up Prompt) string {
	var str string
	switch up {
	case ReadAcceleration:
		str = "enter acceleration:"
	case ReadVelocity:
		str = "enter initial velocity"
	case ReadDisplacement:
		str = "enter initial displacement"
	case ReadTime:
		str = "enter time"
	case RegexFound:
		str = "Valid Input!"
	case RegexNotFound:
		str = "Invalid Input!"
	}
	return str
}

func readUserInput(prompt Prompt, regex string) (float64, error) {
	var inputVal float64
	var err error
	re := regexp.MustCompile(regex)
	scanner := bufio.NewScanner(os.Stdin)
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
	fmt.Printf(GetUserPrompt(prompt))
	for scanner.Scan() {
		var strVal string = scanner.Text()
		match := re.Match([]byte(strVal))
		inputVal, err = strconv.ParseFloat(strVal, 64)
		if match == true && err == nil {
			return inputVal, err
		}
		fmt.Println(GetUserPrompt(RegexNotFound))
		fmt.Printf(GetUserPrompt(prompt))
	}
	return inputVal, err
}

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	fn := func(t float64) float64 {
		return (a * (math.Pow(t, 2) * 0.5) + (v0 * t) + (s0))
	}
	return fn
}

func main(){
	prompts := []Prompt{ReadAcceleration, ReadVelocity, ReadDisplacement, ReadTime}
	var acceleration, initialVelocity, initialDisplacement, time float64
	var err error

	for _, prompt := range prompts {
		switch prompt {
		case ReadAcceleration:
			acceleration, err = readUserInput(prompt, RegexFloat)
		case ReadVelocity:
			initialVelocity, err = readUserInput(prompt, RegexFloat)
		case ReadDisplacement:
			initialDisplacement, err = readUserInput(prompt, RegexFloat)
		case ReadTime:
			time, err = readUserInput(prompt, RegexFloat)
		}
	}

	if err == nil {
		fmt.Printf("acceleration: %f\n", acceleration)
		fmt.Printf("velocity:     %f\n", initialVelocity)
		fmt.Printf("displacement: %f\n", initialDisplacement)
		fmt.Printf("time:         %f\n", time)
	}
	fn := GenDisplaceFn(acceleration, initialVelocity, initialDisplacement)
	fmt.Printf("Displacement: %f\n", fn(time))

}