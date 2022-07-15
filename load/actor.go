package load

import (
	"math/rand"
)

type Point struct {
	row int
	col int
	startRow int
	startCol int
}

var Player Point

func MakeMove(oldRow, oldCol int, cmd string) (newRow, newCol int) {
	newRow, newCol = oldRow, oldCol

	switch cmd {
	case "UP":
		newRow = newRow - 1
		if newRow < 0 {
			newRow = len(arr) - 1
		}
	case "DOWN":
		newRow = newRow + 1
		if newRow == len(arr) {
			newRow = 0
		}
	case "RIGHT":
		newCol = newCol + 1
		if newCol == len(arr[0]) {
			newCol = 0
		}
	case "LEFT":
		newCol = newCol - 1
		if newCol < 0 {
			newCol = len(arr[0]) - 1
		}
	}

	if arr[newRow][newCol] == '#' {
		newRow = oldRow
		newCol = oldCol
	}

	return
}

func DrawDirection() string {
	dir := rand.Intn(4)
	move := map[int]string{
		0: "UP",
		1: "DOWN",
		2: "RIGHT",
		3: "LEFT",
	}
	return move[dir]
}

func MovePlayer(cmd string) {
	Player.row, Player.col = MakeMove(Player.row, Player.col, cmd)

	removeDot := func(row, col int) {
		arr[row] = arr[row][0:col] + " " + arr[row][col + 1:]
	}

	switch arr[Player.row][Player.col] {
	case '.':
		NumDots--
		Score++
		removeDot(Player.row, Player.col)
	case 'U':
		Score += 10
		removeDot(Player.row, Player.col)
		go ProcessPill()
	}
}

func MoveEnemy() {
	for _, e := range Enemies {
		dir := DrawDirection()
		e.position.row, e.position.col = MakeMove(e.position.row, e.position.col, dir)
		// remove dot
		arr[Player.row] = arr[Player.row][0:Player.col] + " " + arr[Player.row][Player.col + 1:]
	}
}

