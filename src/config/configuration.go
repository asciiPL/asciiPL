package config

import "awesomeProject/src/model"

type AreaConfig struct {
	Configs []model.Area `json:"configs"`
	Version string       `json:"version"`
}

type Configuration struct {
	AreaConfig map[int]model.Area
}
