package game

//Player represents black and white player
type Player uint8

//IOTA to represent the player
const (
	Empty Player = iota
	White Player = iota
	Black Player = iota
)
