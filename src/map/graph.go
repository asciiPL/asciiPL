package _map

import (
	"awesomeProject/src/model"
	"awesomeProject/src/util"
	"math/rand"
	"sync"
	"time"
)

func Generate(size int, cfg map[int]model.Area) [][]*model.Grid {
	rand.Seed(time.Now().UnixNano())
	arr := make([][]*model.Grid, size)
	for i := range arr {
		arr[i] = make([]*model.Grid, size)
	}

	var wg sync.WaitGroup
	wg.Add(9)

	for i := 0; i <= 2; i++ {
		for j := 0; j <= 2; j++ {
			i := i
			j := j
			x := rand.Intn(size/3) + size/3*i
			y := rand.Intn(size/3) + size/3*j
			area := cfg[3*i+j+1]
			arr[x][y] = model.NewGrid(&area, x, y)
			go func() {
				defer wg.Done()
				bfs(&arr, x, y, &area)
			}()
		}
	}

	wg.Wait()

	area2Grid := map[*model.Area][]model.Grid{}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			grid := arr[i][j]
			area2Grid[grid.Area] = append(area2Grid[grid.Area], *grid)
		}
	}

	// Put building
	for area, grids := range area2Grid {
		buildCfgs := convertBuilding(area.Building)
		// Convert a string to a two-dimensional array

		gridSize := len(grids)

		noConstructionGrids := make([]model.Grid, 0)
		for _, g := range grids {
			if g.Construction == nil {
				noConstructionGrids = append(noConstructionGrids, g)
			}
		}
		initUse := gridSize - len(noConstructionGrids)

		for use := initUse; len(noConstructionGrids) > 0 && use <= int(float64(gridSize)*util.DivideString(area.ConstructionRate)); {
			randomIndex := rand.Intn(len(buildCfgs))

			buildCfg := buildCfgs[randomIndex]

			// Choose one of the directions
			structure := buildCfg.structure[rand.Intn(len(buildCfg.structure))]
			// Check condition from random grid
			index := rand.Intn(len(noConstructionGrids))
			grid := noConstructionGrids[index]
			if checkBuilding(grid, structure, &arr, buildCfg.size) {
				// create construct
				construct := model.NewConstruction(area)
				building(grid, structure, &arr, construct)
				// If pass: remove grid, use+= sizeBuilding
				use += buildCfg.size
			}

			noConstructionGrids = removeIndex(noConstructionGrids, index)
		}

	}

	// Put road_test
	MinimumConnectIslands(arr)
	// Print map
	// PrintMap(arr)
	return arr
}

func removeIndex(s []model.Grid, index int) []model.Grid {
	return append(s[:index], s[index+1:]...)
}

func building(grid model.Grid, structure [][]int, arr *[][]*model.Grid, construction *model.Construction) {
	x, y := grid.X, grid.Y
	n, m := len(structure), len(structure[0])
	g := *arr
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if structure[i][j] == 1 {
				g[x+i][y+j].SetConstruction(construction)
			}
		}
	}
}

func checkBuilding(grid model.Grid, structure [][]int, arr *[][]*model.Grid, size int) bool {
	x, y := grid.X, grid.Y
	n, m := len(structure), len(structure[0])
	g := *arr
	xl, yl := len(g), len(g[0])
	match := 0
	area := grid.Area
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if structure[i][j] == 1 {
				// check is not construction
				if x+i >= xl || y+j >= yl || g[x+i][y+j].Construction != nil {
					return false
				} else if g[x+i][y+j].Area != area {
					// check is same area
					return false
				} else if (x+i+1 >= xl || g[x+i+1][y+j].Construction != nil) ||
					(x+i-1 < 0 || g[x+i-1][y+j].Construction != nil) ||
					(y+j+1 >= yl || g[x+i][y+j+1].Construction != nil) ||
					(y+j-1 < 0 || g[x+i][y+j-1].Construction != nil) ||
					(g[x+i+1][y+j+1].Construction != nil) ||
					(g[x+i+1][y+j-1].Construction != nil) ||
					(g[x+i-1][y+j+1].Construction != nil) ||
					(g[x+i-1][y+j-1].Construction != nil) {
					// neigh check
					return false
				} else {
					match++
				}
			}
		}
	}
	return match == size
}

type cor struct {
	x int
	y int
}

func bfs(grid *[][]*model.Grid, x int, y int, area *model.Area) {
	g := *grid
	m := len(g)
	n := len(g[0])
	vi := make([][]bool, m)
	numVi := 0
	for i := range vi {
		vi[i] = make([]bool, n)
	}

	queue := make([]cor, 0)
	queue = append(queue, cor{x, y})

	for len(queue) != 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			curNode := queue[0]
			queue = queue[1:]
			if g[curNode.x][curNode.y] == nil || (x == curNode.x && y == curNode.y) {
				g[curNode.x][curNode.y] = model.NewGrid(area, curNode.x, curNode.y)
				numVi++
				if numVi*9 > m*n {
					time.Sleep(1 * time.Nanosecond)
				}
			} else {
				continue
			}
			if curNode.x > 0 && !vi[curNode.x-1][curNode.y] &&
				g[curNode.x-1][curNode.y] == nil {
				vi[curNode.x-1][curNode.y] = true
				queue = append(queue, cor{curNode.x - 1, curNode.y})
			}
			if curNode.y > 0 && !vi[curNode.x][curNode.y-1] && g[curNode.x][curNode.y-1] == nil {
				vi[curNode.x][curNode.y-1] = true
				queue = append(queue, cor{curNode.x, curNode.y - 1})
			}
			if curNode.x < m-1 && !vi[curNode.x+1][curNode.y] && g[curNode.x+1][curNode.y] == nil {
				vi[curNode.x+1][curNode.y] = true
				queue = append(queue, cor{curNode.x + 1, curNode.y})
			}
			if curNode.y < n-1 && !vi[curNode.x][curNode.y+1] && g[curNode.x][curNode.y+1] == nil {
				vi[curNode.x][curNode.y+1] = true
				queue = append(queue, cor{curNode.x, curNode.y + 1})
			}
		}
	}
}
