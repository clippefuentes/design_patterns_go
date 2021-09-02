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
	// m.next.execute(h)s
}

func (m *manaShop) setNext(next guild) {
	m.next = next
}

type weaponShop struct {
	next guild
}

func (w *weaponShop) execute(h *hero) {
	if h.refillWeapon {
		fmt.Println("Weapons are already ready")
		w.next.execute(h)
		return
	}
	fmt.Println("Arrow is refill, Sword is sharped and shield is polished")
	h.refillWeapon = true
	w.next.execute(h)
}

func (w *weaponShop) setNext(next guild) {
	w.next = next
}

type foodShop struct {
	next guild
}

func (f *foodShop) execute(h *hero) {
	if h.refillFood {
		fmt.Println("Food is refilled")
		f.next.execute(h)
		return
	}
	fmt.Println("refilling food bag")
	h.refillFood = true
	f.next.execute(h)
}


func (f *foodShop) setNext(next guild) {
	f.next = next
}

func main () {
	// refillPotion bool
	// refillMana   bool
	// refillWeapon bool
	// refillFood   bool
	manaShop := &manaShop{}
	potionShop := &potionShop{}

	potionShop.setNext(manaShop)

	weaponShop := &weaponShop{}
	weaponShop.setNext(potionShop)
	weaponShop.setNext(potionShop)
	weaponShop.setNext(potionShop)
	weaponShop.setNext(potionShop)
	foodShop := &foodShop{}

	foodShop.setNext(weaponShop)
	hero := &hero{name: "Clyne"}
	foodShop.execute(hero)
}

// func main() {

//     cashier := &cashier{}

//     //Set next for medical department
//     medical := &medical{}
//     medical.setNext(cashier)

//     //Set next for doctor department
//     doctor := &doctor{}
//     doctor.setNext(medical)

//     //Set next for reception department
//     reception := &reception{}
//     reception.setNext(doctor)

//     patient := &patient{name: "abc"}
//     //Patient visiting
//     reception.execute(patient)
// }