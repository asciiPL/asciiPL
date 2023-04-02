package _map

import (
	"fmt"
	"github.com/TwiN/go-color"
	"github.com/asciiPL/asciiPL/src/model"
)

func PrintMap(arr [][]*model.Grid) {
	size := len(arr)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if arr[i][j].Construction != nil {
				print(color.White, arr[i][j].Area.Id, " ", color.Reset)
			} else if arr[i][j].Road != nil {
				print(color.Black, arr[i][j].Area.Id, " ", color.Reset)
			} else {
				if arr[i][j].Area == nil {
					print(color.White, arr[i][j].Area.Id, " ", color.Reset)
				} else if arr[i][j].Area.Color == 1 {
					print(color.Red, arr[i][j].Area.Id, " ", color.Reset)
				} else if arr[i][j].Area.Color == 2 {
					print(color.Green, arr[i][j].Area.Id, " ", color.Reset)
				} else if arr[i][j].Area.Color == 3 {
					print(color.Yellow, arr[i][j].Area.Id, " ", color.Reset)
				} else if arr[i][j].Area.Color == 4 {
					print(color.Blue, arr[i][j].Area.Id, " ", color.Reset)
				} else if arr[i][j].Area.Color == 5 {
					print(color.Purple, arr[i][j].Area.Id, " ", color.Reset)
				} else if arr[i][j].Area.Color == 6 {
					print(color.Cyan, arr[i][j].Area.Id, " ", color.Reset)
				} else if arr[i][j].Area.Color == 7 {
					print(color.Green, arr[i][j].Area.Id, " ", color.Reset)
				} else if arr[i][j].Area.Color == 8 {
					print(color.Red, arr[i][j].Area.Id, " ", color.Reset)
				} else if arr[i][j].Area.Color == 9 {
					print(color.Green, arr[i][j].Area.Id, " ", color.Reset)
				}
			}
			print(" ")
		}
		fmt.Println()
	}
}
