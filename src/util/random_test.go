package util

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetValueFromRange(t *testing.T) {
	exist := map[string]bool{}

	for i := 0; i < 10000; i++ {
		got := GetValueFromRange("1-6")
		exist[got] = true
	}

	require.Len(t, exist, 6)

}
