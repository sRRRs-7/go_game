package load

import (
	"fmt"
	"time"
)

func Restart(e *Point) {
	if Player == *e {
		Lives = Lives - 1
		if Lives != 0 {
			CustomMoveCursor(Player.row, Player.col)
			fmt.Print(cfg.Death)
			CustomMoveCursor(len(arr) + 2, 0)
			time.Sleep(1000 * time.Millisecond)
			Player.row, Player.col = Player.startRow, Player.startCol
		}
	 }
}