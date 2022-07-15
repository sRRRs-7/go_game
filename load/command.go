package load

import (
	"log"
	"os"
	"os/exec"
)

func Initialize() {
	cbreak := exec.Command("stty", "cbreak", "-echo")
	cbreak.Stdin = os.Stdin

	err := cbreak.Run()
	if err != nil {
		log.Fatalln("cbreak mode unable: ", err)
	}
}

func Cleanup() {
	cooked := exec.Command("stty", "-cbreak", "echo")
	cooked.Stdin = os.Stdin

	err := cooked.Run()
	if err != nil {
		log.Fatalln("cooked mode unable restore: ", err)
	}
}
