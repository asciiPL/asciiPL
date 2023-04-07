package main

import (
	_ "github.com/antonmedv/expr"
	"github.com/asciiPL/asciiPL/src/screen"
)

func main() {
	screen.NewAppScreen().Start()
}
