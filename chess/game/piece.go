package game

import (
	"strings"
)

var knightMoves = []BoardPos{
	{Row: 2, Col: 1},
	{Row: 2, Col: -1},
	{Row: -2, Col: 1},
	{Row: -2, Col: -1},
	{Row: 1, Col: 2},
	{Row: -1, Col: 2},
	{Row: 1, Col: -2},
	{Row: -1, Col: -2},
}

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
	boardpos  BoardPos
	sign      string
	player    Player
	firstMove bool
}

type BoardPos struct {
	Row int
	Col int
}

func createPiece(sign string, row int, col int) Piece {
	var player Player
	if sign == strings.ToUpper(sign) {
		player = Black
	} else {
		player = White
	}

	return Piece{
		boardpos: BoardPos{
			Row: row,
			Col: col,
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
	for i:= 0; i< boardSize; i++{
		for j:=0; j<boardSize; j++{
			
			if(i==p.boardpos.Row &&j==p.boardpos.Col){
				continue
			} else if (){
				
			}

		}
	}
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
	for i := p.boardpos.Col; i > 0; i-- {
		for j := p.boardpos.Row; j > 0; j-- {
			pos := BoardPos{
				Row: p.boardpos.Row,
				Col: p.boardpos.Col,
			}
			if board[int(p.boardpos.Row)][int(p.boardpos.Col)] == Empty {
				moves = append(moves, pos)
			} else if canMove(p.player, pos) {
				moves = append(moves, pos)
				break
			} else {
				break
			}
		}
	}

	for i := p.boardpos.Col; i < boardSize; i++ {
		for j := p.boardpos.Row; j < boardSize; j++ {
			pos := BoardPos{
				Row: p.boardpos.Row,
				Col: p.boardpos.Col,
			}
			if board[int(p.boardpos.Row)][int(p.boardpos.Col)] == Empty {
				moves = append(moves, pos)
			} else if canMove(p.player, pos) {
				moves = append(moves, pos)
				break
			} else {
				break
			}
		}
	}
	for i := p.boardpos.Col; i > 0; i-- {
		for j := p.boardpos.Row; j < boardSize; j++ {
			pos := BoardPos{
				Row: p.boardpos.Row,
				Col: p.boardpos.Col,
			}
			if board[int(p.boardpos.Row)][int(p.boardpos.Col)] == Empty {
				moves = append(moves, pos)
			} else if canMove(p.player, pos) {
				moves = append(moves, pos)
				break
			} else {
				break
			}
		}
	}
	for i := p.boardpos.Col; i < boardSize; i++ {
		for j := p.boardpos.Row; j > 0; j-- {
			pos := BoardPos{
				Row: p.boardpos.Row,
				Col: p.boardpos.Col,
			}
			if board[int(p.boardpos.Row)][int(p.boardpos.Col)] == Empty {
				moves = append(moves, pos)
			} else if canMove(p.player, pos) {
				moves = append(moves, pos)
				break
			} else {
				break
			}
		}
	}
	return moves
}

func (p *Piece) getKnightMoves() []BoardPos {
	var moves []BoardPos
	for _, pos := range knightMoves {
		pos := BoardPos{
			Row: p.boardpos.Row + pos.Row,
			Col: p.boardpos.Col + pos.Col,
		}
		if (pos.Row < 8 && pos.Row > 0) && (pos.Col > 0 && pos.Col < 8) {
			if board[pos.Row][pos.Row] == Empty {
				moves = append(moves, pos)
			} else if canMove(p.player, pos) {
				moves = append(moves, pos)
			}
		}
	}
	return moves
}

func (p *Piece) getRookMoves() []BoardPos {
	var moves []BoardPos
	//Move to the left
	for i := p.boardpos.Col; i > 0; i-- {
		pos := BoardPos{
			Row: p.boardpos.Row,
			Col: i,
		}
		if board[int(p.boardpos.Row)][i] == Empty {
			moves = append(moves, pos)
		} else if canMove(p.player, pos) {
			moves = append(moves, pos)
			break
		} else {
			break
		}
	}
	//Move to the right
	for i := p.boardpos.Col; i <= boardSize; i++ {
		pos := BoardPos{
			Row: p.boardpos.Row,
			Col: i,
		}
		if board[int(p.boardpos.Row)][i] == Empty {
			moves = append(moves, pos)
		} else if canMove(p.player, pos) {
			moves = append(moves, pos)
			break
		} else {
			break
		}
	}
	//Move up
	for i := p.boardpos.Row; i <= boardSize; i++ {
		pos := BoardPos{
			Row: i,
			Col: p.boardpos.Col,
		}
		if board[i][int(p.boardpos.Col)] == Empty {
			moves = append(moves, pos)
		} else if canMove(p.player, pos) {
			moves = append(moves, pos)
			break
		} else {
			break
		}
	}
	//Move down
	for i := p.boardpos.Row; i > 0; i-- {
		pos := BoardPos{
			Row: i,
			Col: p.boardpos.Col,
		}
		if board[i][int(p.boardpos.Col)] == Empty {
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
	//Possible moves for white pawn:
	//w,b = current pos of white/black pawn, x = possible, 0 = impossible
	//0 x 0  |  0 b 0
	//x x x  |  x x x
	//0 w 0  |  0 x 0
	var moves []BoardPos
	var val int
	if p.player == White {
		val = 1
	} else if p.player == Black {
		val = -1
	}

	if board[p.boardpos.Row+val][p.boardpos.Col] == Empty {
		moves = append(moves, BoardPos{p.boardpos.Row + val, p.boardpos.Col})
		if (board[p.boardpos.Row+2*val][p.boardpos.Col] == Empty) && p.firstMove {
			moves = append(moves, BoardPos{p.boardpos.Row + 2*val, p.boardpos.Col})
		}
	}

	if canMove(p.player, BoardPos{p.boardpos.Row + val, p.boardpos.Col + 1}) {
		moves = append(moves, BoardPos{p.boardpos.Row + val, p.boardpos.Col + 1})
	}

	if canMove(p.player, BoardPos{p.boardpos.Row + val, p.boardpos.Col - 1}) {
		moves = append(moves, BoardPos{p.boardpos.Row + val, p.boardpos.Col - 1})
	}

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
