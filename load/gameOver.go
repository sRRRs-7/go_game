package load

import (
	"fmt"
	"log"
)

func GameOver() {
	if NumDots == 0 || Lives == 0 {
		if Lives == 0 {
			CustomMoveCursor(Player.row, Player.col)
			fmt.Print(cfg.Death)
			CustomMoveCursor(len(arr) + 2, 0)
		}
		log.Panic("Game Over")
	}
}