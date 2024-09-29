package main

import (
	"bufio"
	"fmt" // log messages to stdout with
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter a Guess \n")
	var reader = bufio.NewReader(os.Stdin)
	var input, err = reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	// to trim \n and other spaces like \t
	input = strings.TrimSpace(input)
	fmt.Println("User Input", input, reflect.TypeOf(input))
	//parsedInput, err := strconv.Atoi(input) // to reassign atoi error back to err variable
	// err is overriden
	//if err != nil {
	//	log.Fatal(err)

	//}
	var parsedInput, parsedInputErr = strconv.Atoi(input)
	if parsedInputErr != nil {
		log.Fatal("Invalid Input: ", parsedInputErr)
	}
	fmt.Println("TypeConverted", input, reflect.TypeOf(parsedInput))

	var isEligible = false
	if parsedInput > 18 {
		isEligible = true
	}
	fmt.Println("Result", isEligible)
	//return isEligible;
}
