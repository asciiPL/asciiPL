package model

import (
	"fmt"
	"github.com/asciiPL/asciiPL/src/util"
)

type Record struct {
	Name      string      `yaml:"name" json:"name"`
	ID        int         `yaml:"id" json:"id"`
	Attribute []Attribute `yaml:"attribute" json:"attribute"`
}

type Attribute struct {
	Name        string      `yaml:"name" json:"name,omitempty"`
	Value       string      `yaml:"value" json:"value,omitempty"`
	Description string      `yaml:"description" json:"description,omitempty"`
	Attribute   []Attribute `yaml:"attribute" json:"attribute,omitempty"`
}

type Action struct {
	Name       string       `yaml:"name" json:"name"`
	ID         int          `yaml:"id" json:"id"`
	Source     Character    `yaml:"source" json:"source"`
	Target     Character    `yaml:"source" json:"target"`
	Expression []Expression `yaml:"expression" json:"expression"`
}

func (action Action) execute(source Character, target Character) {
	if action.Source.Physic.ID != source.Physic.ID || action.Source.Psychology.ID != source.Psychology.ID {
		return
	}
	if action.Target.Physic.ID != target.Physic.ID || action.Target.Psychology.ID != target.Psychology.ID {
		return
	}
	expressions := action.Expression
	for _, exp := range expressions {
		fmt.Print(exp)
	}
}

type Expression struct {
	Index   int    `yaml:"index" json:"index"`
	Command string `yaml:"command" json:"command"`
}

type Character struct {
	x int
	y int
	// Character have same structure in life cycle
	Physic Record `yaml:"physic" json:"physic"`
	// each Character have each structure
	// change in character development
	Psychology Record `yaml:"psychology" json:"psychology"`
	// each Character have each structure
	// change in character development
	powerAttribute Record
	// define character's fate
	plot Plot
	// define target (Will get interesting Mission with good target)
	target Target
}

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
//   - Open ending: Do not solve all the problems that occur, leaving doubts and questions
type Plot struct {
	event []Event
	// time start plot
	start Time
	// time end plot
	end Time
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
	stateGrid      [][]*Grid
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

func NewCharacter(physicConfig Record, psychologyConfig Record) *Character {
	return &Character{
		Physic:     GetRecordValue(physicConfig),
		Psychology: GetRecordValue(psychologyConfig),
	}
}

func GetRecordValue(record Record) Record {
	record.Attribute = getAttributeValue(record.Attribute)
	return record
}

func getAttributeValue(attribute []Attribute) []Attribute {
	for i, _ := range attribute {
		attribute[i].Value = util.GetValueFromRange(attribute[i].Value)
		if attribute[i].Attribute != nil || len(attribute[i].Attribute) != 0 {
			attribute[i].Attribute = getAttributeValue(attribute[i].Attribute)
		}
	}
	return attribute
}

type Mission struct {
	// location where process mission
	grid *Grid
}
