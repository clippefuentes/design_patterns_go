package main

import "fmt"

type Pizza struct {
	ownsTo     string
	doughReady bool
	slices     int
	toppings   []string
}

type Modifier interface {
	Add(Modifier)
	Handle()
}

type PizzaModifier struct {
	pizza *Pizza
	next  Modifier
}

func (p *PizzaModifier) Add(m Modifier) {
	if p.next != nil {
		p.next.Add(m)
	} else {
		p.next = m
	}
}

func (p *PizzaModifier) Handle() {
	if p.next != nil {
		p.next.Handle()
	}
}

type DoughModifier struct {
	PizzaModifier
}

func (d DoughModifier) Handle() {
	if d.pizza.doughReady == false {
		fmt.Println("Dough Ready")
		d.pizza.doughReady = true
	}
	d.PizzaModifier.Handle()
}

func Dough(p *Pizza) *DoughModifier {
	return &DoughModifier{PizzaModifier{pizza: p}}
}

type SliceModifier struct {
	PizzaModifier
	slice int
}

func (s SliceModifier) Handle() {
	s.pizza.slices = s.slice
	fmt.Printf("Slicing pizza into %d pcs/pc", s.slice)
	s.PizzaModifier.Handle()
}

func Slice(p *Pizza, slice int) *SliceModifier {
	return &SliceModifier{PizzaModifier{pizza: p}, slice}
}

type ToppingModifier struct {
	PizzaModifier
	topping []string
}

func (t ToppingModifier) Handle() {
	for _, topping := range t.topping {
		fmt.Println("Putting " + topping + " into it")
		t.pizza.toppings = append(t.pizza.toppings, topping)
	}
	t.PizzaModifier.Handle()
}

func Topper(p *Pizza, toppings ...string) *ToppingModifier {
	return &ToppingModifier{PizzaModifier{pizza: p}, toppings}
}

func main() {
	pizza1 := &Pizza{ownsTo: "Clyne"}
	fmt.Println(pizza1)
	pizzaOneProcess := PizzaModifier{pizza: pizza1}
	pizzaOneProcess.Add(Slice(pizza1, 4))
	pizzaOneProcess.Add(Dough(pizza1))
	pizzaOneProcess.Add(Topper(pizza1, "cheese", "pepporoni", "mushroom", "pineapple"))
	pizzaOneProcess.Handle()
	fmt.Println(pizza1)
}