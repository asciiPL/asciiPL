package main

import (
	"github.com/asciiPL/asciiPL/src/config"
	"github.com/asciiPL/asciiPL/src/map"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	_map.Generate(60, config.LoadCfg(true).AreaConfig)
}
