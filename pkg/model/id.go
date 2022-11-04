package model

type ObjectID struct {
	Db640 string `yaml:"db640,omitempty"`
	ExtId string `yaml:"extId,omitempty"`
	Ifopt Ifopt  `yaml:"ifopt,omitempty"`
}

type Ifopt struct {
	Country  string
	State    string
	Stop     string
	Area     string
	Platform string
}
