package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func execinput(input string) error {
	input = strings.TrimSuffix(input, "\n")
	args := strings.Split(input, " ")
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		if err = execinput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}

}
