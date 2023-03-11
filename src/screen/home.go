package screen

import "github.com/rivo/tview"

func NewHomeScreen(appScreen *AppScreen) *Page {
	imgType := tview.NewList().
		ShowSecondaryText(false).
		AddItem("New Game", "", 'n', func() { appScreen.SetScreen(appScreen.NewGameScreen) }).
		AddItem("Continue", "", 'c', func() {}).
		AddItem("Exit", "", 'e', func() { appScreen.app.Stop() })
	imgType.SetTitle("Home").SetBorder(true)
	grid := tview.NewGrid().
		SetBorders(false).
		SetColumns(3, -1).
		SetRows(12, 12, 12, -1).
		AddItem(imgType, 0, 0, 12, 12, 0, 0, true)
	appScreen.page.AddPage("home", grid, true, true)
	return &Page{
		grid,
		"home",
	}

}
