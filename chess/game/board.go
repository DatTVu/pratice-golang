package game

var board [8][8]Player

const boardSize int = 8

func canMove(p Player, pos BoardPos) bool {
	return board[pos.Row][pos.Col] != p
}
