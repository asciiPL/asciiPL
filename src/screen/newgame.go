package screen

import "github.com/rivo/tview"

func NewNewGameScreen(appScreen *AppScreen) *Page {
	imgType := tview.NewList().
		ShowSecondaryText(false).
		AddItem("Generate new world", "", 'g', func() { appScreen.SetScreen(appScreen.GenerateScreen) }).
		AddItem("Back", "", 'b', func() { appScreen.Back() })
	imgType.SetTitle("New game").SetBorder(true)
	grid := tview.NewGrid().
		SetBorders(false).
		SetColumns(3, -1).
		SetRows(12, 12, 12, -1).
		AddItem(imgType, 0, 0, 12, 12, 0, 0, true)
	appScreen.page.AddPage("new_game", grid, true, false)
	return &Page{
		grid,
		"new_game",
	}

}
