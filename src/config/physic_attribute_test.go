package config

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLoadCfg(t *testing.T) {
	got, got1 := LoadCfg()
	require.Len(t, got, 1)
	require.Len(t, got1, 1)
	require.NotNil(t, got[0].Areas)
	require.NotNil(t, got1[0].Physics)
}
