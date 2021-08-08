package main

import "fmt"

// Hero Adapter Pattern Exercise

// client: HeroAssociation
// service: Hero
// unknown: Citizen
// Adaptor: HeroCitizenAdaptor

type HeroAssociation struct {}

type SuperHero struct {}

type SimpleHuman struct{}

type SimpleHumanAdapter struct {
	simpleHuman *SimpleHuman
}

type Hero interface {
	usePower()
}

func (h *HeroAssociation) useHeroPower(hero Hero) {
	hero.usePower()
}

func (hero *SuperHero) usePower() {
	fmt.Println("I am a super so I will use super power")
}

func (human *SimpleHuman) fight() {
	fmt.Println("I dont have super powers but i can fight")
}

func (simpleHumanHero *SimpleHumanAdapter) usePower() {
	simpleHumanHero.simpleHuman.fight()
	fmt.Println("But now i have adapted some powers")
}

func main() {
	ha := &HeroAssociation{}
	shero := &SuperHero{}
	shuman := &SimpleHuman{}
	shumanhero := &SimpleHumanAdapter{shuman}
	ha.useHeroPower(shero)
	shuman.fight()
	ha.useHeroPower(shumanhero)
}