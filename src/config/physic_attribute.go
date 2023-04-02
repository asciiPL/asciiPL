package config

type Physics struct {
	Name      string      `yaml:"name" json:"name"`
	ID        int         `yaml:"id" json:"id"`
	Attribute []Attribute `yaml:"attribute" json:"attribute"`
}

type Attribute struct {
	Name      string      `yaml:"name" json:"name"`
	Value     string      `yaml:"value" json:"value"`
	Attribute []Attribute `yaml:"attribute" json:"attribute"`
}

type PhysicMigration struct {
	Physics []Physics `yaml:"physics"`
}
