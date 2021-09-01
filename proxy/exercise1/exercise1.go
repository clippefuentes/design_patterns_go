package main

import "fmt"

type MagicPower interface {
	useMagic()
}

type Magician struct {
	name string
}

func (m *Magician) useMagic() {
	fmt.Println("Use Magic", m.name)
}

type LazyMagician struct {
	name     string
	magician *Magician
}

func (l *LazyMagician) useMagic() {
	if l.magician == nil {
		l.magician = &Magician{name: l.name}
	}
	l.magician.useMagic()
}

func useMagicPowers(magicPower MagicPower) {
	fmt.Println("About to use powers")
	magicPower.useMagic()
	fmt.Println("Done using powers")
}

func main() {
	bmp := &LazyMagician{name: "Clyne"}
	useMagicPowers(bmp)
}
