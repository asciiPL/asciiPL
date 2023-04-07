package config

import (
	"github.com/asciiPL/asciiPL/src/model"
	"github.com/spf13/viper"
	"log"
)

type AreaConfig struct {
	Configs []model.Area `json:"configs"`
	Version string       `json:"version"`
}

type PhysicConfig struct {
	Configs []model.Record `json:"configs"`
	Version string         `json:"version"`
}

type PsychologyConfig struct {
	Configs []model.Record `json:"configs"`
	Version string         `json:"version"`
}

type ActionConfig struct {
	Configs []model.Action `json:"configs"`
	Version string         `json:"version"`
}

type Configuration struct {
	AreaConfig   map[int]model.Area
	PhysicConfig map[int]model.Record
	PsychoConfig map[int]model.Record
	ActionConfig map[int]model.Action
}

func LoadCfg(isRoot bool) *Configuration {
	areaConfig := AreaConfig{}
	err := parseConfig(isRoot, "config/data", "area.json", &areaConfig)
	if err != nil {
		log.Printf(err.Error())
		return nil
	}

	physicConfig := PhysicConfig{}
	err = parseConfig(isRoot, "config/data", "character_attribute.physic.json", &physicConfig)
	if err != nil {
		log.Printf(err.Error())
		return nil
	}
	psychoConfig := PsychologyConfig{}
	err = parseConfig(isRoot, "config/data", "character_attribute.psychology.json", &psychoConfig)
	if err != nil {
		log.Printf(err.Error())
		return nil
	}
	actionConfig := ActionConfig{}
	err = parseConfig(isRoot, "config/data", "action.json", &actionConfig)
	if err != nil {
		log.Printf(err.Error())
		return nil
	}
	return &Configuration{
		AreaConfig:   mappingArea(areaConfig),
		PhysicConfig: mappingPhysic(physicConfig),
		PsychoConfig: mappingPsycho(psychoConfig),
		ActionConfig: mappingAction(actionConfig),
	}
}

func mappingAction(config ActionConfig) map[int]model.Action {
	m := make(map[int]model.Action, 0)
	for _, config := range config.Configs {
		m[config.ID] = config
	}
	return m
}

func mappingPsycho(config PsychologyConfig) (m map[int]model.Record) {
	m = make(map[int]model.Record, 0)
	for _, config := range config.Configs {
		m[config.ID] = config
	}
	return m
}

func mappingPhysic(config PhysicConfig) (m map[int]model.Record) {
	m = make(map[int]model.Record, 0)
	for _, config := range config.Configs {
		m[config.ID] = config
	}
	return m
}

func mappingArea(config AreaConfig) (m map[int]model.Area) {
	m = make(map[int]model.Area, 0)
	for _, config := range config.Configs {
		m[config.Id] = config
	}
	return m
}

func parseConfig(isRoot bool, configPath string, configFile string, config interface{}) error {
	v := viper.New()
	v.SetConfigType("json")
	if !isRoot {
		v.AddConfigPath("../../" + configPath + "/")
	} else {
		v.AddConfigPath(configPath)
	}
	v.SetConfigName(configFile)
	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
		return err
	}

	if err := v.Unmarshal(&config); err != nil {
		log.Fatalf("Error unmarshaling config, %s", err)
		return err
	}
	return nil
}
