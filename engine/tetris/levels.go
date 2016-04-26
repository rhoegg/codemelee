package tetris

var Levels = [...]Level{
	Level{number: 1, speed: 1,
		NextPiece: func() TetrisPiece {
			return Pieces.O
		},
		Score: func(lines int) int {
			return lines * lines
		}}}