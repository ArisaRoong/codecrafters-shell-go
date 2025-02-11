package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func main() {

	for {
		// Uncomment this block to pass the first stage
		fmt.Fprint(os.Stdout, "$ ")
		// Wait for user input
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Errorf("Error processing command: %s", input)
		}

		// Split the input into the command and the value
		input = strings.Trim(input, "\n")
		inputSlice := strings.SplitAfterN(input, " ", 2)
		cmd := inputSlice[0]
		cmd = strings.Trim(cmd, " ")
		var val string
		if len(inputSlice) > 1 {
			val = inputSlice[1]
		}

		switch cmd {
		// Exit command
		case "exit":
			if val == "0" {
				os.Exit(0)
			}
		// Type command
		case "type":
			GetType(val)
		// Echo command
		case "echo":
			fmt.Fprintf(os.Stdout, "%s", val)
		// Invalid command
		default:
			fmt.Fprintf(os.Stdout, "%s: command not found", cmd)
		}

		fmt.Fprintf(os.Stdout, "\n")

	}

}

func GetType(t string) {
	if t == "echo" || t == "exit" || t == "type" {
		fmt.Fprintf(os.Stdout, "%s is a shell builtin", t)
	} else {
		fmt.Fprintf(os.Stdout, "%s: not found", t)
	}
}
