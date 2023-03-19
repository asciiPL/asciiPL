package _map

import (
	"fmt"
	"github.com/TwiN/go-color"
)

func PrintMap(arr [][]*Grid) {
	size := len(arr)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if arr[i][j].area == nil {
				print(color.White, arr[i][j].area.Id, " ", color.Reset)
			} else if arr[i][j].area.Color == 1 {
				print(color.Red, arr[i][j].area.Id, " ", color.Reset)
			} else if arr[i][j].area.Color == 2 {
				print(color.Green, arr[i][j].area.Id, " ", color.Reset)
			} else if arr[i][j].area.Color == 3 {
				print(color.Yellow, arr[i][j].area.Id, " ", color.Reset)
			} else if arr[i][j].area.Color == 4 {
				print(color.Blue, arr[i][j].area.Id, " ", color.Reset)
			} else if arr[i][j].area.Color == 5 {
				print(color.Purple, arr[i][j].area.Id, " ", color.Reset)
			} else if arr[i][j].area.Color == 6 {
				print(color.Cyan, arr[i][j].area.Id, " ", color.Reset)
			} else if arr[i][j].area.Color == 7 {
				print(color.Gray, arr[i][j].area.Id, " ", color.Reset)
			} else if arr[i][j].area.Color == 8 {
				print(color.Red, arr[i][j].area.Id, " ", color.Reset)
			} else if arr[i][j].area.Color == 9 {
				print(color.Green, arr[i][j].area.Id, " ", color.Reset)
			}
			print(" ")
		}
		fmt.Println()
	}
}
