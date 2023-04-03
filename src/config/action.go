package config

import "github.com/asciiPL/asciiPL/src/model"

type ActionMigration struct {
	Actions []model.Action `yaml:"actions"`
}
