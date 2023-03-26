package util

import (
	"os"
	"path"
	"path/filepath"
	"runtime"
	"sort"
	"strconv"
	"strings"
)

var (
	_, b, _, _ = runtime.Caller(0)

	// Root folder of this project
	cwd = filepath.Join(filepath.Dir(b), "../..")
)

func ReadFile(filepath string) ([]byte, error) {
	fullPath := path.Join(cwd, filepath)

	return os.ReadFile(fullPath)
}

func DivideString(input string) float64 {
	parts := strings.Split(input, "/")
	numerator, _ := strconv.ParseFloat(parts[0], 64)
	denominator, _ := strconv.ParseFloat(parts[1], 64)
	return numerator / denominator
}

func ListFileConfig(folderPath string) []string {
	var res []string

	fullPath := path.Join(cwd, folderPath)

	files, err := os.ReadDir(fullPath)
	if err != nil {
		return res
	}

	for _, file := range files {
		if !file.IsDir() && strings.Contains(file.Name(), ".json") {
			res = append(res, file.Name())
		}
	}

	sort.Strings(res)

	return res
}

func WriteFile(filepath string, data []byte) error {

	fullPath := path.Join(cwd, filepath)

	return os.WriteFile(fullPath, data, 0644)
}
