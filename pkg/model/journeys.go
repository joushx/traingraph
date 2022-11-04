package model

type Journey struct {
	ID    string     `yaml:"id"`
	Name  string     `yaml:"name"`
	Stops []StopTime `yaml:"stops"`
}

type StopTime struct {
	Id        ObjectID `yaml:"id"`
	Arrival   string   `yaml:"arrival,omitempty"`
	Departure string   `yaml:"departure,omitempty"`
}
