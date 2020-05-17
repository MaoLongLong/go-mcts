package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"strings"
	"tictactoe/game"
	"tictactoe/mcts"
	"time"
)

var (
	reader  *bufio.Reader
	side    string
	num     int
	lines   string
	human   int
	ai      int
	message string
	logFile *os.File
	state   game.DotAndBoxState
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	rand.Seed(time.Now().UnixNano())
	logFile, _ = os.Create(fmt.Sprintf("./log/%v.txt", time.Now().Unix()))
	defer logFile.Close()
	log.SetOutput(logFile)
	reader = bufio.NewReader(os.Stdin)
	for {
		message, _ = reader.ReadString('\n')
		trim(&message)

		log.Println("[SAU] " + message)

		if message == "name?" {
			log.Println("[AI] " + "name hello")
			fmt.Println("name 50")
		}

		if message == "quit" {
			break
		}

		if strings.HasPrefix(message, "new") {
			temp := strings.Split(message, " ")
			side = temp[1]

			if side == "black" {
				ai = game.Black
				human = game.White
				run()
			} else {
				ai = game.White
				human = game.Black
			}
		}

		if strings.HasPrefix(message, "move") {
			temp := strings.Split(message, " ")
			num, _ = strconv.Atoi(temp[1])
			lines = temp[2]

			for t := 0; t < num; t++ {
				k := int(lines[3*t] - 'A')
				i := int(lines[3*t+1] - 'A')
				j := int(lines[3*t+2] - 'A')
				state = state.Move(game.NewDotAndBoxMove(k, i, j, human))
			}
			run()
		}

		if message == "end" {
			log.Printf("[BOX] %v\n", state.Box)
			log.Printf("[BOARD] %v\n", state.Board)

			var newState game.DotAndBoxState
			state = newState
		}

		if message == "error" {
			log.Printf("[ERROR] [BOX] %v\n", state.Box)
			log.Printf("[ERROR] [BOARD] %v\n", state.Board)
		}
	}
}

func trim(s *string) {
	*s = strings.TrimLeft(*s, "\x00")
	*s = strings.TrimRight(*s, "\n")
	*s = strings.TrimRight(*s, "\r")
}

func run() {
	lines := make([]byte, 0)
	state.NextToMove = ai
	state.Depth = 0
	for state.NextToMove == ai && !state.IsGameOver() {
		state.Depth = 0
		root := mcts.NewNode(state, nil)
		search := mcts.NewSearch(root)
		node := search.BestMove()
		log.Printf("[B/W] %d/%d\n", node.Results[game.Black], node.Results[game.White])
		state = node.State
		lines = append(lines, chr(state.K))
		lines = append(lines, chr(state.I))
		lines = append(lines, chr(state.J))
	}
	fmt.Printf("move %d %s\n", len(lines)/3, lines)
	log.Printf("[AI] move %d %s\n", len(lines)/3, lines)
}

func chr(num int) byte {
	return byte('A' + num)
}
