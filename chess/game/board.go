package game

var board [8][8]Player

func canMove(p Player, pos BoardPos) bool {
	return board[pos.row][pos.col] != p
}
