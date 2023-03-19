package main

import (
	"awesomeProject/src/config"
	"awesomeProject/src/map"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	areaCfg := config.MigrationConfig()
	_map.Generate(60, areaCfg)
}
