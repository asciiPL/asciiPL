package screen

import (
	"awesomeProject/src/screen/component"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func NewGenerateWorldScreen(appScreen *AppScreen) *Page {

	grid := tview.NewGrid().
		SetRows(5, 0, 0).
		SetColumns(0, 0, 0).
		SetBorders(false)

	log := tview.NewList().
		ShowSecondaryText(false).
		AddItem("_GEN WORLD_ Process generate map", "", 0, func() {}).
		AddItem("_GEN WORLD_ Complete generate map", "", 0, func() {}).
		AddItem("_GEN WORLD_ Process generate people", "", 0, func() {}).
		AddItem("_GEN WORLD_ Complete generate 32 people", "", 0, func() {}).
		AddItem("_GEN WORLD_ Process generate event", "", 0, func() {}).
		AddItem("_GEN WORLD_ Complete generate 1963 people", "", 0, func() {}).
		AddItem("_GEN WORLD_ Start with year 0", "", 0, func() {}).
		AddItem("_GEN WORLD_ Process event, current year 200", "", 0, func() {}).
		AddItem("_GEN WORLD_ Complete generate world, ready to play", "", 0, func() {})

	log.SetTitle("Log").SetBorder(true)

	option := tview.NewList().
		ShowSecondaryText(false).
		AddItem("Play this world", "", 'p', func() {}).
		AddItem("Save world and back", "", 's', func() { appScreen.Back() }).
		AddItem("Back", "", 'b', func() { appScreen.Back() })
	option.SetTitle("Option").SetBorder(true)

	selections := []*tview.List{option, log}

	for i, box := range selections {
		(func(index int) {
			box.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
				switch event.Key() {
				case tcell.KeyTab:
					appScreen.app.SetFocus(selections[(index+1)%len(selections)])
					return nil
				case tcell.KeyBacktab:
					appScreen.app.SetFocus(selections[(index+len(selections)-1)%len(selections)])
					return nil
				}
				return event
			})
		})(i)
	}

	grid.AddItem(option, 0, 0, 1, 1, 0, 100, true).
		AddItem(log, 1, 0, 2, 1, 0, 100, false)

	appScreen.page.AddPage("generate_world", grid, true, false)
	return &Page{
		grid,
		"generate_world",
	}

}

var currentWorld = tview.NewGrid()

func GenWorld(page *Page) {
	grid := page.grid
	grid.RemoveItem(currentWorld)
	currentWorld = component.WorldScreen(30)
	grid.AddItem(currentWorld, 0, 1, 3, 2, 0, 100, false)
}
