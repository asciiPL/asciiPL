package config

import (
	"awesomeProject/src/model"
)

type AreaMigration struct {
	Areas []model.Area `yaml:"areas"`
}
