package shared

type Message struct {
	Type   string   `json:"type"`
	Player string   `json:"player,omitempty"`
	Move   int      `json:"move,omitempty"`
	Board  []string `json:"board,omitempty"`
	Result string   `json:"result,omitempty"`
	Error  string   `json:"error,omitempty"`
}

func NewBoard() []string {
	board := make([]string, 9)
	for i := range board {
		board[i] = " "
	}
	return board
}

func CheckWinner(board []string) string {
	winningCombos := [][]int{
		{0, 1, 2}, {3, 4, 5}, {6, 7, 8},
		{0, 3, 6}, {1, 4, 7}, {2, 5, 8},
		{0, 4, 8}, {2, 4, 6},
	}

	for _, combo := range winningCombos {
		a, b, c := combo[0], combo[1], combo[2]
		if board[a] != " " && board[a] == board[b] && board[b] == board[c] {
			return board[a]
		}
	}
	return ""
}

func IsDraw(board []string) bool {
	for _, v := range board {
		if v == " " {
			return false
		}
	}
	return true
}
