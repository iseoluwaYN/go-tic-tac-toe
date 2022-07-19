package game

import "github.com/iseoluwaYN/tic-tac-toe/player"

type Value string

type TicTacToe struct {
	Board   [3][3]Value
	Players [2]player.Player
}

//constructor
func NewTicTacToe() TicTacToe {
	ticTacToe := TicTacToe{}
	for i := 0; i < len(ticTacToe.Board); i++ {
		for j := 0; j < len(ticTacToe.Board[i]); j++ {
			ticTacToe.Board[i][j] = "EMPTY"
		}
	}

	return ticTacToe
}
