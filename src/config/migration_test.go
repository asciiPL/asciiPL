package config

import (
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

func TestLoadMigration(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	got, got1, got2, got3 := LoadMigration()
	require.Len(t, got, 1)
	require.Len(t, got1, 1)
	require.Len(t, got2, 1)
	require.Len(t, got3, 1)
	require.NotNil(t, got[0].Areas)
	require.NotNil(t, got1[0].Physics)
	require.NotNil(t, got[0].Areas)
	require.NotNil(t, got2[0].Psychology)
	require.NotNil(t, got3[0].Actions)
	require.NotNil(t, got3[0].Actions)
}

func Test_buildCfg(t *testing.T) {
	require.NoError(t, storeCfg())
}
