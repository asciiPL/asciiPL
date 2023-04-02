package main

import (
	"github.com/asciiPL/asciiPL/src/config"
	_map "github.com/asciiPL/asciiPL/src/map"
	"github.com/asciiPL/asciiPL/src/model"
)

func main() {
	area := config.LoadCfg(true).AreaConfig[1]
	grids := [][]*model.Grid{
		{model.NewGrid(&area, 0, 0), model.NewGrid(&area, 0, 1),
			model.NewGrid(&area, 0, 2), model.NewGrid(&area, 0, 3),
			model.NewGrid(&area, 0, 4)},

		{model.NewGrid(&area, 1, 0), model.NewGrid(&area, 1, 1),
			model.NewGrid(&area, 1, 2), model.NewGrid(&area, 1, 3),
			model.NewGrid(&area, 1, 4)},

		{model.NewGrid(&area, 2, 0), model.NewGrid(&area, 2, 1),
			model.NewGrid(&area, 2, 2), model.NewGrid(&area, 2, 3),
			model.NewGrid(&area, 2, 4)},

		{model.NewGrid(&area, 3, 0), model.NewGrid(&area, 3, 1),
			model.NewGrid(&area, 3, 2), model.NewGrid(&area, 3, 3),
			model.NewGrid(&area, 3, 4)},

		{model.NewGrid(&area, 4, 0), model.NewGrid(&area, 4, 1),
			model.NewGrid(&area, 4, 2), model.NewGrid(&area, 4, 3),
			model.NewGrid(&area, 4, 4)},
	}
	cons1 := model.NewConstruction(&area)
	grids[0][0].SetConstruction(cons1)
	grids[1][0].SetConstruction(cons1)
	grids[1][1].SetConstruction(cons1)

	cons2 := model.NewConstruction(&area)
	grids[0][4].SetConstruction(cons2)

	cons3 := model.NewConstruction(&area)
	grids[4][0].SetConstruction(cons3)

	cons4 := model.NewConstruction(&area)
	grids[4][2].SetConstruction(cons4)

	cons5 := model.NewConstruction(&area)
	grids[4][4].SetConstruction(cons5)

	_map.MinimumConnectIslands(grids)
	_map.PrintMap(grids)
}
