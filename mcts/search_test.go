package mcts

import (
	"math/rand"
	"runtime"
	"testing"
	"tictactoe/game"
	"time"
)

func TestBestMove(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	rand.Seed(time.Now().UnixNano())
	box := [5][5]int{
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}
	board := [2][6][5]int{
		{
			{-1, 0, 0, 0, 0},
			{1, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
		},
		{
			{-1, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
		},
	}
	state := game.NewDotAndBoxState(box, board, game.White)
	root := NewNode(state, nil)
	search := NewSearch(root)
	node := search.BestMove()
	t.Log(node.State.Board)
}

//func BenchmarkBestMove(b *testing.B) {
//	board := [3][3]int{
//		{0, 0, 0},
//		{0, 0, 0},
//		{0, 0, 0},
//	}
//	state := game.NewTicTacToeGameState(board, game.X)
//	root := NewNode(state, nil)
//	search := NewSearch(root)
//	node := search.BestMove()
//	b.Log(node.State.Board)
//}
