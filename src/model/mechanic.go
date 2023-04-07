package model

import (
	"errors"
	"fmt"
	"github.com/antonmedv/expr"
	"github.com/antonmedv/expr/vm"
	"github.com/asciiPL/asciiPL/src/util"
	"strconv"
	"strings"
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

func (action Action) Execute(source *Character, target *Character) error {
	if action.Source.Physic.ID != source.Physic.ID || action.Source.Psychology.ID != source.Psychology.ID {
		return errors.New("config is not correct")
	}
	if action.Target.Physic.ID != target.Physic.ID || action.Target.Psychology.ID != target.Psychology.ID {
		return errors.New("config is not correct")
	}
	expressions := action.Expression
	data := buildData(*source, *target)
	for _, exp := range expressions {

		env := map[string]interface{}{
			"PathSet":    strings.TrimSpace(strings.Split(exp.Command, "=")[0]),
			"Calculate":  strings.TrimSpace(strings.Split(exp.Command, "=")[1]),
			"Data":       data,
			"SetMapFunc": setMapFunc,
		}

		code := `SetMapFunc(PathSet, Calculate, Data)`

		program, err := expr.Compile(code, expr.Env(env))
		if err != nil {
			fmt.Printf("error compiling expression: %v\n", err)
			return err
		}

		output, err := expr.Run(program, env)
		if err != nil {
			fmt.Printf("error executing expression: %v\n", err)
			return err
		}

		data, _ = convertToMap(output)
	}
	delete(data, "tmp")
	source.Physic = reverseTransformRecord(data["source"].(map[string]interface{})["physic"].(map[string]interface{}))
	source.Psychology = reverseTransformRecord(data["source"].(map[string]interface{})["psychology"].(map[string]interface{}))
	target.Physic = reverseTransformRecord(data["target"].(map[string]interface{})["physic"].(map[string]interface{}))
	target.Psychology = reverseTransformRecord(data["target"].(map[string]interface{})["psychology"].(map[string]interface{}))
	return nil
}

func convertToMap(input interface{}) (map[string]interface{}, error) {
	if input == nil {
		return nil, errors.New("input value is nil")
	}

	if m, ok := input.(map[string]interface{}); ok {
		return m, nil
	}

	return nil, errors.New("input value is not a map[string]interface{}")
}

func buildData(source Character, target Character) map[string]interface{} {
	return map[string]interface{}{
		"source": map[string]interface{}{
			"physic":     transformRecord(source.Physic),
			"psychology": transformRecord(source.Psychology),
		},
		"target": map[string]interface{}{
			"physic":     transformRecord(target.Physic),
			"psychology": transformRecord(target.Psychology),
		},
	}
}

func setMapFunc(pathSet string, calculate string, data map[string]interface{}) map[string]interface{} {
	env := data
	typeConvert := ".toInt()"
	if strings.Contains(calculate, ".toInt()") {
		typeConvert = ".toInt()"
		calculate = strings.ReplaceAll(calculate, ".toInt()", "")
	}
	if strings.Contains(calculate, ".toString()") {
		typeConvert = ".toString()"
		calculate = strings.ReplaceAll(calculate, ".toString()", "")
	}

	program, err := expr.Compile(calculate, expr.Env(env))
	if err != nil {
		panic(err)
	}

	// Reuse this vm instance between runs
	v := vm.VM{}

	out, err := v.Run(program, env)
	if err != nil {
		panic(err)
	}

	if typeConvert == ".toInt()" {
		v, err := strconv.Atoi(fmt.Sprintf("%s", out))
		if err != nil {
			panic(err)
		}
		var i interface{} = v
		SetMap(data, pathSet, i)
		return data
	}

	if typeConvert == ".toString()" {
		v := fmt.Sprint(out)
		var i interface{} = v
		SetMap(data, pathSet, i)
		return data
	}

	SetMap(data, pathSet, out)
	return data
}

func SetMap(m map[string]interface{}, mapPath string, value interface{}) {
	// Split the mapPath into its parts
	keys := strings.Split(mapPath, ".")
	// Traverse the map based on the keys
	for i := 0; i < len(keys)-1; i++ {
		if m[keys[i]] == nil {
			m[keys[i]] = make(map[string]interface{})
		}
		m = m[keys[i]].(map[string]interface{})
	}
	// Set the value at the final key
	m[keys[len(keys)-1]] = value
}

func transformRecord(record Record) map[string]interface{} {
	result := make(map[string]interface{})
	result["id"] = record.ID
	result["name"] = record.Name

	attrMap := make(map[string]interface{})
	for _, attr := range record.Attribute {
		attrName := attr.Name
		if attrName == "" {
			continue
		}

		attrValue := make(map[string]interface{})
		for _, subAttr := range attr.Attribute {
			attrValue[subAttr.Name] = subAttr.Value
		}

		attrMap[attrName] = attrValue
	}

	result["attribute"] = attrMap
	return result
}

func reverseTransformRecord(m map[string]interface{}) Record {
	var record Record
	record.ID = m["id"].(int)
	record.Name = m["name"].(string)

	attributes := make([]Attribute, 0)
	for key, value := range m["attribute"].(map[string]interface{}) {
		attr := Attribute{Name: key}
		subAttrs := make([]Attribute, 0)
		for subKey, subValue := range value.(map[string]interface{}) {
			subAttr := Attribute{Name: subKey, Value: subValue.(string)}
			subAttrs = append(subAttrs, subAttr)
		}
		if len(subAttrs) > 0 {
			attr.Attribute = subAttrs
		}
		attributes = append(attributes, attr)
	}
	if len(attributes) > 0 {
		record.Attribute = attributes
	}
	return record
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
	for i := range attribute {
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
