package testing

import (
	"github.com/asciiPL/asciiPL/src/config"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPraiseAction(t *testing.T) {

	cfg := config.LoadCfg(false)
	char1 := createHumanCharacter(cfg)
	char2 := createHumanCharacter(cfg)

	praiseAction := cfg.ActionConfig[0]

	//praiseAction.Execute()

	require.NotNil(t, praiseAction)

	require.NotNil(t, char1)
	require.NotNil(t, char2)
}
