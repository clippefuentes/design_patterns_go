package main

import "fmt"

type Human struct {
	name           string
	SchoolDone     bool
	wife           string
	money          int
}

type Modifier interface {
	Add(Modifier)
	Handle()
}

type HumanModifer struct {
	human *Human
	next  Modifier
}

func (h *HumanModifer) Add(m Modifier) {
	if h.next != nil {
		h.next.Add(m)
	} else {
		h.next = m
	}
}

func (h *HumanModifer) Handle() {
	if h.next != nil {
		h.next.Handle()
	}
}

type SchoolHandler struct {
	HumanModifer
}

func (p SchoolHandler) Handle() {
	if p.human.SchoolDone == false {
		fmt.Println("learning abc 123")
		p.human.SchoolDone = true
	}
	p.HumanModifer.Handle()
}

func KinderGarden(h *Human) *SchoolHandler {
	return &SchoolHandler{HumanModifer{human: h}}
}


type WifeHandler struct {
	HumanModifer
	wife string
}

func (p WifeHandler) Handle() {
	if p.human.wife == "" {
		fmt.Println("Date ", p.wife)
		p.human.wife = p.wife
	}
	p.HumanModifer.Handle()
}

func MatchMaker(h *Human, wife string) *WifeHandler {
	return &WifeHandler{HumanModifer{human: h}, wife}
}

type MoneyHandler struct {
	HumanModifer
	money int
}

func (p MoneyHandler) Handle() {
	if p.human.money == 0 {
		fmt.Println("Put bank ", p.money)
		p.human.money = p.money
	}
	p.HumanModifer.Handle()
}

func Lotto(h *Human, money int) *MoneyHandler {
	return &MoneyHandler{HumanModifer{human: h}, money}
}

func main() {
	human1 := &Human{name: "Clyne"}
	fmt.Println(human1)
	root := HumanModifer{human: human1}
	root.Add(KinderGarden(human1))
	root.Add(MatchMaker(human1, "Sheena"))
	root.Add(Lotto(human1, 50000000))
	root.Handle()
	fmt.Println(human1)
}