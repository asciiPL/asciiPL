package component

import (
	"awesomeProject/src/config"
	_map "awesomeProject/src/map"
	"awesomeProject/src/model"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"strconv"
)

func WorldScreen(size int) *tview.Grid {
	rows := make([]int, size)
	cols := make([]int, size)

	//[][]*model.Grid
	world := _map.Generate(size, config.LoadCfg(true).AreaConfig)

	grid := tview.NewGrid().
		SetRows(rows...).
		SetColumns(cols...)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			grid.AddItem(GridPrimitive(world[i][j]), i, j, 1, 1, 0, 100, false)
		}
	}

	return grid
}

var number2tcellColor = map[int]tcell.Color{
	1:  tcell.ColorRed,
	2:  tcell.ColorGreen,
	3:  tcell.ColorYellow,
	4:  tcell.ColorBlue,
	5:  tcell.ColorPurple,
	6:  tcell.ColorLightCyan,
	7:  tcell.ColorAqua,
	8:  tcell.ColorGold,
	9:  tcell.ColorLime,
	10: tcell.ColorWhite, // construction
	11: tcell.ColorBlack, // road
}

func GridPrimitive(grid *model.Grid) tview.Primitive {
	areaView := tview.NewTextView().
		SetTextAlign(tview.AlignCenter).SetText(strconv.Itoa(grid.Area.Id))

	if grid.Construction == nil && grid.Road == nil {
		areaView.SetTextColor(number2tcellColor[grid.Area.Color])
	} else if grid.Construction != nil {
		areaView.SetTextColor(number2tcellColor[10])
	} else {
		areaView.SetTextColor(number2tcellColor[11])
	}
	return areaView
}
