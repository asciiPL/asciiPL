package util

import (
	"os"
	"path"
	"sort"
	"strconv"
	"strings"
)

func ReadFile(filepath string) ([]byte, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

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
	cwd, err := os.Getwd()
	if err != nil {
		return res
	}

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
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	fullPath := path.Join(cwd, filepath)

	return os.WriteFile(fullPath, data, 0644)
}
