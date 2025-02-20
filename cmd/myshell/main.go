package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func main() {

	for {
		fmt.Fprint(os.Stdout, "$ ")
		// Wait for user input
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Errorf("Error processing command: %s\n", input)
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
			path, valid := IsValidCommand(val)
			if valid && val != "echo" {
				fmt.Fprintf(os.Stdout, "%s is %s\n", val, path)
			} else {
				TypeCommand(val)
			}
		// Echo command
		case "echo":
			fmt.Fprintf(os.Stdout, "%s\n", val)
		default:
			// Executable
			command := exec.Command(cmd, val)
			command.Stderr = os.Stderr
			command.Stdout = os.Stdout
			err := command.Run()
			if err != nil {
				fmt.Fprintf(os.Stdout, "%s: command not found\n", cmd)
			}
		}

	}

}

func TypeCommand(t string) {
	if t == "echo" || t == "exit" || t == "type" {
		fmt.Fprintf(os.Stdout, "%s is a shell builtin\n", t)
	} else {
		fmt.Fprintf(os.Stdout, "%s: not found\n", t)
	}
}

func IsValidCommand(input string) (string, bool) {
	paths := strings.Split(os.Getenv("PATH"), ":")
	for _, path := range paths {
		fullPath := path + "/" + input
		if _, err := os.Stat(fullPath); err == nil {
			return fullPath, true
		}
	}
	return "", false
}
