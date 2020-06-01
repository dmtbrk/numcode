package main

import (
	"fmt"
	"os"

	"github.com/ortymid/numcode/pkg/numcode"
)

func main() {
	if len(os.Args) == 3 {
		switch os.Args[1] {
		case "-e":
			fmt.Println(numcode.Encode(os.Args[2]))
		case "-d":
			fmt.Println(numcode.Decode(os.Args[2]))
		default:
			printHelp()
		}
	} else {
		printHelp()
	}
}

func printHelp() {
	fmt.Println("numcode 0.1.0")
	fmt.Println("Usage: numcode [-e | -d] string")
	fmt.Println("\t-e\tspecifies that the following string should be encoded")
	fmt.Println("\t-d\tspecifies that the following string should be decoded")
}
