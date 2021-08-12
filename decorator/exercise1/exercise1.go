package main

import "fmt"

type Human struct{}
type Magician struct{}
type Hero struct{
	human *Human
	magic *Magician
}

func (*Human) DoHumanThings() {
	fmt.Println("Do Human Things")
}

func (*Magician) DoMagicianThings() {
	fmt.Println("Do Human Things")
}

func (h *Hero) DoHumanThings() {
	h.human.DoHumanThings()
	fmt.Println("But like a hero human")
}

func (h *Hero) DoMagicianThings() {
	h.magic.DoMagicianThings()
	fmt.Println("But like a hero magician")
}

func main() {
	h1 := &Human{}
	ma := &Magician{}
	hero :=  &Hero{h1, ma}
	hero.DoHumanThings()
	hero.DoMagicianThings()
}