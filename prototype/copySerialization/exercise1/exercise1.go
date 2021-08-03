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


func main() {
	hero1 := &Hero{
		"Hero",
		&Clan{"Light", "North", "Good"},
		[]string{"Light"},
	}

	hero2 := hero1.DeepCopy()
	hero2.Name = "Hero2"
	hero2.Clan.Allegiance = "IDK"
	hero2.Skills = append(hero1.Skills, "Hero v2")
	fmt.Println(hero1, hero1.Clan)
	fmt.Println(hero2, hero2.Clan)

}
