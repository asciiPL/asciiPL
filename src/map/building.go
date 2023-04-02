package _map

import (
	"fmt"
)

type Building struct {
	structure [][][]int
	size      int
}

func convertBuilding(buildings []string) []*Building {
	res := make([]*Building, 0)
	for _, building := range buildings {
		structure := make([][]int, 0)
		row := make([]int, 0)
		s := 0
		for _, char := range building {
			if char == '*' {
				row = append(row, 1)
				s++
			} else if char == ' ' {
				row = append(row, 0)
			} else if char == '\n' {
				structure = append(structure, row)
				row = make([]int, 0)
			}
		}
		structure = append(structure, row)
		maxCol := 0
		for i := range structure {
			if maxCol < len(structure[i]) {
				maxCol = len(structure[i])
			}
		}
		for i := range structure {
			for len(structure[i]) < maxCol {
				structure[i] = append(structure[i], 0)
			}
		}
		res = append(res, &Building{
			structure: generateUniqueArrays(structure),
			size:      s,
		})
	}
	return res
}

func generateUniqueArrays(matrix [][]int) [][][]int {
	arrays := [][][]int{}
	for i := 0; i < 4; i++ {
		arrays = append(arrays, copyArr(matrix))
		matrix = rotate(matrix)
	}
	matrix = mirror(matrix)
	for i := 0; i < 4; i++ {
		arrays = append(arrays, copyArr(matrix))
		matrix = rotate(matrix)
	}
	return removeDuplicate(arrays)
}

func copyArr(matrix [][]int) [][]int {
	n, m := len(matrix), len(matrix[0])
	rotated := make([][]int, n)
	for j := range rotated {
		rotated[j] = make([]int, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			rotated[i][j] = matrix[i][j]
		}
	}
	return rotated
}

func removeDuplicate(uniqueArrays [][][]int) [][][]int {
	seen := make(map[string]bool)
	result := [][][]int{}
	for _, arr := range uniqueArrays {
		str := fmt.Sprintf("%v", arr)
		if !seen[str] {
			seen[str] = true
			result = append(result, arr)
		}
	}
	return result
}

func mirror(matrix [][]int) [][]int {
	n, m := len(matrix), len(matrix[0])

	for i := 0; i < n; i++ {
		for j := 0; j < m/2; j++ {
			matrix[i][j], matrix[i][m-j-1] = matrix[i][m-j-1], matrix[i][j]
		}
	}
	return matrix
}

func rotate(matrix [][]int) [][]int {
	m := len(matrix)
	n := len(matrix[0])
	rotated := make([][]int, n)
	for i := range rotated {
		rotated[i] = make([]int, m)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			rotated[j][m-i-1] = matrix[i][j]
		}
	}
	return rotated
}
