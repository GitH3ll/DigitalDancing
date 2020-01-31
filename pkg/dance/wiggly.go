package dance

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

type Wiggly struct {
	moves []string
}

func NewWiggly() *Wiggly {
	return &Wiggly{
		moves: []string{
			`♪┏(・o･)┛♪`,
			`♪┗ ( ･o･) ┓♪`,
		},
	}
}

func (w Wiggly) Dance(file *os.File) {
	for {
		for _, move := range w.moves {
			clear(file)
			fmt.Println(move)
			time.Sleep(time.Second)
		}
	}
}

func clear(file *os.File) {
	cmd := exec.Command("clear")
	cmd.Stdout = file
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
