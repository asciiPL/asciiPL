package config

import (
	"encoding/json"
	"errors"
	"github.com/asciiPL/asciiPL/src/model"
	"github.com/asciiPL/asciiPL/src/util"
	"github.com/spf13/viper"
	"golang.org/x/exp/maps"
	"log"
	"strings"
)

var (
	gridConfigFiles       = util.ListFileConfig("config/grid")
	physicAttributeFiles  = util.ListFileConfig("config/character_attribute/physics")
	psychologyConfigFiles = util.ListFileConfig("config/character_attribute/psychology")
	actionConfigFiles     = util.ListFileConfig("config/action")
)

func LoadMigration() ([]AreaMigration, []PhysicMigration,
	[]PsychologyMigration, []ActionMigration) {
	physicAttributeMigrations := make([]PhysicMigration, len(physicAttributeFiles))
	err := parseMigration("config/character_attribute/physics", physicAttributeFiles, append(make([]interface{}, 0), physicAttributeMigrations))
	if err != nil {
		log.Printf(err.Error())
		return nil, nil, nil, nil
	}

	areaMigrations := make([]AreaMigration, len(gridConfigFiles))
	err = parseMigration("config/grid", gridConfigFiles, append(make([]interface{}, 0), areaMigrations))
	if err != nil {
		log.Printf(err.Error())
		return nil, nil, nil, nil
	}

	psychologyMigration := make([]PsychologyMigration, len(psychologyConfigFiles))
	err = parseMigration("config/character_attribute/psychology", psychologyConfigFiles, append(make([]interface{}, 0), psychologyMigration))
	if err != nil {
		log.Printf(err.Error())
		return nil, nil, nil, nil
	}

	actionMigration := make([]ActionMigration, len(actionConfigFiles))
	err = parseMigration("config/action", actionConfigFiles, append(make([]interface{}, 0), actionMigration))
	if err != nil {
		log.Printf(err.Error())
		return nil, nil, nil, nil
	}

	return areaMigrations, physicAttributeMigrations,
		psychologyMigration, actionMigration
}

func parseMigration(configPath string, configFiles []string, migrationCfgs []interface{}) error {
	if len(configFiles) != len(migrationCfgs) {
		return errors.New("can't parse config with configFiles, not equals len")
	}
	v := viper.New()
	v.SetConfigType("yaml")
	v.AddConfigPath("../../" + configPath + "/")
	for i := 0; i < len(configFiles); i++ {
		v.SetConfigName(configFiles[i])
		if err := v.ReadInConfig(); err != nil {
			log.Fatalf("Error reading config file, %s", err)
		}

		if err := v.Unmarshal(&migrationCfgs[i]); err != nil {
			log.Fatalf("Error unmarshaling config, %s", err)
		}
	}
	return nil
}

type Data struct {
	Configs interface{} `json:"configs"`
	Version string      `json:"version"`
}

func storeConfig(configFiles string, migrationCfgs []interface{}, version string) error {
	v := viper.New()
	v.SetConfigType("json")
	v.SetConfigName(configFiles)
	v.AddConfigPath("../../config/data" + "/")
	b, err := json.Marshal(Data{
		Configs: migrationCfgs[0],
		Version: version,
	})
	if err != nil {
		return err
	}
	err = v.ReadConfig(strings.NewReader(string(b)))
	if err != nil {
		return err
	}
	err = v.WriteConfig()
	if err != nil {
		return err
	}
	return nil
}

func storeCfg() error {
	areaMigrations, physicAttributeMigrations,
		psychologyMigration, actionMigration := LoadMigration()
	err := storeConfig("area.json", append(make([]interface{}, 0), buildAreaConfig(areaMigrations)), gridConfigFiles[len(gridConfigFiles)-1])
	if err != nil {
		return err
	}
	err = storeConfig("character_attribute.physic.json", append(make([]interface{}, 0), buildPhysicConfig(physicAttributeMigrations)), physicAttributeFiles[len(physicAttributeFiles)-1])
	if err != nil {
		return err
	}
	err = storeConfig("character_attribute.psychology.json", append(make([]interface{}, 0), buildPsychologyConfig(psychologyMigration)), psychologyConfigFiles[len(psychologyConfigFiles)-1])
	if err != nil {
		return err
	}
	err = storeConfig("action.json", append(make([]interface{}, 0), buildActionConfig(actionMigration)), actionConfigFiles[len(actionConfigFiles)-1])
	if err != nil {
		return err
	}
	return nil
}

func buildActionConfig(migrations []ActionMigration) interface{} {
	id2Cfs := map[int]model.Action{}
	for _, cfg := range migrations {
		for _, action := range cfg.Actions {
			id2Cfs[action.ID] = action
		}
	}
	return maps.Values(id2Cfs)
}

func buildPsychologyConfig(migrations []PsychologyMigration) []model.Record {
	id2Cfs := map[int]model.Record{}
	for _, cfg := range migrations {
		for _, physic := range cfg.Psychology {
			id2Cfs[physic.ID] = physic
		}
	}
	return maps.Values(id2Cfs)
}

func buildAreaConfig(migrations []AreaMigration) []model.Area {
	id2Cfs := map[int]model.Area{}
	for _, cfg := range migrations {
		for _, area := range cfg.Areas {
			id2Cfs[area.Id] = area
		}
	}
	return maps.Values(id2Cfs)
}

func buildPhysicConfig(migrations []PhysicMigration) []model.Record {
	id2Cfs := map[int]model.Record{}
	for _, cfg := range migrations {
		for _, physic := range cfg.Physics {
			id2Cfs[physic.ID] = physic
		}
	}
	return maps.Values(id2Cfs)
}
