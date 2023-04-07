package config

import "github.com/asciiPL/asciiPL/src/model"

type PhysicMigration struct {
	Physics []model.Record `yaml:"physics"`
}
