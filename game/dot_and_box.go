package game

const (
	Black   = -1
	White   = 1
	NotOver = iota
)

type DotAndBoxMove struct {
	k, i, j, color int
}

type DotAndBoxState struct {
	Box        [5][5]int
	Board      [2][6][5]int
	NextToMove int
}

func NewDotAndBoxMove(k, i, j, color int) DotAndBoxMove {
	return DotAndBoxMove{
		k:     k,
		i:     i,
		j:     j,
		color: color,
	}
}

func NewDotAndBoxState(box [5][5]int, board [2][6][5]int, nextToMove int) DotAndBoxState {
	return DotAndBoxState{
		Box:        box,
		Board:      board,
		NextToMove: nextToMove,
	}
}

func (t DotAndBoxState) GameResult() int {
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
	if blackCount+whiteCount < 25 {
		return NotOver
	}
	if blackCount > whiteCount {
		return Black
	}
	return White
}

func (t DotAndBoxState) IsGameOver() bool {
	return t.GameResult() != NotOver
}

func (t DotAndBoxState) IsMoveLegal(move DotAndBoxMove) bool {
	if move.color != t.NextToMove {
		return false
	}

	if move.k != 0 && move.k != 1 {
		return false
	}

	if move.i < 0 || move.i > 5 {
		return false
	}

	if move.j < 0 || move.j > 4 {
		return false
	}

	return t.Board[move.k][move.i][move.j] == 0
}

func (t DotAndBoxState) Move(move DotAndBoxMove) DotAndBoxState {
	newBoard := t.Board
	newBox := t.Box
	newBoard[move.k][move.i][move.j] = move.color
	flag := false
	// 横
	if move.k == 0 {
		// 0 1 0
		// 0 0 0
		// 1 0 0
		// 1 1 0
		upCount := 0
		if move.i-1 >= 0 {
			// 上
			if t.Board[0][move.i-1][move.j] != 0 {
				upCount++
			}
			// 左上
			if t.Board[1][move.j][move.i-1] != 0 {
				upCount++
			}
			// 右上
			if t.Board[1][move.j+1][move.i-1] != 0 {
				upCount++
			}
		}
		if upCount == 3 {
			newBox[move.i-1][move.j] = move.color
			flag = true
		}

		// 0 1 0
		// 0 2 0
		// 1 0 1
		// 1 1 1
		downCount := 0
		// 下
		if move.i+1 < 6 {
			if t.Board[0][move.i+1][move.j] != 0 {
				downCount++
			}
		}
		if move.i < 5 {
			// 左下
			if t.Board[1][move.j][move.i] != 0 {
				downCount++
			}
			// 右下
			if t.Board[1][move.j+1][move.i] != 0 {
				downCount++
			}
		}
		if downCount == 3 {
			newBox[move.i][move.j] = move.color
			flag = true
		}
	} else {
		// 竖
		// 1 1 0
		// 1 0 0
		// 0 0 0
		// 0 1 0
		leftCount := 0
		if move.i-1 >= 0 {
			// 左
			if t.Board[1][move.i-1][move.j] != 0 {
				leftCount++
			}
			// 左上
			if t.Board[0][move.j][move.i-1] != 0 {
				leftCount++
			}
			if t.Board[0][move.j+1][move.i-1] != 0 {
				leftCount++
			}
		}
		if leftCount == 3 {
			newBox[move.j][move.i-1] = move.color
			flag = true
		}

		// 1 1 0
		// 1 2 0
		// 0 0 1
		// 0 1 1
		rightCount := 0
		if move.i+1 < 6 {
			if t.Board[1][move.i+1][move.j] != 0 {
				rightCount++
			}
		}
		if move.i < 5 {
			if t.Board[0][move.j][move.i] != 0 {
				rightCount++
			}
			if t.Board[0][move.j+1][move.i] != 0 {
				rightCount++
			}
		}
		if rightCount == 3 {
			newBox[move.j][move.i] = move.color
			flag = true
		}
	}
	var nextToMove int
	if flag {
		nextToMove = move.color
	} else {
		if move.color == Black {
			nextToMove = White
		} else {
			nextToMove = Black
		}
	}
	return NewDotAndBoxState(newBox, newBoard, nextToMove)
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
