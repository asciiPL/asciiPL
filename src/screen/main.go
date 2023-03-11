package screen

import (
	"github.com/rivo/tview"
)

type AppScreen struct {
	stack          *Stack
	app            *tview.Application
	HomeScreen     *Page
	NewGameScreen  *Page
	GenerateScreen *Page
	page           *tview.Pages
}

type Page struct {
	grid *tview.Grid
	name string
}

func (appScreen *AppScreen) SetScreen(screen *Page) {
	appScreen.stack.Push(screen)
	appScreen.page.SwitchToPage(screen.name)
}

func (appScreen *AppScreen) Back() {
	appScreen.stack.Pop()
	screen := appScreen.stack.Peek()
	appScreen.page.SwitchToPage(screen.name)
}

func (appScreen *AppScreen) Start() {
	layout := tview.NewGrid().
		SetBorders(false).
		SetColumns(3, -1).
		SetRows(12, 12, 12, -1).
		AddItem(appScreen.page, 0, 0, 12, 12, 0, 0, true)
	if err := appScreen.app.SetRoot(layout, true).Run(); err != nil {
		panic(err)
	}
}

func (appScreen *AppScreen) StartScreen(page Page) {
	layout := tview.NewGrid().
		SetBorders(false).
		SetColumns(3, -1).
		SetRows(12, 12, 12, -1).
		AddItem(page.grid, 0, 0, 12, 12, 0, 0, true)
	if err := appScreen.app.SetRoot(layout, true).Run(); err != nil {
		panic(err)
	}
}

func NewAppScreen() *AppScreen {
	app := tview.NewApplication()
	screen := &AppScreen{
		stack: NewStack(),
		app:   app,
	}
	pages := tview.NewPages()
	screen.page = pages
	screen.HomeScreen = NewHomeScreen(screen)
	screen.NewGameScreen = NewNewGameScreen(screen)
	screen.GenerateScreen = NewGenerateWorldScreen(screen)

	screen.stack.Push(screen.HomeScreen)
	return screen
}
