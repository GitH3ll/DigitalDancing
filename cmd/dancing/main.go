package main

import (
	"github.com/GitH3ll/DigitalDancing/pkg/dance"
	"os"
)

func main() {
	wiggly:=dance.NewWiggly()
	wiggly.Dance(os.Stdout)
}
