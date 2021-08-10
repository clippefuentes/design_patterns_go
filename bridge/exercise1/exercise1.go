package main

import "fmt"

// Hero Bridge Pattern Exercise

// ABSTRACTION HERO
// IMPLEMENTATION POWERS

type Superpower interface {
	usePower(name string)
}

type DarkPowerHero struct {}

func (d *DarkPowerHero) usePower(name string) {
	fmt.Println("Use Dark Powers:", name)
}

type LightPowerHero struct {}

func (l *LightPowerHero) usePower(name string) {
	fmt.Println("Use Light Powers:", name)
}

type Hero struct {
	superPower Superpower
	name string
}

func (h *Hero) useHeroPowers () {
	h.superPower.usePower(h.name)
}

func NewHero(superPower Superpower, name string) *Hero {
	return &Hero{superPower: superPower, name: name}
}

func main() {
	darkSuperpower := &DarkPowerHero{}
	lightSuperpower := &LightPowerHero{}
	hero1 := NewHero(darkSuperpower, "dark1")
	hero2 := NewHero(lightSuperpower, "light2")
	hero1.useHeroPowers()
	hero2.useHeroPowers()
}