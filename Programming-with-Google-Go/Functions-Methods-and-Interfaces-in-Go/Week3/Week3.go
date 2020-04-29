package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food string
	locomotion string
	noise string
}
var animalMap map[string]*Animal

func main() {
	cow := new(Animal)
	bird := new(Animal)
	snake := new(Animal)
	cow.food = "grass"
	cow.locomotion = "walk"
	cow.noise = "moo"

	bird.food = "worms"
	bird.locomotion = "fly"
	bird.noise = "peep"

	snake.food = "mice"
	snake.locomotion = "slither"
	snake.noise = "hsss"
	animalMap = make(map[string]*Animal)
	animalMap["cow"] = cow
	animalMap["bird"] = bird
	animalMap["snake"] = snake

	reader := bufio.NewReader(os.Stdin)
	for true{
		fmt.Printf("> ")
		text, _ := reader.ReadString('\n')
		tstArr := strings.Fields(text)

		animalName := tstArr[0]
		requestName := tstArr[1]

		res := processRequest(animalName, requestName)
		fmt.Printf("%s\n", res)


	}
}

func (a *Animal) Eat() string {
	return a.food
}
func (a *Animal) Speak() string {
	return a.noise
}
func (a *Animal) Move() string {
	return a.locomotion
}

func processRequest(strAnimal string, strRequest string) string{
	animal := animalMap[strAnimal]
	switch strRequest {
	case "eat":
		return animal.Eat()
	case "move":
		return animal.Move()
	case "speak":
		return animal.Speak()
	default:
		return "Error"
	}
}

