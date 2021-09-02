package main

import "fmt"

type hero struct {
	name         string
	refillPotion bool
	refillMana   bool
	refillWeapon bool
	refillFood   bool
}

type guild interface {
	execute(*hero)
	setNext(guild)
}

type potionShop struct {
	next guild
}

func (p *potionShop) execute(h *hero) {
	if h.refillPotion {
		fmt.Println("Hero has refill potion already")
		p.next.execute(h)
		return
	}
	fmt.Println("Refilling potions")
	h.refillWeapon = true
	p.next.execute(h)
}

func (p *potionShop) setNext(next guild) {
	p.next = next
}

type manaShop struct {
	next guild
}

func (m *manaShop) execute(h *hero) {
	if h.refillMana {
		fmt.Println("Mana is already full")
		m.next.execute(h)
		return
	}
	fmt.Println("Refilling mana")
	h.refillMana = true
	m.next.execute(h)
}

func (m *manaShop) setNext(next guild) {
	m.next = next
}

type 