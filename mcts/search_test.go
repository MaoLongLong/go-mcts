package mcts

import (
	"testing"
	"tictactoe/game"
)

func TestBestMove(t *testing.T) {
	board := [3][3]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}
	state := game.NewTicTacToeGameState(board, game.X)
	root := NewNode(state, nil)
	search := NewSearch(root)
	node := search.BestMove()
	t.Log(node.State.Board)
}

func BenchmarkBestMove(b *testing.B) {
	board := [3][3]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}
	state := game.NewTicTacToeGameState(board, game.X)
	root := NewNode(state, nil)
	search := NewSearch(root)
	node := search.BestMove()
	b.Log(node.State.Board)
}
