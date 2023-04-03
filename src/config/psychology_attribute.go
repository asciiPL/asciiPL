package config

import "github.com/asciiPL/asciiPL/src/model"

type PsychologyMigration struct {
	Psychology []model.Record `yaml:"psychology"`
}
