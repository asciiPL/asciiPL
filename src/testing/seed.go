package testing

import (
	"github.com/asciiPL/asciiPL/src/config"
	"github.com/asciiPL/asciiPL/src/model"
)

func createHumanCharacter(cfg *config.Configuration) *model.Character {
	humanPhysicCfg := cfg.PhysicConfig[1]
	humanPsychoCfg := cfg.PsychoConfig[1]
	character := model.NewCharacter(humanPhysicCfg, humanPsychoCfg)
	return character
}
