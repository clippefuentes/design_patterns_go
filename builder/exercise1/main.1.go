package main

import (
	"fmt"
)

type Hero struct {
	name string
	race string
	//
	job string
	faction string
}

type HeroBuilder struct {
	hero *Hero
}

type HeroOriginBuilder struct {
	HeroBuilder
}

type HeroJobBuilder struct {
	HeroBuilder
}

func NewHeroBuilder() *HeroBuilder {
	return &HeroBuilder{&Hero{}}
}

func (it *HeroBuilder) Build() *Hero {
	return it.hero
}

func (it *HeroBuilder) Born() *HeroOriginBuilder {
	return &HeroOriginBuilder{*it}
}

func (it *HeroBuilder) Works() *HeroJobBuilder {
	return &HeroJobBuilder{*it}
}

// Origin Builder Methods

func (hob *HeroOriginBuilder) NameAs(name string) *HeroOriginBuilder {
	hob.hero.name = name
	return hob
}

func (hob *HeroOriginBuilder) BornAs(race string) *HeroOriginBuilder{
	hob.hero.race = race
	return hob
}

// Job Builder Methods

func (hjb *HeroJobBuilder) WorkAs(job string) *HeroJobBuilder {
	hjb.hero.job = job
	return hjb
}

func (hjb *HeroJobBuilder) JoinIn(faction string) *HeroJobBuilder {
	hjb.hero.faction = faction
	return hjb
}

func main() {
	hob := NewHeroBuilder()
	hob.
		Born().
			NameAs("Clyne").
			BornAs("Human").
		Works().
			WorkAs("Magician").
			JoinIn("The Dark Side")
	hero1 := hob.Build()
	fmt.Println(*hero1)
}