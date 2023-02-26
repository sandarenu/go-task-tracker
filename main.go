package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("=== Welcome to Task Tracker ===")
	fmt.Println("Usage: help, list, add, complete, exit")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(">> ")
		enteredCmd, _ := reader.ReadString('\n')
		cmd := strings.TrimSpace(enteredCmd)
		fmt.Println("You entered: ", cmd)

		i := strings.Index(cmd, " ")
		if i > 0 {
			c := cmd[0:i]
			d := cmd[i:len(cmd)]

			vc, e := validateCmd(c)
			if e != nil {
				fmt.Println("Invalid command: ", e)
				continue
			}

			fmt.Printf("cmd[%s] data[%s]\n", vc, strings.TrimSpace(d))

		} else {
			vc, e := validateCmd(cmd)
			if e != nil {
				fmt.Println("Invalid command: ", e)
				continue
			}

			fmt.Printf("cmd[%s]\n", vc)
		}

		if cmd == "exit" {
			fmt.Println("bye...")
			os.Exit(0)
		}

		if cmd == "help" {
			printHelp()
		}
	}
}

func validateCmd(cmd string) (string, error) {
	validCommands := []string{"help", "list", "add", "complete", "exit"}
	for _, v := range validCommands {
		if v == cmd {
			return cmd, nil
		}
	}
	return "", errors.New("invalid Cmd")
}

func printHelp() {
	fmt.Println("Usage: list, add, complete, exit")
}
