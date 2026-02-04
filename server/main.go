package main

import (
	"encoding/json"
	"fmt"
	"net"
	"sync"

	"tic-tac-toe/shared"
)

type Client struct {
	conn   net.Conn
	player string
	enc    *json.Encoder
	dec    *json.Decoder
}

var waitingClients = make(chan *Client, 100)

func main() {
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}
	fmt.Println("Server running on :9000")

	go matchMaker()

	for {
		conn, _ := listener.Accept()
		client := &Client{
			conn: conn,
			enc:  json.NewEncoder(conn),
			dec:  json.NewDecoder(conn),
		}
		waitingClients <- client
	}
}

func matchMaker() {
	for {
		p1 := <-waitingClients
		p2 := <-waitingClients
		go startGame(p1, p2)
	}
}

func startGame(p1, p2 *Client) {
	board := shared.NewBoard()
	turn := "X"

	p1.player = "X"
	p2.player = "O"

	p1.enc.Encode(shared.Message{Type: "start", Player: "X"})
	p2.enc.Encode(shared.Message{Type: "start", Player: "O"})

	players := map[string]*Client{
		"X": p1,
		"O": p2,
	}

	var mu sync.Mutex

	for {
		current := players[turn]

		var msg shared.Message
		err := current.dec.Decode(&msg)
		if err != nil {
			return
		}

		mu.Lock()

		if msg.Type != "move" || msg.Move < 0 || msg.Move > 8 || board[msg.Move] != " " {
			current.enc.Encode(shared.Message{
				Type:  "error",
				Error: "Invalid move",
			})
			mu.Unlock()
			continue
		}

		board[msg.Move] = turn

		// Send board update
		for _, c := range players {
			c.enc.Encode(shared.Message{
				Type:  "update",
				Board: board,
			})
		}

		// Check win
		winner := shared.CheckWinner(board)
		if winner != "" {
			for _, c := range players {
				c.enc.Encode(shared.Message{
					Type:   "end",
					Result: winner + " wins",
				})
			}
			mu.Unlock()
			return
		}

		// Check draw
		if shared.IsDraw(board) {
			for _, c := range players {
				c.enc.Encode(shared.Message{
					Type:   "end",
					Result: "Draw",
				})
			}
			mu.Unlock()
			return
		}

		// Switch turn
		if turn == "X" {
			turn = "O"
		} else {
			turn = "X"
		}

		mu.Unlock()
	}
}
