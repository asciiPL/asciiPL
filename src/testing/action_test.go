package testing

import (
	"github.com/asciiPL/asciiPL/src/config"
	"github.com/stretchr/testify/require"
	"log"
	"reflect"
	"testing"
)

func TestPraiseAction(t *testing.T) {

	cfg := config.LoadCfg(false)
	char1 := createHumanCharacter(cfg)
	char2 := createHumanCharacter(cfg)

	char2.Psychology.Attribute[0].Attribute[0].Value = "50" // happy = 50

	praiseAction := cfg.ActionConfig[1]
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	err := praiseAction.Execute(char1, char2)
	require.NoError(t, err)

	require.NotNil(t, praiseAction)

	require.NotNil(t, char1)
	require.NotNil(t, char2)

	require.True(t, reflect.DeepEqual(char1, char1))
	require.Equal(t, char2.Psychology.Attribute[0].Attribute[0].Value, "55")
	char2.Psychology.Attribute[0].Attribute[0].Value = "50" // happy = 50
	require.True(t, reflect.DeepEqual(char2, char2))
}
