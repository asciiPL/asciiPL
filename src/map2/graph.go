package map2

import (
	"math/rand"
	"sync"
	"time"
)

type Area struct {
	Id               int      `json:"id"`
	Name             string   `json:"name"`
	Color            int      `json:"color"`
	Size             string   `json:"size"`
	ConstructionRate string   `json:"constructionRate"`
	Building         []string `json:"building"`
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

func NewGrid(area *Area) *Grid {
	grid := Grid{area: area}
	return &grid
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

func Generate(size int, cfg map[int]Area) {
	rand.Seed(time.Now().UnixNano())
	arr := make([][]*Grid, size)
	for i := range arr {
		arr[i] = make([]*Grid, size)
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
			arr[x][y] = NewGrid(&area)
			go func() {
				defer wg.Done()
				bfs(&arr, x, y, &area)
			}()
		}
	}

	wg.Wait()

	PrintMap(arr)
}

type cor struct {
	x int
	y int
}

func bfs(grid *[][]*Grid, x int, y int, area *Area) {
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
				g[curNode.x][curNode.y] = NewGrid(area)
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
