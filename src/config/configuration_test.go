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
	require.NotEqual(t, len(got.PhysicConfig), 0)
	require.NotEqual(t, len(got.AreaConfig), 0)
	require.NotEqual(t, len(got.PsychoConfig), 0)
}
