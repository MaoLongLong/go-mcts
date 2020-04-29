package mcts

import (
	"math"
	"math/rand"
	"tictactoe/game"
)

type MonteCarloTreeSearchNode struct {
	NumberOfVis int
	Results     map[int]int
	State       game.TicTacToeGameState
	Parent      *MonteCarloTreeSearchNode
	Children    []*MonteCarloTreeSearchNode
	UntriedMove []game.TicTacToeMove
}

func NewNode(state game.TicTacToeGameState, parent *MonteCarloTreeSearchNode) *MonteCarloTreeSearchNode {
	return &MonteCarloTreeSearchNode{
		NumberOfVis: 0,
		Results:     make(map[int]int),
		State:       state,
		Parent:      parent,
		Children:    make([]*MonteCarloTreeSearchNode, 0),
		UntriedMove: state.GetLegalMove(),
	}
}

// 评估节点
func (n *MonteCarloTreeSearchNode) Q() float64 {
	wins := n.Results[n.Parent.State.NextToMove]
	loses := n.Results[-1*n.Parent.State.NextToMove]
	return float64(wins - loses) // 净胜场数
}

func (n *MonteCarloTreeSearchNode) N() float64 {
	return float64(n.NumberOfVis)
}

func (n *MonteCarloTreeSearchNode) Expand() *MonteCarloTreeSearchNode {
	move := n.UntriedMove[len(n.UntriedMove)-1]
	n.UntriedMove = n.UntriedMove[:len(n.UntriedMove)-1]
	nextState := n.State.Move(move)
	childNode := NewNode(nextState, n)
	n.Children = append(n.Children, childNode)
	return childNode
}

func (n *MonteCarloTreeSearchNode) IsTerminalNode() bool {
	return n.State.IsGameOver()
}

func (n *MonteCarloTreeSearchNode) RollOut() int {
	cur := game.NewTicTacToeGameState(n.State.Board, n.State.NextToMove)
	for !cur.IsGameOver() {
		moves := cur.GetLegalMove()
		move := moves[rand.Intn(len(moves))]
		//r := rand.Intn(len(moves))
		//move := moves[r]
		//moves = append(moves[:r], moves[r+1:]...)
		cur = cur.Move(move)
	}
	return cur.GameResult()
}

func (n *MonteCarloTreeSearchNode) Back(result int) {
	n.NumberOfVis++
	n.Results[result]++
	if n.Parent != nil {
		n.Parent.Back(result)
	}
}

func (n *MonteCarloTreeSearchNode) IsFullyExpanded() bool {
	return len(n.UntriedMove) == 0
}

func (n *MonteCarloTreeSearchNode) BestChild(param float64) *MonteCarloTreeSearchNode {
	max, child := -float64(0x3f3f3f3f), new(MonteCarloTreeSearchNode)
	for _, c := range n.Children {
		x := c.Q()/c.N() + param*math.Sqrt(2*math.Log(n.N())/c.N())
		if x > max {
			max = x
			child = c
		}
	}
	return child
}
