package util

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var isSeed = false

func GetValueFromRange(rangeStr string) string {
	if !isSeed {
		rand.Seed(time.Now().UnixNano())
		isSeed = true
	}

	// Split the range string into the minimum and maximum values.
	rangeArr := strings.Split(rangeStr, "-")
	if len(rangeArr) != 2 {
		// Invalid range string.
		return rangeStr
	}
	minVal, err := strconv.Atoi(rangeArr[0])
	if err != nil {
		// Invalid minimum value.
		return rangeStr
	}
	maxVal, err := strconv.Atoi(rangeArr[1])
	if err != nil {
		// Invalid maximum value.
		return rangeStr
	}

	// Generate a random integer within the specified range.
	randVal := rand.Intn(maxVal-minVal+1) + minVal

	// Convert the random integer to a string and return it.
	return strconv.Itoa(randVal)
}
