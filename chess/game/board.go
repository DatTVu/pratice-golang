package game

type Board struct {
	board [8][8]Player
}

//var board [8][8]Player

const boardSize int = 8

func (b Board) canMove(p Player, pos BoardPos) bool {
	return b.board[pos.Row][pos.Col] != p
}

func (b *Board) CreateBoard() bool {
	return false
}

func (b *Board) LoadBoardSettings(path string) bool {

	return true
}
