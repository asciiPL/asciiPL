package config

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLoadMigration(t *testing.T) {
	got, got1, got2 := LoadMigration()
	require.Len(t, got, 1)
	require.Len(t, got1, 1)
	require.Len(t, got2, 1)
	require.NotNil(t, got[0].Areas)
	require.NotNil(t, got1[0].Physics)
	require.NotNil(t, got[0].Areas)
	require.NotNil(t, got2[0].Psychology)
}

func Test_buildCfg(t *testing.T) {
	require.NoError(t, storeCfg())
}
