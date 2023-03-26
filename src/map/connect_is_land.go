package _map

import (
	"awesomeProject/src/model"
	"fmt"
	"math"
)

type maxCons struct {
	i int
	j int
}

func MinimumConnectIslands(islands [][]*model.Grid) {
	xl, yl := len(islands), len(islands[0])
	pointer2Construction := map[string]bool{}
	constructions := make([]*model.Construction, 0)
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			cons := islands[i][j].Construction
			str := fmt.Sprint(islands[i][j].Construction)
			if cons != nil && !pointer2Construction[str] {
				pointer2Construction[str] = true
				constructions = append(constructions, cons)
			}
		}
	}

	for len(constructions) > 2 {
		maxDistance := 0

		maxCon := maxCons{
			i: 0,
			j: 0,
		}

		for i := 0; i < len(constructions); i++ {
			for j := i + 1; j < len(constructions); j++ {
				consI := constructions[i]
				consJ := constructions[j]
				dis := distance(consI.Grids[0], consJ.Grids[0])
				if maxDistance < dis {
					maxDistance = dis
					maxCon = maxCons{
						i: i,
						j: j,
					}
				}
			}
		}

		additionRemove := fillRoad(islands, constructions[maxCon.i], constructions[maxCon.j])
		if maxCon.i > maxCon.j {
			constructions = removeConstructionIndex(constructions, maxCon.i)
		} else {
			constructions = removeConstructionIndex(constructions, maxCon.j)
		}

		for _, c := range additionRemove {
			for i := len(constructions) - 1; i >= 0; i-- {
				if c == constructions[i] {
					constructions = removeConstructionIndex(constructions, i)
				}
			}
		}

	}

}

func fillRoad(grids [][]*model.Grid, construction *model.Construction, construction2 *model.Construction) (additionRemove []*model.Construction) {
	additionRemove = make([]*model.Construction, 0)

	worldStr := transformToInputWorld(grids, construction, construction2)
	worldWithRoad := Path(worldStr)
	xl, yl := len(grids), len(grids[0])
	i, j := 0, 0
	fromArea := construction.Grids[0].Area
	toArea := construction2.Grids[0].Area

	for _, r := range worldWithRoad {
		j++
		if r == '\n' {
			i++
			j = 0
		} else if i < xl && j < yl && r == '●' && grids[i][j].Construction == nil {
			grids[i][j].Road = model.NewRoad(fromArea, toArea)
			if i+1 < xl && grids[i+1][j].Construction != nil {
				additionRemove = append(additionRemove, grids[i+1][j].Construction)
			}
			if i-1 >= 0 && grids[i-1][j].Construction == nil {
				additionRemove = append(additionRemove, grids[i-1][j].Construction)
			}
			if j+1 < yl && grids[i][j+1].Construction == nil {
				additionRemove = append(additionRemove, grids[i][j+1].Construction)
			}
			if j-1 >= 0 && grids[i][j-1].Construction == nil {
				additionRemove = append(additionRemove, grids[i][j-1].Construction)
			}
		} else {
			continue
		}
	}

	return additionRemove
}

func transformToInputWorld(grids [][]*model.Grid, construction *model.Construction, construction2 *model.Construction) string {
	str := "\n"

	xl, yl := len(grids), len(grids[0])

	neighbors1 := make([]*model.Grid, 0)
	for _, g := range construction.Grids {
		if g.X+1 < xl && grids[g.X+1][g.Y].Construction == nil {
			neighbors1 = append(neighbors1, grids[g.X][g.Y])
		}
		if g.X-1 >= 0 && grids[g.X-1][g.Y].Construction == nil {
			neighbors1 = append(neighbors1, grids[g.X][g.Y])
		}
		if g.Y+1 < yl && grids[g.X][g.Y+1].Construction == nil {
			neighbors1 = append(neighbors1, grids[g.X][g.Y])
		}
		if g.Y-1 >= 0 && grids[g.X][g.Y-1].Construction == nil {
			neighbors1 = append(neighbors1, grids[g.X][g.Y])
		}
	}

	neighbors2 := make([]*model.Grid, 0)
	for _, g := range construction2.Grids {
		if g.X+1 < xl && grids[g.X+1][g.Y].Construction == nil {
			neighbors2 = append(neighbors2, grids[g.X][g.Y])
		}
		if g.X-1 >= 0 && grids[g.X-1][g.Y].Construction == nil {
			neighbors2 = append(neighbors2, grids[g.X][g.Y])
		}
		if g.Y+1 < yl && grids[g.X][g.Y+1].Construction == nil {
			neighbors2 = append(neighbors2, grids[g.X][g.Y])
		}
		if g.Y-1 >= 0 && grids[g.X][g.Y-1].Construction == nil {
			neighbors2 = append(neighbors2, grids[g.X][g.Y])
		}
	}

	grid1, grid2 := maxGridDistance(neighbors1, neighbors2)

	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			if grids[i][j] == grid1 {
				str += "F"
			} else if grids[i][j] == grid2 {
				str += "T"
			} else if grids[i][j].Construction == nil {
				str += "."
			} else {
				str += "X"
			}
		}
		str += "\n"
	}

	return str
}

func maxGridDistance(neighbors1 []*model.Grid, neighbors2 []*model.Grid) (*model.Grid, *model.Grid) {
	maxDistance := 0

	maxCon := maxCons{
		i: 0,
		j: 0,
	}

	for i, n1 := range neighbors1 {
		for j, n2 := range neighbors2 {
			dis := distance(n1, n2)
			if maxDistance < dis {
				maxDistance = dis
				maxCon = maxCons{
					i: i,
					j: j,
				}
			}
		}
	}

	return neighbors1[maxCon.i], neighbors2[maxCon.j]
}

func removeConstructionIndex(s []*model.Construction, index int) []*model.Construction {
	return append(s[:index], s[index+1:]...)
}

func distance(i *model.Grid, j *model.Grid) int {
	xi, yi := i.X, i.Y
	xj, yj := j.X, j.Y

	return int(math.Sqrt(float64((xi-xj)*(xi-xj) + (yi-yj)*(yi-yj))))
}
