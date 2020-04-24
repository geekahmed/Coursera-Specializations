package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main()  {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter a String: ")
	scanner.Scan()
	str := scanner.Text()
	newStr := strings.ToLower(str)
	if strings.Index(newStr, "i") == 0   && strings.Contains(newStr, "a") && strings.HasSuffix(newStr, "n") {
		fmt.Printf("Found!")
	} else {
		fmt.Printf("Not Found!")
	}

}
