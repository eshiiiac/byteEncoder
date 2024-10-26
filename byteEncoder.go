package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func userInput(prompt string, reader *bufio.Reader) int{
	fmt.Println(prompt)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	
	number, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("invalid input")
	}
	return number
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	prompt := userInput("enter a number: ",reader)
	println(prompt)

}