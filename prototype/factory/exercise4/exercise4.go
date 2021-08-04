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

func (song *Song) DeepCopy() *Song {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	d := gob.NewDecoder(&b)
	result := Song{}
	e.Encode(song)
	d.Decode(&result)
	return &result
}

func NewSongCopy(proto *Song, title, name string) *Song {
	s := proto.DeepCopy()
	s.Title = title
	s.Singer.Name = name
	return s
}

var maleSong = &Song{
	"", &Singer{"", "Male"},
}

var femaleSong = &Song{
	"", &Singer{"", "Female"},
}

func MaleSong(title, name string) *Song {
	return NewSongCopy(maleSong, title, name)
}

func FemaleSong(title, name string) *Song {
	return NewSongCopy(femaleSong, title, name)
}

func main() {
	song1 := MaleSong("Golden", "Harry Styles")
	song2 := FemaleSong("Everything I wanted", "Billie")
	fmt.Println(song1, song1.Singer)
	fmt.Println(song2, song2.Singer)
}