package config

import (
	"awesomeProject/src/map"
	"awesomeProject/src/util"
	"encoding/json"
	"log"
	"strings"
)

type Config struct {
	Areas []_map.Area `json:"areas"`
}

type Migration struct {
	Version string `json:"version"`
	Name    string `json:"name"`
}

func MigrationConfig() map[int]_map.Area {
	gridConfigFiles := util.ListFileConfig("config/grid")
	for _, fileName := range gridConfigFiles {
		MigrationGridConfig(fileName)
	}
	areaByte, err := util.ReadFile("config/data/areas.json")
	if err != nil {
		log.Fatal(err)
	}
	var areaCfg []_map.Area
	err = json.Unmarshal(areaByte, &areaCfg)
	if err != nil {
		log.Fatal(err)
	}
	mapCfg := map[int]_map.Area{}
	for _, area := range areaCfg {
		mapCfg[area.Id] = area
	}
	return mapCfg
}

func MigrationGridConfig(fileName string) {
	//configFile, err := util.ReadFile("config/grid/0.0.1_init_config.json")
	configFile, err := util.ReadFile("config/grid/" + fileName)
	if err != nil {
		log.Fatal(err)
	}

	var config Config
	err = json.Unmarshal(configFile, &config)
	if err != nil {
		log.Fatal(err)
	}

	areasJSON, err := json.Marshal(config.Areas)
	if err != nil {
		log.Fatal(err)
	}

	err = util.WriteFile("config/data/areas.json", areasJSON)
	if err != nil {
		log.Fatal(err)
	}

	version := strings.Split(fileName, "_")[0]
	migrationName := strings.ReplaceAll(fileName, version+"_", "")
	migrationName = strings.ReplaceAll(migrationName, ".json", "")

	migrationJSON, err := json.Marshal(Migration{
		Version: version,
		Name:    migrationName,
	})
	err = util.WriteFile("config/data/migration.json", migrationJSON)
	if err != nil {
		log.Fatal(err)
	}
}
