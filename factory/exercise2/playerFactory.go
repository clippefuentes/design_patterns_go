package main

import "fmt"

type Player struct {
	name, role, region string
	age                int
}

// functional
func NewPlayerFactory(role, region string) func(name string, age int) *Player {
	return func(name string, age int) *Player {
		return &Player{name, role, region, age}
	}
}

type PlayerFactory struct {
	role, region string
}

func NewPlayerFactory2(role, region string) *PlayerFactory {
	return &PlayerFactory{role, region}
}

func (pf *PlayerFactory) Create(name string, age int) *Player {
	return &Player{name, pf.role, pf.region, age}
}

func main() {
	NAJunglerPlayerFactory := NewPlayerFactory("Jungler", "NA")
	naJungler1 := NAJunglerPlayerFactory("Spica", 19)
	fmt.Println(naJungler1)
	EUMidPlayerFactory := NewPlayerFactory2("Mid", "EU")
	euMid1 := EUMidPlayerFactory.Create("Perkz", 23)
	fmt.Println(euMid1)
}