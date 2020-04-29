package mcts

type MonteCarloTreeSearch struct {
	root *MonteCarloTreeSearchNode
}

func NewSearch(root *MonteCarloTreeSearchNode) MonteCarloTreeSearch {
	return MonteCarloTreeSearch{root}
}

func (s MonteCarloTreeSearch) TreePolicy() *MonteCarloTreeSearchNode {
	cur := s.root
	for !cur.IsTerminalNode() {
		if !cur.IsFullyExpanded() {
			return cur.Expand()
		} else {
			cur = cur.BestChild(1.44)
		}
	}
	return cur
}

func (s MonteCarloTreeSearch) BestMove() *MonteCarloTreeSearchNode {
	for i := 0; i < 100000; i++ {
		v := s.TreePolicy()
		reward := v.RollOut()
		v.Back(reward)
	}
	return s.root.BestChild(float64(0))
}
