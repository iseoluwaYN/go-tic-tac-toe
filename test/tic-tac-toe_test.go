package test

import (
	"github.com/iseoluwaYN/tic-tac-toe/game"
	"github.com/stretchr/testify/assert"
	"testing"
)

var tictactoe game.TicTacToe
var gameBoard [3][3]game.Value

func beforeEach() {
	tictactoe = game.NewTicTacToe()
	gameBoard = tictactoe.Board

}
func TestThatGameExist(t *testing.T) {
	//given
	tictactoe := game.NewTicTacToe()
	assert.NotNil(t, tictactoe)
}

func TestThatBoardExists(t *testing.T) {
	//given
	beforeEach()
	assert.NotNil(t, tictactoe)
	assert.NotNil(t, gameBoard)
}

func TestThatBoardCellIsEmptyByDefault(t *testing.T) {
	//given
	beforeEach()
	assert.NotNil(t, tictactoe)
	assert.NotNil(t, gameBoard)
	for i := 0; i < len(gameBoard); i++ {
		for j := i; j < i; j++ {
			var isEmpty bool
			valueInCell := gameBoard[i][j]
			if valueInCell == "EMPTY" {
				isEmpty = true
			}
			assert.True(t, isEmpty, valueInCell)
			//assert.Equal(t, "EMPTY", valueInCell)
		}
	}
}

func TestThatThereAreOnlyTwoPlayers(t *testing.T) {
	beforeEach()
	assert.NotNil(t, tictactoe)
	assert.NotNil(t, gameBoard)
	players := tictactoe.Players

	for i := 0; i < len(players); i++ {
		player := players
		assert.Equal(t, 2, len(player))
	}

}
