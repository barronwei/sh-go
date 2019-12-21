package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("[sh-go] ")
		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		input = strings.TrimSpace(input)
		args := strings.Fields(input)

		switch args[0] {
		case "cd":
			if len(args) > 1 {
				os.Chdir(args[1])
			}

		default:
			cmd := exec.Command(args[0], args[1:]...)

			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			cmd.Run()
		}
	}
}
