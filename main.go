package main

import (
	"fmt"
	game2 "github.com/iseoluwaYN/tic-tac-toe/game"
)

func main() {
	game := game2.NewTicTacToe()
	fmt.Println(game.Board)
}
