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
	//timeout := time.After(10 * time.Second)
	//quit := make(chan struct{})
	//go func() {
	//	for {
	//		select {
	//		case <-timeout:
	//			quit <- struct{}{}
	//			return
	//		default:
	//			v := s.TreePolicy()
	//			reward := v.RollOut()
	//			v.Back(reward)
	//		}
	//	}
	//}()
	//<-quit
	for i := 0; i < 500000; i++ {
		v := s.TreePolicy()
		res := v.RollOut()
		v.Back(res)
	}
	return s.root.BestChild(float64(0))
}
