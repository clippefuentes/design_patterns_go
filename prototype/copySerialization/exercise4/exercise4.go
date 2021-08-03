package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)


type Singer struct {
	Name, Gender string
}

type Song struct {
	Title  string
	Singer *Singer
}

func (s *Song) DeepCopy() *Song {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	d := gob.NewDecoder(&b)
	e.Encode(s)
	result := Song{}
	d.Decode(&result)
	return &result
}

func main() {
	song := &Song{
		"Somebody that i use to know",
		&Singer{"Gotye", "Male"},
	}

	song2 := song.DeepCopy()
	song2.Singer.Name = "Maroon 5"

	fmt.Println(song, song.Singer)
	fmt.Println(song2, song2.Singer)

}