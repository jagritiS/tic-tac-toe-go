package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"

	"tic-tac-toe/shared"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:9000")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	enc := json.NewEncoder(conn)
	dec := json.NewDecoder(conn)

	reader := bufio.NewReader(os.Stdin)

	var myPlayer string

	go func() {
		for {
			var msg shared.Message
			err := dec.Decode(&msg)
			if err != nil {
				os.Exit(0)
			}

			switch msg.Type {
			case "start":
				myPlayer = msg.Player
				fmt.Println("Game started. You are:", myPlayer)

			case "update":
				printBoard(msg.Board)

			case "end":
				fmt.Println("Game over:", msg.Result)
				os.Exit(0)

			case "error":
				fmt.Println("Error:", msg.Error)
			}
		}
	}()

	for {
		fmt.Print("Enter move (0-8): ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		move, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println("Enter a number between 0 and 8")
			continue
		}

		enc.Encode(shared.Message{
			Type: "move",
			Move: move,
		})
	}
}

func printBoard(board []string) {
	fmt.Println()
	for i := 0; i < 9; i += 3 {
		fmt.Printf(" %s | %s | %s \n", board[i], board[i+1], board[i+2])
		if i < 6 {
			fmt.Println("---+---+---")
		}
	}
	fmt.Println()
}
