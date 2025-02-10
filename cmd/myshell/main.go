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
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		input = strings.Trim(input, "\n")

		switch {
		// Exit condition
		case input == "exit 0":
			os.Exit(0)
		// Echo condition
		case strings.HasPrefix(input, "echo"):
			fmt.Fprintf(os.Stdout, "%s", input)
		// Invalid command
		default:
			fmt.Fprintf(os.Stdout, "%s: command not found", input)
		}

		fmt.Fprintf(os.Stdout, "\n")

	}

}
