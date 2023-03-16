package main

import (
	"awesomeProject/src/config"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	config.MigrationConfig()
}
