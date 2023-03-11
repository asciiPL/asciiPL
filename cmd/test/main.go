package main

import (
	"awesomeProject/src/screen"
)

func main() {
	sc := screen.NewAppScreen()
	sc.StartScreen(*sc.GenerateScreen)
}
