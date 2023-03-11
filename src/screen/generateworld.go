package screen

import "github.com/rivo/tview"

func NewGenerateWorldScreen(appScreen *AppScreen) *Page {

	log := tview.NewList().
		ShowSecondaryText(false).
		AddItem("Generate new world", "", 'g', func() {}).
		AddItem("Back", "", 'b', func() { appScreen.Back() })
	log.SetTitle("Log").SetBorder(true)

	option := tview.NewList().
		ShowSecondaryText(false).
		AddItem("Generate new world", "", 'g', func() {}).
		AddItem("Back", "", 'b', func() { appScreen.Back() })
	option.SetTitle("Option").SetBorder(true)

	grid := tview.NewGrid().
		SetColumns(0, 0).
		SetBorders(false)

	grid.AddItem(option, 0, 0, 1, 1, 0, 100, true).
		AddItem(log, 0, 1, 1, 1, 0, 100, false)

	appScreen.page.AddPage("generate_world", grid, true, false)
	return &Page{
		grid,
		"generate_world",
	}

}
