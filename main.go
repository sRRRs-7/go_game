package main

import (
	"flag"
	"log"
	"time"

	"github.com/sRRRs-7/go_game/load"
)

func main() {
	// processing CLI flag
	flag.Parse()

	load.Initialize()
	defer load.Cleanup()

	// load field
	err := load.LoadFile(*load.TextFile)
	if err != nil {
		log.Fatal("cannot load text.txt")
		return
	}

	// load emojis
	err = load.LoadConfig(*load.ConfigFile)
	if err != nil {
		log.Fatal("cannot load config.json", err)
		return
	}

	// process input async
	input := make(chan string)
	go func(ch chan<- string) {
		for {
			input, err := load.ReadInput()
			if err != nil {
				log.Print("error reading input", err)
				ch <- "ESC"
			}
			ch <- input
		}
	}(input)

	// game loop
	for {
		// process movement
		select {
		case inp := <- input:
			if inp == "ESC" {
				load.Lives = 0
			}
			load.MovePlayer(inp)
		default:
		}
		load.MoveEnemy()

		// update screen
		load.PrintScreen()

		// process collision
		load.Collision()

		// check game over
		load.GameOver()

		// repeat
		time.Sleep(500 * time.Millisecond)
	}
}