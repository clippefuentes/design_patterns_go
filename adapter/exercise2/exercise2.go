package main

import "fmt"

// Player Adapter Pattern Exercise

// client: Game
// service: Carry
// unknown: Support
// Adaptor: CarrySupportAdaptor

type Game struct {}
type CarryHero struct {}
type SupportHero struct {}
type CarrySupportAdapter struct {
	supportHero *SupportHero
}

type Carry interface {
	inflictDamage()
}

func (g *Game) clash(c Carry) {
	c.inflictDamage()
}

func (ch *CarryHero) inflictDamage() {
	fmt.Println("Carry Do Damage")
}

func (sh *SupportHero) heal() {
	fmt.Println("Heal Teammates")
}

func (cs *CarrySupportAdapter) inflictDamage() {
	cs.supportHero.heal()
	fmt.Println("AND CURSE ENEMIES!!")
}

func main() {
	game := &Game{}
	ch := &CarryHero{}
	sh := &SupportHero{}
	csh := &CarrySupportAdapter{sh}
	game.clash(ch)
	sh.heal()
	fmt.Println("-----")
	game.clash(csh)
}