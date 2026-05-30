package main

//i only added fmt, istg everything else just came
import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	//taking input
	for {
		cwd, _ := os.Getwd()
		fmt.Printf("%s > ", cwd)
		//new line par break karing
		input, err := reader.ReadString('\n')
		if err != nil { //error
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if err = execinput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}

}

func execinput(input string) error {
	input = strings.TrimSuffix(input, "\n")
	//earlier i had strings.Split...but this is better because of reasons like multiple spaces
	args := strings.Fields(input)

	if len(args) == 0 { //no input
		return nil
	}
	switch args[0] {
	case "cd":
		if len(args) < 2 {
			return errors.New("destination path required")
		}

		return os.Chdir(args[1])
	case "exit":
		os.Exit(0)
	}
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	return cmd.Run()
}
