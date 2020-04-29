package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime/pprof"
	"tictactoe/game"
	"tictactoe/mcts"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	file, _ := os.Create("cpu.pprof")
	pprof.StartCPUProfile(file)
	defer pprof.StopCPUProfile()
	for i := 0; i < 100; i++ {
		board := [3][3]int{
			{0, 0, 0},
			{0, 0, 0},
			{0, 0, 0},
		}
		state := game.NewTicTacToeGameState(board, game.X)
		root := mcts.NewNode(state, nil)
		search := mcts.NewSearch(root)
		node := search.BestMove()
		fmt.Println(node.State.Board)
	}
}
