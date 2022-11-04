package model

type InfrastructureObject struct {
	Name     string   `yaml:"name"`
	Location float32  `yaml:"location"`
	Distance float32  `yaml:"distance"`
	Type     string   `yaml:"type"`
	Id       ObjectID `yaml:"id"`
}
