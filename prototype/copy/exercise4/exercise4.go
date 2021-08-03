package main

import "fmt"

type Singer struct {
	name, gender string
}

type Song struct {
	title  string
	singer *Singer
}

func (s *Singer) DeepCopy() *Singer {
	return &Singer{
		s.name,
		s.gender,
	}
}

func (s *Song) DeepCopy() *Song {
	q := *s
	q.singer = s.singer.DeepCopy()
	return &q
}

func main() {
	song := &Song{
		"Somebody that i use to know",
		&Singer{"Gotye", "Male"},
	}

	song2 := song.DeepCopy()
	song2.singer.name = "Maroon 5"

	fmt.Println(song, song.singer)
	fmt.Println(song2, song2.singer)

}