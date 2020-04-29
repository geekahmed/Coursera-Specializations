package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var animalMap map[string]Animal
var animalArray = []string{"cow", "bird", "snake"}
var requestArray = []string{"eat", "move", "speak"}
var commandArray = []string{"newanimal", "query"}

func main() {
	reader := bufio.NewReader(os.Stdin)
	animalMap = make(map[string]Animal)
	for true{
		fmt.Printf("> ")
		text, _ := reader.ReadString('\n')
		tstArr := strings.Fields(text)
		nameOfNewAnimal := tstArr[1]
		switch tstArr[0] {
		case commandArray[0]:
			typeOfNewAnimal := tstArr[2]
			createNewAnimal(nameOfNewAnimal, typeOfNewAnimal)
		case commandArray[1]:
			info := tstArr[2]
			processRequest(nameOfNewAnimal, info)
		}


	}
}

func processRequest(strAnimal string, strRequest string){
	animal := animalMap[strAnimal]
	switch strRequest {
	case requestArray[0]:
		fmt.Printf("%s eats %s\n", strAnimal, animal.Eat())
	case requestArray[1]:
		fmt.Printf("%s moves by %s\n", strAnimal, animal.Move())
	case requestArray[2]:
		fmt.Printf("%s speaks %s\n", strAnimal, animal.Speak())
	default:
		fmt.Errorf("%s: Invalid Request", strRequest)
	}
}
func createNewAnimal(animalName, animalType string) {
	switch animalType {
	case animalArray[0]:
		animal := Cow{food: "grass", locomotion: "walk", noise: "moo"}
		animalMap[animalName] = animal
	case animalArray[1]:
		animal := Bird{food: "worms", locomotion: "fly", noise: "peep"}
		animalMap[animalName] = animal
	case animalArray[2]:
		animal := Snake{food: "mice", locomotion: "slither", noise: "hsss"}
		animalMap[animalName] = animal
	default:
		fmt.Errorf("%s: Invalid animal type", animalType)
		return
	}
	fmt.Println("Created it!")
}


type Animal interface {
	Eat() string
	Move() string
	Speak() string
}

type Cow struct{ food, locomotion, noise string }

func (a Cow) Eat() string {
	return a.food
}
func (a Cow) Move() string {
	return a.locomotion
}
func (a Cow) Speak() string {
	return a.noise
}

type Bird struct{ food, locomotion, noise string }

func (a Bird) Eat() string {
	return a.food
}

func (a Bird) Move() string {
	return a.locomotion
}

func (a Bird) Speak() string {
	return a.noise
}

type Snake struct{ food, locomotion, noise string }

func (a Snake) Eat() string {
	return a.food
}

func (a Snake) Move() string {
	return a.locomotion
}

func (a Snake) Speak() string {
	return a.noise
}

