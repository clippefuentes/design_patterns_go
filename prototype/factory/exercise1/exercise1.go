package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Clan struct {
	Name, Region, Allegiance string
}

type Hero struct {
	Name   string
	Clan   *Clan
	Skills []string
}

func (h *Hero) DeepCopy() *Hero {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	e.Encode(h)
	d := gob.NewDecoder(&b)
	result := Hero{}
	d.Decode(&result)
	return &result
}


func NewHero(proto *Hero, name string) *Hero {
	newHero := proto.DeepCopy()
	newHero.Name = name
	return newHero
}

var lightClanHero = &Hero{
	"",
	&Clan{"Light", "North", "Human"},
	[]string{"Heal"},
}

var darkClanHero = &Hero{
	"",
	&Clan{"Dark", "NA", "Orc"},
	[]string{"Dark Magic"},
}

func NewLightNorthClanHeroCopy(name string) *Hero {
	return NewHero(lightClanHero, name)
}

func NewDarkNAClanHeroCopy(name string) *Hero {
	return NewHero(darkClanHero, name)
}

func main() {
	hero1 := NewLightNorthClanHeroCopy("Hero 1")
	fmt.Println(hero1, hero1.Clan)
}
