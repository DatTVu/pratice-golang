package game

import (
	"strings"
)

//String representation of piece
const (
	WhiteKing = "\u2654"
	BlackKing = "\u265A"

	WhiteQueen = "\u2655"
	BlackQueen = "\u265B"

	WhiteRook = "\u2656"
	BlackRook = "\u265C"

	WhiteBishop = "\u2657"
	BlackBishop = "\u265D"

	WhiteKnight = "\u2658"
	BlackKnight = "\u265E"

	WhitePawn = "\u2659"
	BlackPawn = "\u265F"
)

type Piece struct {
	boardpos BoardPos
	sign     string
	player   Player
}

type BoardPos struct {
	row uint8
	col uint8
}

func createPiece(sign string, row uint8, col uint8) Piece {
	var player Player
	if sign == strings.ToUpper(sign) {
		player = Black
	} else {
		player = White
	}

	return Piece{
		boardpos: BoardPos{
			row: row,
			col: col,
		},
		sign:   sign,
		player: player,
	}
}

func (p *Piece) getPieceSymbol() string {

	switch p.sign {
	case "k":
		return WhiteKing
	case "K":
		return BlackKing
	case "q":
		return WhiteQueen
	case "Q":
		return BlackQueen
	case "b":
		return WhiteBishop
	case "B":
		return BlackBishop
	case "n":
		return WhiteKnight
	case "N":
		return BlackKnight
	case "r":
		return WhiteRook
	case "R":
		return BlackRook
	case "p":
		return WhitePawn
	case "P":
		return BlackPawn

	default:
		panic("Cannot Print Piece. Unknown Sign:" + p.sign)
	}
}

func (p *Piece) String() string {
	return p.getPieceSymbol()
}

func (p *Piece) getKingMoves() []BoardPos {
	var moves []BoardPos
	//moves.append(moves, )
	return moves
}

func (p *Piece) getQueenMoves() []BoardPos {
	var moves []BoardPos
	moves = append(moves, p.getRookMoves()...)
	moves = append(moves, p.getBishopMoves()...)
	return moves
}

func (p *Piece) getBishopMoves() []BoardPos {
	var moves []BoardPos
	return moves
}

func (p *Piece) getKnightMoves() []BoardPos {
	var moves []BoardPos
	return moves
}

func (p *Piece) getRookMoves() []BoardPos {
	var moves []BoardPos
	//Move to the left
	for i := p.boardpos.col; i > 0; i-- {
		pos := BoardPos{
			row: p.boardpos.row,
			col: i,
		}
		if board[p.boardpos.row][i] == Empty {
			moves = append(moves, pos)
		} else if canMove(p.player, pos) {
			moves = append(moves, pos)
			break
		} else {
			break
		}
	}
	//Move to the right
	for i := p.boardpos.col; i <= 8; i++ {
		pos := BoardPos{
			row: p.boardpos.row,
			col: i,
		}
		if board[p.boardpos.row][i] == Empty {
			moves = append(moves, pos)
		} else if canMove(p.player, pos) {
			moves = append(moves, pos)
			break
		} else {
			break
		}
	}
	//Move up
	for i := p.boardpos.row; i <= 8; i++ {
		pos := BoardPos{
			row: i,
			col: p.boardpos.col,
		}
		if board[i][p.boardpos.col] == Empty {
			moves = append(moves, pos)
		} else if canMove(p.player, pos) {
			moves = append(moves, pos)
			break
		} else {
			break
		}
	}
	//Move down
	for i := p.boardpos.row; i > 0; i-- {
		pos := BoardPos{
			row: i,
			col: p.boardpos.col,
		}
		if board[i][p.boardpos.col] == Empty {
			moves = append(moves, pos)
		} else if canMove(p.player, pos) {
			moves = append(moves, pos)
			break
		} else {
			break
		}
	}
	return moves
}

func (p *Piece) getPawnMoves() []BoardPos {
	var moves []BoardPos
	return moves
}

func (p *Piece) getMoves() []BoardPos {
	switch p.String() {
	case WhiteKing, BlackKing:
		return p.getKingMoves()
	case WhiteQueen, BlackQueen:
		return p.getQueenMoves()
	case WhiteBishop, BlackBishop:
		return p.getBishopMoves()
	case WhiteKnight, BlackKnight:
		return p.getKnightMoves()
	case WhiteRook, BlackRook:
		return p.getRookMoves()
	case WhitePawn, BlackPawn:
		return p.getPawnMoves()
	default:
		return nil
	}
}

func (p *Piece) isKing() bool {
	switch p.String() {
	case WhiteKing, BlackKing:
		return true
	default:
		return false
	}
}

func isCheck() bool {
	return false
}
