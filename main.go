package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	connectToDb()

	fmt.Println("=== Welcome to Task Tracker ===")
	fmt.Println("Usage: help, list, add, complete, exit")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(">> ")
		enteredCmd, _ := reader.ReadString('\n')
		cmd, err := buildCmdData(enteredCmd)
		if err != nil {
			fmt.Println("Invalid command: ", err)
			continue
		}

		if cmd.Cmd == "exit" {
			fmt.Println("bye...")
			os.Exit(0)
		}

		if cmd.Cmd == "help" {
			printHelp()
		}

		if cmd.Cmd == "list" {
			var tasks []Task
			var err error

			switch cmd.Data {
			case "all":
				tasks, err = readTasks()
			default:
				tasks, err = readPending()
			}

			if err != nil {
				log.Fatal("Error reading tasks", err)
			}
			printTasks(tasks)
		}

		if cmd.Cmd == "add" {
			task := Task{Title: cmd.Data}
			id, err := addTask(task)
			if err != nil {
				log.Fatalf("Error while inserting new task [%v]\n", err)
			}
			fmt.Printf("New task[%d] inserted\n", id)
		}

		if cmd.Cmd == "complete" {
			id, err := markComplete(cmd.Data)
			if err != nil {
				fmt.Printf("Error updating task [%v]\n", err)
			} else {

				if id > 0 {
					fmt.Printf("Task [%s] updated\n", cmd.Data)
				} else {
					fmt.Printf("No task to update\n")
				}
			}
		}
	}
}

func printTasks(tasks []Task) {
	for _, t := range tasks {
		fmt.Printf("%d) %s (%s)\n", t.TaskId, t.Title, statusStr(t.Status))
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

func buildCmdData(str string) (CmdData, error) {
	cmd := strings.TrimSpace(str)
	//fmt.Println("You entered: ", cmd)

	i := strings.Index(cmd, " ")
	if i > 0 {
		c := cmd[0:i]
		d := cmd[i:len(cmd)]

		vc, e := validateCmd(c)
		if e != nil {
			fmt.Println("Invalid command: ", e)
			return CmdData{}, fmt.Errorf("Invalid command [%s]", cmd)
		}

		return CmdData{Cmd: vc, Data: strings.TrimSpace(d)}, nil

	} else {
		vc, e := validateCmd(cmd)
		if e != nil {
			fmt.Println("Invalid command: ", e)
			return CmdData{}, fmt.Errorf("Invalid command [%s]", cmd)
		}

		return CmdData{Cmd: vc, Data: ""}, nil
	}
}

func statusStr(s TaskStatus) string {
	switch s {
	case Pending:
		return "pending"
	case Completed:
		return "done"
	default:
		return ""
	}
}
