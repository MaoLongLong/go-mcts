package game

const (
	X = -1
	XIsWin
	O = 1
	OIsWin
	Draw = iota
	Continue
)

type TicTacToeMove struct {
	x     int
	y     int
	value int
}

type TicTacToeGameState struct {
	Board      [3][3]int
	NextToMove int
}

func NewTicTacToeMove(x, y, value int) TicTacToeMove {
	return TicTacToeMove{
		x:     x,
		y:     y,
		value: value,
	}
}

func NewTicTacToeGameState(board [3][3]int, nextToMove int) TicTacToeGameState {
	return TicTacToeGameState{
		Board:      board,
		NextToMove: nextToMove,
	}
}

func (t TicTacToeGameState) GameResult() int {
	if t.Board[0][0] != 0 && t.Board[0][0] == t.Board[0][1] && t.Board[0][1] == t.Board[0][2] {
		if t.Board[0][0] == X {
			return XIsWin
		}
		return OIsWin
	}
	if t.Board[1][0] != 0 && t.Board[1][0] == t.Board[1][1] && t.Board[1][1] == t.Board[1][2] {
		if t.Board[1][0] == X {
			return XIsWin
		}
		return OIsWin
	}
	if t.Board[2][0] != 0 && t.Board[2][0] == t.Board[2][1] && t.Board[2][1] == t.Board[2][2] {
		if t.Board[2][0] == X {
			return XIsWin
		}
		return OIsWin
	}
	if t.Board[0][0] != 0 && t.Board[0][0] == t.Board[1][0] && t.Board[1][0] == t.Board[2][0] {
		if t.Board[0][0] == X {
			return XIsWin
		}
		return OIsWin
	}
	if t.Board[0][1] != 0 && t.Board[0][1] == t.Board[1][1] && t.Board[1][1] == t.Board[2][1] {
		if t.Board[0][1] == X {
			return XIsWin
		}
		return OIsWin
	}
	if t.Board[0][2] != 0 && t.Board[0][2] == t.Board[1][2] && t.Board[1][2] == t.Board[2][2] {
		if t.Board[0][2] == X {
			return XIsWin
		}
		return OIsWin
	}
	if t.Board[0][0] != 0 && t.Board[0][0] == t.Board[1][1] && t.Board[1][1] == t.Board[2][2] {
		if t.Board[0][0] == X {
			return XIsWin
		}
		return OIsWin
	}
	if t.Board[0][2] != 0 && t.Board[0][2] == t.Board[1][1] && t.Board[1][1] == t.Board[2][0] {
		if t.Board[0][2] == X {
			return XIsWin
		}
		return OIsWin
	}

	for _, row := range t.Board {
		for _, x := range row {
			if x == 0 {
				return Continue
			}
		}
	}

	return Draw
}

func (t TicTacToeGameState) IsGameOver() bool {
	return t.GameResult() != Continue
}

func (t TicTacToeGameState) IsMoveLegal(move TicTacToeMove) bool {
	if move.value != t.NextToMove {
		return false
	}

	if move.x < 0 || move.x > 2 {
		return false
	}

	if move.y < 0 || move.y > 2 {
		return false
	}

	return t.Board[move.x][move.y] == 0
}

func (t TicTacToeGameState) Move(move TicTacToeMove) TicTacToeGameState {
	newBoard := t.Board
	newBoard[move.x][move.y] = move.value
	var nextToMove int
	if t.NextToMove == O {
		nextToMove = X
	} else {
		nextToMove = O
	}
	return NewTicTacToeGameState(newBoard, nextToMove)
}

func (t TicTacToeGameState) GetLegalMove() []TicTacToeMove {
	ret := make([]TicTacToeMove, 0)
	for x, row := range t.Board {
		for y, val := range row {
			if val == 0 {
				ret = append(ret, NewTicTacToeMove(x, y, t.NextToMove))
			}
		}
	}
	return ret
}
