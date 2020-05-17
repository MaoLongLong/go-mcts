package game

const (
	Black   = -1
	White   = 1
	NotOver = iota
)

type DotAndBoxMove struct {
	K, I, J, Color int
}

type DotAndBoxState struct {
	DotAndBoxMove
	Depth      int
	Box        [5][5]int
	Board      [2][6][5]int
	NextToMove int
}

func NewDotAndBoxMove(k, i, j, color int) DotAndBoxMove {
	return DotAndBoxMove{
		K:     k,
		I:     i,
		J:     j,
		Color: color,
	}
}

func NewDotAndBoxState(box [5][5]int, board [2][6][5]int, nextToMove, depth int) DotAndBoxState {
	return DotAndBoxState{
		Depth:      depth,
		Box:        box,
		Board:      board,
		NextToMove: nextToMove,
	}
}

func NewDotAndBoxStateWithMove(box [5][5]int, board [2][6][5]int, nextToMove, depth int, move DotAndBoxMove) DotAndBoxState {
	return DotAndBoxState{
		Depth:         depth,
		DotAndBoxMove: move,
		Box:           box,
		Board:         board,
		NextToMove:    nextToMove,
	}
}

func (t DotAndBoxState) GetCount() (int, int) {
	whiteCount, blackCount := 0, 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if t.Box[i][j] == Black {
				blackCount++
			} else if t.Box[i][j] == White {
				whiteCount++
			}
		}
	}
	return blackCount, whiteCount
}

func (t DotAndBoxState) GameResult(must bool) int {
	blackCount, whiteCount := t.GetCount()
	if !must && whiteCount+blackCount < 25 {
		return NotOver
	}
	return whiteCount - blackCount
}

func (t DotAndBoxState) NeedToTrim() bool {
	blackCount, whiteCount := t.GetCount()
	return t.Depth > 6 || blackCount > 12 || whiteCount > 12
}

func (t DotAndBoxState) IsGameOver() bool {
	return t.GameResult(false) != NotOver
}

func (t DotAndBoxState) IsMoveLegal(move DotAndBoxMove) bool {
	if move.Color != t.NextToMove {
		return false
	}

	if move.K != 0 && move.K != 1 {
		return false
	}

	if move.I < 0 || move.I > 5 {
		return false
	}

	if move.J < 0 || move.J > 4 {
		return false
	}

	return t.Board[move.K][move.I][move.J] == 0
}

func (t DotAndBoxState) Move(move DotAndBoxMove) DotAndBoxState {
	newBoard := t.Board
	newBox := t.Box
	newBoard[move.K][move.I][move.J] = move.Color
	flag := false
	// 横
	if move.K == 0 {
		// 0 1 0
		// 0 0 0
		// 1 0 0
		// 1 1 0
		upCount := 0
		if move.I-1 >= 0 {
			// 上
			if t.Board[0][move.I-1][move.J] != 0 {
				upCount++
			}
			// 左上
			if t.Board[1][move.J][move.I-1] != 0 {
				upCount++
			}
			// 右上
			if t.Board[1][move.J+1][move.I-1] != 0 {
				upCount++
			}
		}
		if upCount == 3 {
			newBox[move.I-1][move.J] = move.Color
			flag = true
		}

		// 0 1 0
		// 0 2 0
		// 1 0 1
		// 1 1 1
		downCount := 0
		// 下
		if move.I+1 < 6 {
			if t.Board[0][move.I+1][move.J] != 0 {
				downCount++
			}
		}
		if move.I < 5 {
			// 左下
			if t.Board[1][move.J][move.I] != 0 {
				downCount++
			}
			// 右下
			if t.Board[1][move.J+1][move.I] != 0 {
				downCount++
			}
		}
		if downCount == 3 {
			newBox[move.I][move.J] = move.Color
			flag = true
		}
	} else {
		// 竖
		// 1 1 0
		// 1 0 0
		// 0 0 0
		// 0 1 0
		leftCount := 0
		if move.I-1 >= 0 {
			// 左
			if t.Board[1][move.I-1][move.J] != 0 {
				leftCount++
			}
			// 左上
			if t.Board[0][move.J][move.I-1] != 0 {
				leftCount++
			}
			if t.Board[0][move.J+1][move.I-1] != 0 {
				leftCount++
			}
		}
		if leftCount == 3 {
			newBox[move.J][move.I-1] = move.Color
			flag = true
		}

		// 1 1 0
		// 1 2 0
		// 0 0 1
		// 0 1 1
		rightCount := 0
		if move.I+1 < 6 {
			if t.Board[1][move.I+1][move.J] != 0 {
				rightCount++
			}
		}
		if move.I < 5 {
			if t.Board[0][move.J][move.I] != 0 {
				rightCount++
			}
			if t.Board[0][move.J+1][move.I] != 0 {
				rightCount++
			}
		}
		if rightCount == 3 {
			newBox[move.J][move.I] = move.Color
			flag = true
		}
	}
	var nextToMove int
	if flag {
		nextToMove = move.Color
	} else {
		if move.Color == Black {
			nextToMove = White
		} else {
			nextToMove = Black
		}
	}
	return NewDotAndBoxStateWithMove(newBox, newBoard, nextToMove, t.Depth+1, move)
}

func (t DotAndBoxState) GetLegalMove() []DotAndBoxMove {
	ret := make([]DotAndBoxMove, 0, 64)
	for k := 0; k < 2; k++ {
		for i := 0; i < 6; i++ {
			for j := 0; j < 5; j++ {
				if t.Board[k][i][j] == 0 {
					ret = append(ret, NewDotAndBoxMove(k, i, j, t.NextToMove))
				}
			}
		}
	}
	return ret
}
