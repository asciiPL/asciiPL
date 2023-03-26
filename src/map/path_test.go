package _map

import (
	"github.com/beefsack/go-astar"
	"testing"
)

// testPath takes a string encoded world, decodes it, calculates a path and
// checks the expected distance matches.  An expectedDist of -1 expects that no
// path will be found.
func testPath(worldInput string, t *testing.T, expectedDist float64) {
	world := ParseWorld(worldInput)
	t.Logf("Input world\n%s", world.RenderPath([]astar.Pather{}))
	p, dist, found := astar.Path(world.From(), world.To())
	if !found {
		t.Log("Could not find a path")
	} else {
		t.Logf("Resulting path\n%s", world.RenderPath(p))
	}
	if !found && expectedDist >= 0 {
		t.Fatal("Could not find a path")
	}
	if found && dist != expectedDist {
		t.Fatalf("Expected dist to be %v but got %v", expectedDist, dist)
	}
}

// TestStraightLine checks that having no obstacles results in a straight line
// path.
func TestStraightLine(t *testing.T) {
	testPath(`
.....~......
.....MM.....
.F........T.
....MMM.....
............
`, t, 9)
}

// TestPathAroundMountain checks that having a round mountain in the path
// results in a path around the mountain.
func TestPathAroundMountain(t *testing.T) {
	testPath(`
.....~......
.....MM.....
.F..MMMM..T.
....MMM.....
............
`, t, 13)
}

// TestBlocked checks that no path is returned when there is no possible path.
func TestBlocked(t *testing.T) {
	testPath(`
............
.........XXX
.F.......XTX
.........XXX
............
`, t, -1)
}

// TestMaze checks that paths can double back on themselves to reach the goal.
func TestMaze(t *testing.T) {
	testPath(`
FX.X........
.X...XXXX.X.
.X.X.X....X.
...X.X.XXXXX
.XX..X.....T
`, t, 27)
}

// TestMountainClimber checks that a path will choose to go over a mountain,
// which has a movement penalty of 3, if it's faster than going around the
// mountain.
func TestMountainClimber(t *testing.T) {
	testPath(`
..F..M......
.....MM.....
....MMMM..T.
....MMM.....
............
`, t, 12)
}

// TestRiverSwimmer checks that the path will prefer to cross a river, which
// has a movement penalty of 2, over a mountain which has a movement penalty of
// 3.
func TestRiverSwimmer(t *testing.T) {
	testPath(`
.....~......
.....~......
.F...X...T..
.....M......
.....M......
`, t, 11)
}

func BenchmarkLarge(b *testing.B) {
	world := ParseWorld(`
F............................~.................................................
.............................~.................................................
........M...........X........~.................................................
.......MMM.........X.........~~................................................
........MM........X...........~................................................
.......MM........X............~................................................
................X.............~................................................
...............X..............~~...............................................
..............X................~...............................................
.............X.................~...X...............~...........................
............X.......................X..............~...........................
...........X.........................X.............~...........................
..........X..................~........X............~...........................
.........X...................~.........X...........~...........................
.............................~..........X..........~...............XXXXXXXXXXXX
............................~............X..........~..............X...X...X...
............................~.............X.........~......MMM.....X.X.X.X.X.X.
............................~..............X........~......MM......X.X.X.X.X.X.
............................~...............X.......~....MMMM......X.X.X.X.X.X.
...........................~.................X.....~......MMM......X.X.X.X.X.X.
..............................................X....~.......MM......X.X.X.X.X.X.
...............................................X...~.......M.........X...X...XT
`)
	for i := 0; i < b.N; i++ {
		astar.Path(world.From(), world.To())
	}
}
