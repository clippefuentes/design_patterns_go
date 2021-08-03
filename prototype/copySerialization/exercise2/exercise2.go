package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Team struct {
	Name, Region string
}

type Player struct {
	Name      string
	Age       int
	Team      *Team
	PlayStyle []string
}

func (p *Player) DeepCopy() *Player {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	e.Encode(p)
	d := gob.NewDecoder(&b)
	result := Player{}
	d.Decode(&result)
	return &result
}

func main() {
	p1 := &Player{
		"Spica",
		19,
		&Team{
			"TSM", "NA",
		},
		[]string{},
	}

	p2 := p1.DeepCopy()
	p2.Name = "Huni"
	p2.PlayStyle = append(p2.PlayStyle, "TOP")
	p1.PlayStyle = append(p1.PlayStyle, "JUNGLE")

	fmt.Println(p1, p1.Team)
	fmt.Println(p2, p2.Team)
}