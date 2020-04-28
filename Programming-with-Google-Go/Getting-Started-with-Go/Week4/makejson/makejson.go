package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	var My_map = make(map[string]string)

	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Enter a name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSuffix(name, "\n")

	fmt.Printf("Enter an address: ")
	address, _ := reader.ReadString('\n')
	address = strings.TrimSuffix(address, "\n")

	My_map["name"] = name
	My_map["address"] = address

	jsonObject, _ := json.Marshal(My_map)

	fmt.Println(string(jsonObject))


}


