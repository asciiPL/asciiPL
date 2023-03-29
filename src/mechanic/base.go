package mechanic

import "awesomeProject/src/model"

// Plot example: simulator three-act structure
// 1: Setup
//   - Introduction background and character
//   - First Incident
//
// 2: Confrontation
//   - Turning point 1 (Change character's normal life)
//   - Mid-point (Character can't back to normal life)
//   - Turning point 2 (Character failure and fall into a stalemate)
//
// 3: Resolution
//   - Climax: (Character make decisions to overcome stalemate)
//     End:
//   - Happy ending: Character win
//   - Sad ending: Character failure or die
//   - Open ending: Do not solve all the problems that occur in the movie, leaving doubts and questions
type Plot struct {
	event []Event
	// time start plot
	start Time
	// time end plot
	end Time
}

type Event struct {
	// ratio time in Plot like "20/100"
	ratioTime string
	// load id from Plot config file to get Event name, get ratioTime and generate expression
	id   int64
	name string
	// using Character (A) owner Plot and scan other Character (B) around for generate Action
	expression string
}

type Time int64

// PlotController Each character have plot, each plot have timing Event
// PlotController coordinate all events to make decision and order what Event execute,
// what Event change for service around Player
type PlotController struct {
	currentEvent   []Event
	stateGrid      [][]*model.Grid
	stateCharacter []*Character
	statePlayer    *Player
	// select decisionExpression will execute using state
	controllerExpression string
	// process decisionExpression to make decision how Event execute
	decisionExpression []string
}

// Player like Character but have more events around than normal Character
type Player struct {
	Character
}

type Character struct {
	x int
	y int
	// all Character have same structure
	physicsAttribute []Attribute
	// each Character have each structure
	// change in character development
	psychologyAttribute []Attribute
	// each Character have each structure
	// change in character development
	powerAttribute []Attribute
	// define character's fate
	plot Plot
	// define target (Will get interesting Mission with good target)
	target Target
}

type Target struct {
	id         int64
	name       string
	expression string
	// Get mission from other character
	mission []Mission
	// Transfer mission for other character to get Target
	transfer []Mission
}

type Mission struct {
	// location where process mission
	grid *model.Grid
}

type Attribute struct {
	code  string
	value string
}

type Action struct {
	id int64
	// expression define logic Character 1 -> Action -> Character 2
	//      => change Attribute Character 1 && change Attribute Character 2
	// Can create or change Target
	expression string
}
