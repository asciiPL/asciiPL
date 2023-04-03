package model

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
	Source     Record       `yaml:"source" json:"source"`
	Target     Record       `yaml:"source" json:"target"`
	Expression []Expression `yaml:"expression" json:"expression"`
}

type Expression struct {
	Index   int    `yaml:"index" json:"index"`
	Command string `yaml:"command" json:"command"`
}
