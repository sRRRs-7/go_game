package load

import (
	"bufio"
	"log"
	"os"
)

var arr []string
var Score int
var NumDots int
var Lives = 3

var Enemies []*Enemy

const (
	EnemyStatusNormal EnemyStatus = "Normal"
	EnemyStatusBlue EnemyStatus = "Blue"
)

type EnemyStatus string

type Enemy struct {
	position Point
	status EnemyStatus
}

func LoadFile(file string) error {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal("cannot load file")
		return err
	}
	defer f.Close()

	scan := bufio.NewScanner(f)
	for scan.Scan() {
		line := scan.Text()
		arr = append(arr, line)
	}
	if scan.Err() != nil {
		log.Fatal("text file append error")
		return err
	}

	for row, line := range arr {
		for col, char := range line {
			switch char {
			case 'P':
				Player = Point{row, col, row, col}
			case 'E':
				Enemies = append(Enemies, &Enemy{Point{row, col, row, col}, EnemyStatusNormal})
			case '.':
				NumDots++
			}
		}
	}
	return nil
}

func ReadInput() (string, error) {
	buffer := make([]byte, 100)

	cnt, err := os.Stdin.Read(buffer)
	if err != nil {
		return "", err
	}

	// 0x1b === "ESC"
	if cnt == 1 && buffer[0] == 0x1b {
		return "ESC", nil
	} else if cnt >= 3 {
		if buffer[0] == 0x1b && buffer[1] == '[' {
			switch buffer[2] {
			case 'A':
				return "UP", nil
			case 'B':
				return "DOWN", nil
			case 'C':
				return "RIGHT", nil
			case 'D':
				return "LEFT", nil
			}
		}
	}

	return "", nil
}