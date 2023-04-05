package testing

import (
	"github.com/asciiPL/asciiPL/src/config"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

func TestPraiseAction(t *testing.T) {

	cfg := config.LoadCfg(false)
	char1 := createHumanCharacter(cfg)
	char2 := createHumanCharacter(cfg)

	praiseAction := cfg.ActionConfig[1]
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	praiseAction.Execute(*char1, *char2)

	//program, err := expr.Compile("expression[0].command")
	//if err != nil {
	//	fmt.Printf("error compiling expression: %v\n", err)
	//	return
	//}
	//
	//output, err := expr.Run(program, data)
	//if err != nil {
	//	fmt.Printf("error executing expression: %v\n", err)
	//	return
	//}
	//
	//fmt.Println(output)

	require.NotNil(t, praiseAction)

	require.NotNil(t, char1)
	require.NotNil(t, char2)
}
