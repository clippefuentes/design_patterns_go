package main

import "fmt"

type Player struct {
	// Info
	name string
	role string
	// Team
	team         string
	jerseyNumber int
}

type PlayerBuilder struct {
	player *Player
}

type PlayerInfoBuilder struct {
	PlayerBuilder
}

type PlayerTeamInfoBuilder struct {
	PlayerBuilder
}

func (pb *PlayerBuilder) Build() *Player {
	return pb.player
}

func NewPlayerBuilder() *PlayerBuilder {
	return &PlayerBuilder{&Player{}}
}

func (pb *PlayerBuilder) Info() *PlayerInfoBuilder {
	return &PlayerInfoBuilder{*pb}
}

func (pb *PlayerBuilder) Joins() *PlayerTeamInfoBuilder {
	return &PlayerTeamInfoBuilder{*pb}
}

// Info Builder Methods
func (pib *PlayerInfoBuilder) Name(name string) *PlayerInfoBuilder {
	pib.player.name = name
	return pib
}

func (pib *PlayerInfoBuilder) Role(role string) *PlayerInfoBuilder {
	pib.player.role = role
	return pib
}

// Team Builder Methods

func (ptib *PlayerTeamInfoBuilder) Team(team string) *PlayerTeamInfoBuilder {
	ptib.player.team = team
	return ptib
}

func (ptib *PlayerTeamInfoBuilder) JerseyNumber(jerseyNumber int) *PlayerTeamInfoBuilder {
	ptib.player.jerseyNumber = jerseyNumber
	return ptib
}

func main() {
	pB := NewPlayerBuilder()
	pB.
		Info().
		Name("Lebron James").
		Role("Point Guard").
		Joins().
		Team("Los Angeles Lakers").
		JerseyNumber(23)

	p1 := pB.Build()
	fmt.Println(p1)
}