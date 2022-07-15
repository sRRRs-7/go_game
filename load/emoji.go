package load

import (
	"encoding/json"
	"os"

	"github.com/sRRRs-7/go_game/ansi"
)

func LoadConfig(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	decoder := json.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		return err
	}

	return nil
}

func CustomMoveCursor(row, col int) {
	if cfg.UseEmoji {
		ansi.MoveCursor(row, col * 2)
	} else {
		ansi.MoveCursor(row, col)
	}
}