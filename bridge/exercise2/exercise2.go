package main

import "fmt"

// ABSTRACTION: Hero
// IMPLEMENTATION: Role

type RoleHero interface {
	useSkill(name, target string)
}

type JunglerHero struct {}

func (j *JunglerHero) useSkill(name, target string) {
	fmt.Println(name, "uses Smite on", target)
}

type MagicHero struct {}

func (m *MagicHero) useSkill(name, target string) {
	fmt.Println(name, "burst", target)
}

type Hero struct {
	roleHero RoleHero
	name string
}

func (h *Hero) attack(name string) {
	h.roleHero.useSkill(h.name, name)
}

func main() {
	jH := &JunglerHero{}
	mH := &MagicHero{}
	h1 := &Hero{jH, "Zac"}
	h2 := &Hero{mH, "Ziggs"}
	h1.attack("Lee")
	h2.attack("Ori")
}