package main

import (
	"awesomeProject/src/config"
	"awesomeProject/src/map2"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	areaCfg := config.MigrationConfig()
	map2.Generate(60, areaCfg)
}
