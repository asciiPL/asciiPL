package config

import (
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

func TestLoadCfg(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	got := LoadCfg(false)
	require.NotNil(t, got)
	require.NotNil(t, got.PhysicConfig)
	require.NotNil(t, got.AreaConfig)
}
