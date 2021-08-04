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

func NewPlayerCopy(proto *Player ,name string, age int) *Player {
	p := proto.DeepCopy()
	p.Name = name
	p.Age = age
	return p
}

var TSMJunglerPlayer = &Player{
	"", 0, &Team{"TSM", "NA"}, []string{"Jungler"},
}

var C9TopPlayer = &Player{
	"", 0, &Team{"C9", "NA"}, []string{"Top"},
}

func NewTSMJunglePlayer(name string, age int) *Player {
	return NewPlayerCopy(TSMJunglerPlayer, name, age)
}

func NewC9TopPlayer(name string, age int) *Player{
	return NewPlayerCopy(C9TopPlayer, name, age)
}

func main() {
	p1 := NewTSMJunglePlayer("POE", 25)
	p2 := NewC9TopPlayer("Zven", 21)
	fmt.Println(p1, p1.Team)
	fmt.Println(p2, p2.Team)
}