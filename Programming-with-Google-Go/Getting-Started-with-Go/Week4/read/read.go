package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	type name struct {
		fname string
		lname string
	}

	myslc := make([]name, 0)
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Enter the Name of the file: ")
	fileName, _ := reader.ReadString('\n')
	fileName = strings.TrimSuffix(fileName, "\n")

	// Open file for reading
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		testArray := strings.Fields(line)
		currentName := name{fname: testArray[0]}
		currentName.lname = testArray[1]
		myslc = append(myslc, currentName)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for _, v := range myslc{
		fmt.Printf("FirstName: %s LastName: %s\n", v.fname, v.lname)
	}

}


