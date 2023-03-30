package config

import (
	"awesomeProject/src/util"
	"errors"
	"github.com/spf13/viper"
	"log"
)

type Physics struct {
	Name      string      `yaml:"name"`
	ID        int64       `yaml:"id"`
	Attribute []Attribute `yaml:"attribute"`
}

type Attribute struct {
	Name      string      `yaml:"name"`
	Value     string      `yaml:"value"`
	Attribute []Attribute `yaml:"attribute"`
}

var (
	gridConfigFiles      = util.ListFileConfig("config/grid")
	physicAttributeFiles = util.ListFileConfig("config/character_attribute/physics")
)

type MigrationCfg struct {
	Physics []Physics `yaml:"physics"`
}

type Configuration struct {
	v *viper.Viper
}

func LoadCfg() ([]Config, []MigrationCfg) {

	physicAttributeMigrations := make([]MigrationCfg, len(physicAttributeFiles))
	err := parseConfig("config/character_attribute/physics", physicAttributeFiles, append(make([]interface{}, 0), physicAttributeMigrations))
	if err != nil {
		log.Printf(err.Error())
		return nil, nil
	}

	areaMigration := make([]Config, len(gridConfigFiles))
	err = parseConfig("config/grid", gridConfigFiles, append(make([]interface{}, 0), areaMigration))
	if err != nil {
		log.Printf(err.Error())
		return nil, nil
	}

	return areaMigration, physicAttributeMigrations
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

func BuildCfg() {

}
