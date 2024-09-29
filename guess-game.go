package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var input, err = reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.Trimspaces(input)

}
