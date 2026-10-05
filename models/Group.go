package models

type Group struct {
	Name string
	Description string
}

func NewGroup(name string, description string) *Group {
	return &Group{Name: name, Description: description}
}
