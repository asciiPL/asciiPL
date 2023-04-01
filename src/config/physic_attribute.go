package config

import (
	"awesomeProject/src/model"
	"awesomeProject/src/util"
	"encoding/json"
	"errors"
	"github.com/spf13/viper"
	"golang.org/x/exp/maps"
	"log"
	"strings"
)

type Physics struct {
	Name      string      `yaml:"name" json:"name"`
	ID        int         `yaml:"id" json:"id"`
	Attribute []Attribute `yaml:"attribute" json:"attribute"`
}

type Attribute struct {
	Name      string      `yaml:"name" json:"name"`
	Value     string      `yaml:"value" json:"value"`
	Attribute []Attribute `yaml:"attribute" json:"attribute"`
}

var (
	gridConfigFiles      = util.ListFileConfig("config/grid")
	physicAttributeFiles = util.ListFileConfig("config/character_attribute/physics")
)

type MigrationCfg struct {
	Physics []Physics `yaml:"physics"`
}

func LoadCfg() ([]Config, []MigrationCfg) {
	physicAttributeMigrations := make([]MigrationCfg, len(physicAttributeFiles))
	err := parseConfig("config/character_attribute/physics", physicAttributeFiles, append(make([]interface{}, 0), physicAttributeMigrations))
	if err != nil {
		log.Printf(err.Error())
		return nil, nil
	}

	areaMigrations := make([]Config, len(gridConfigFiles))
	err = parseConfig("config/grid", gridConfigFiles, append(make([]interface{}, 0), areaMigrations))
	if err != nil {
		log.Printf(err.Error())
		return nil, nil
	}

	return areaMigrations, physicAttributeMigrations
}

func parseConfig(configPath string, configFiles []string, migrationCfgs []interface{}) error {
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
	Configs []interface{} `json:"configs"`
	Version string        `json:"version"`
}

func storeConfig(configFiles string, migrationCfgs []interface{}, version string) error {
	v := viper.New()
	v.SetConfigType("json")
	v.SetConfigName(configFiles)
	v.AddConfigPath("../../config/data" + "/")
	b, err := json.Marshal(Data{
		Configs: migrationCfgs,
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
	areaMigrations, physicAttributeMigrations := LoadCfg()
	err := storeConfig("area.json", append(make([]interface{}, 0), buildAreaConfig(areaMigrations)), gridConfigFiles[len(gridConfigFiles)-1])
	if err != nil {
		return err
	}
	err = storeConfig("character_attribute.physic.json", append(make([]interface{}, 0), buildPhysicConfig(physicAttributeMigrations)), physicAttributeFiles[len(physicAttributeFiles)-1])
	if err != nil {
		return err
	}
	return nil
}

func buildAreaConfig(migrations []Config) []model.Area {
	id2Cfs := map[int]model.Area{}
	for _, cfg := range migrations {
		for _, area := range cfg.Areas {
			id2Cfs[area.Id] = area
		}
	}
	return maps.Values(id2Cfs)
}

func buildPhysicConfig(migrations []MigrationCfg) []Physics {
	id2Cfs := map[int]Physics{}
	for _, cfg := range migrations {
		for _, physic := range cfg.Physics {
			id2Cfs[physic.ID] = physic
		}
	}
	return maps.Values(id2Cfs)
}
