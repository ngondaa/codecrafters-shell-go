package main

import (
	"fmt"
	"os"
	"bufio"
)

var _ = fmt.Print

func main() {
	
	fmt.Print("$ ")

	// Wait for user input
	command , err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println(os.Stderr, "Error reading input:", err)
		os.Exit(1)
	}

	// Remove the newline character
	command = strings.TrimSpace(command)
	fmt.Println(command[:len(command)-1] + ": command not found")

	// Execute the command
}
