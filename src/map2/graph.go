package map2

import (
	"fmt"
	"github.com/TwiN/go-color"
	"math/rand"
	"sync"
	"time"
)

// house: 1 * 2
// house:

//type LandGrid *[][]Land
//
//type Land struct {
//	area *Area
//}
//
//type AreaType string
//
//const (
//	AreaTypeUpperCity AreaType = "Upper City"
//)
//
//type Area struct {
//	Grid[][]
//	areaType AreaType
//	neighborhood *Area
//}
//

type Area struct {
	id    int
	name  string
	color int
}

type Road struct {
	id    int
	name  string
	color int
	*MetaGrid
	grid *Grid
}

type Construction struct {
	id    int
	name  string
	color int
	*MetaGrid
	grid *Grid
}

type MetaGrid struct {
	objects []string
}

func (m MetaGrid) getObjects() []string {
	return m.objects
}

type IMetaGrid interface {
	getObjects() []string
}

var _ IMetaGrid = (*MetaGrid)(nil)

type Grid struct {
	area         *Area
	construction *Construction
	road         *Road
}

func NewConstructionGrid(area *Area, construction *Construction) *Grid {
	grid := Grid{area: area, construction: construction}
	construction.grid = &grid
	return &grid
}

func NewRoadGrid(area *Area, road *Road) *Grid {
	grid := Grid{area: area, road: road}
	road.grid = &grid
	return &grid
}

func DivArray(size int) {
	rand.Seed(time.Now().UnixNano())
	arr := make([][]int, size)
	for i := range arr {
		arr[i] = make([]int, size)
	}

	var wg sync.WaitGroup
	wg.Add(9)

	for i := 0; i <= 2; i++ {
		for j := 0; j <= 2; j++ {
			i := i
			j := j
			x := rand.Intn(size/3) + size/3*i
			y := rand.Intn(size/3) + size/3*j
			arr[x][y] = 3*i + j + 1
			go func() {
				defer wg.Done()
				bfs(&arr, x, y, 3*i+j+1)
			}()
		}
	}

	wg.Wait()

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if arr[i][j] == 0 {
				print(color.White, arr[i][j], " ", color.Reset)
			}
			if arr[i][j] == 1 {
				print(color.Red, arr[i][j], " ", color.Reset)
			}
			if arr[i][j] == 2 {
				print(color.Green, arr[i][j], " ", color.Reset)
			}
			if arr[i][j] == 3 {
				print(color.Yellow, arr[i][j], " ", color.Reset)
			}
			if arr[i][j] == 4 {
				print(color.Blue, arr[i][j], " ", color.Reset)
			}
			if arr[i][j] == 5 {
				print(color.Purple, arr[i][j], " ", color.Reset)
			}
			if arr[i][j] == 6 {
				print(color.Cyan, arr[i][j], " ", color.Reset)
			}
			if arr[i][j] == 7 {
				print(color.Gray, arr[i][j], " ", color.Reset)
			}
			if arr[i][j] == 8 {
				print(color.Red, arr[i][j], " ", color.Reset)
			}
			if arr[i][j] == 9 {
				print(color.Green, arr[i][j], " ", color.Reset)
			}
			print(" ")
		}
		fmt.Println()
	}
}

type cor struct {
	x int
	y int
}

func bfs(grid *[][]int, x, y, value int) {
	g := *grid
	m := len(g)
	n := len(g[0])
	vi := make([][]bool, m)
	numVi := 0
	for i := range vi {
		vi[i] = make([]bool, n)
	}

	queue := make([]cor, 0)
	if g[x][y] == value {
		queue = append(queue, cor{x, y})
	}

	for len(queue) != 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			curNode := queue[0]
			queue = queue[1:]
			if g[curNode.x][curNode.y] == 0 || (x == curNode.x && y == curNode.y) {
				g[curNode.x][curNode.y] = value
				numVi++
				if numVi*9 > m*n {
					time.Sleep(1 * time.Nanosecond)
				}
			} else {
				continue
			}
			if curNode.x > 0 && !vi[curNode.x-1][curNode.y] && g[curNode.x-1][curNode.y] == 0 {
				vi[curNode.x-1][curNode.y] = true
				queue = append(queue, cor{curNode.x - 1, curNode.y})
			}
			if curNode.y > 0 && !vi[curNode.x][curNode.y-1] && g[curNode.x][curNode.y-1] == 0 {
				vi[curNode.x][curNode.y-1] = true
				queue = append(queue, cor{curNode.x, curNode.y - 1})
			}
			if curNode.x < m-1 && !vi[curNode.x+1][curNode.y] && g[curNode.x+1][curNode.y] == 0 {
				vi[curNode.x+1][curNode.y] = true
				queue = append(queue, cor{curNode.x + 1, curNode.y})
			}
			if curNode.y < n-1 && !vi[curNode.x][curNode.y+1] && g[curNode.x][curNode.y+1] == 0 {
				vi[curNode.x][curNode.y+1] = true
				queue = append(queue, cor{curNode.x, curNode.y + 1})
			}
		}
	}
}

//* *
//
//* *
//* *
//
//* * * *
//
//* * *
//	*
//
//* * *
// *
