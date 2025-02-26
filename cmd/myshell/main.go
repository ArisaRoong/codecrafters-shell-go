package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"slices"
	"strings"
)

func GetBuiltIn() []string {
	return []string{"echo", "exit", "type", "pwd"}
}

func TypeExceptions() []string {
	return []string{"echo", "pwd"}
}

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
			if valid && !slices.Contains(GetBuiltIn(), val) {
				fmt.Fprintf(os.Stdout, "%s is %s\n", val, path)
			} else {
				TypeCommand(val)
			}
		// Echo command
		case "echo":
			// Check for single quotes
			if strings.HasPrefix(val, "'") && strings.HasSuffix(val, "'") {
				outputSlice := strings.Split(val, "'")
				var output string
				for _, word := range outputSlice {
					if word != "" {
						output += word
						val = output
					}
				}
			} else {
				val = strings.Join(strings.Fields(val), " ")
			}
			fmt.Fprintf(os.Stdout, "%s\n", val)
		// Print working directory
		case "pwd":
			dir, err := os.Getwd()
			if err != nil {
				fmt.Errorf("Error processing %s command. Error: %v\n", cmd, err)
			}
			fmt.Fprintf(os.Stdout, "%s\n", dir)
		case "cd":
			// Navigate to home directory
			if val == "~" {
				hdir, err := os.UserHomeDir()
				if err != nil {
					fmt.Fprintf(os.Stdout, "cd: %s: Error navigating to home directory\n", val)
				}
				val = hdir
			}
			err = os.Chdir(val)
			if err != nil {
				fmt.Fprintf(os.Stdout, "cd: %s: No such file or directory\n", val)
			}
		default:
			// Executable
			args := ParseArguments(val)
			command := exec.Command(cmd, args...)
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
	if slices.Contains(GetBuiltIn(), t) {
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

func ParseArguments(s string) []string {
	var tokens []string
	for {
		start := strings.Index(s, "'")
		if start == -1 {
			tokens = append(tokens, strings.Fields(s)...)
			break
		}
		tokens = append(tokens, strings.Fields(s[:start])...)
		s = s[start+1:]
		end := strings.Index(s, "'")
		token := s[:end]
		tokens = append(tokens, token)
		s = s[end+1:]
	}
	return tokens
}
