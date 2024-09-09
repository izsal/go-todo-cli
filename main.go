package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/izsal/go-todo-cli/db"
	"github.com/izsal/go-todo-cli/model"
	"github.com/izsal/go-todo-cli/service"
)

func main() {
	log.Print("Start the todo app...")

	ser := service.NewService(db.New())

	scanner := bufio.NewScanner(os.Stdin)

	for {
		// to show command when user input
		fmt.Print("> ")
		scanner.Scan()
		cmd := strings.Split(scanner.Text(), " ")

		if len(cmd) == 0 || cmd[0] == "" {
			fmt.Println("Please enter a command.")
			continue
		}

		switch strings.ToLower(cmd[0]) {
		case "add":
			if len(cmd) < 2 {
				fmt.Println("Please provide a label for the todo.")
				continue
			}
			ser.Add(&model.ToDo{Label: strings.Join(cmd[1:], " ")})

		case "delete":
			if len(cmd) < 2 {
				fmt.Println("Please provide an ID to delete.")
				continue
			}
			id, err := strconv.Atoi(cmd[1])
			if err != nil {
				fmt.Printf("Please provide a valid ID - %v\n", cmd[1])
				continue
			}
			if err := ser.Delete(id); err != nil {
				fmt.Printf("ID %d not found\n", id)
			}

		case "update":
			if len(cmd) < 3 {
				fmt.Println("Please provide an ID and a new label to update.")
				continue
			}
			id, err := strconv.Atoi(cmd[1])
			if err != nil {
				fmt.Printf("Please provide a valid ID - %v\n", cmd[1])
				continue
			}
			ser.Update(&model.ToDo{Id: id, Label: strings.Join(cmd[2:], " ")})

		case "get":
			if len(cmd) < 2 {
				fmt.Println("Please provide an ID to get the todo item.")
				continue
			}
			id, err := strconv.Atoi(cmd[1])
			if err != nil {
				fmt.Printf("Please provide a valid ID - %v\n", cmd[1])
				continue
			}
			todo := ser.GetById(id)
			if todo == nil {
				fmt.Printf("Todo with ID %d not found\n", id)
			} else {
				fmt.Printf("%d    %s\n", todo.Id, todo.Label)
			}

		case "getall":
			todos := ser.GetAll()
			if len(todos) == 0 {
				fmt.Println("No todos available")
				continue
			}
			for _, todo := range todos {
				fmt.Printf("%d    %s\n", todo.Id, todo.Label)
			}

		default:
			fmt.Printf("Invalid operation: %s\n", cmd[0])
		}
	}
}
