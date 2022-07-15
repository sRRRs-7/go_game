package load

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/sRRRs-7/go_game/ansi"
)

func PrintScreen() {
	ansi.ClearScreen()
	for _, line := range arr {
		for _, char := range line {
			switch char {
			case '#':
				fmt.Print(ansi.WithBlueBackground(cfg.Wall))
			case '.':
				fmt.Print(cfg.Dot)
			default:
				fmt.Print(cfg.Space)
			}
		}
		fmt.Println()
	}

	CustomMoveCursor(Player.row, Player.col)
	fmt.Print(cfg.Player)

	for _, e := range Enemies {
		CustomMoveCursor(e.position.row, e.position.col)
		if e.status == EnemyStatusNormal {
			fmt.Println(cfg.Enemy)
		} else if e.status == EnemyStatusBlue {
			fmt.Println(cfg.EnemyBlue)
		}
	}

	livesRemaining := strconv.Itoa(Lives) //converts lives int to a string
    if cfg.UseEmoji {
        livesRemaining = GetLivesAsEmoji()
    }

	// print score
	CustomMoveCursor(len(arr) + 1, 0)
	fmt.Println("Score:", Score, "\tLives:", livesRemaining)
}

func GetLivesAsEmoji() string {
	buf := bytes.Buffer{}
	for i := Lives; i > 0; i-- {
		buf.WriteString(cfg.Player)
	}
	return buf.String()
}



