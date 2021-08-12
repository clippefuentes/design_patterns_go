package main

import (
	"fmt"
	"strings"
)

type Human struct {
	name     string
	children []*Human
}

func (h *Human) print() {
	sb := strings.Builder{}

	sb.WriteString("This is \n")
	sb.WriteString(h.name + "\n")
	sb.WriteString("Childs are \n")
	for _, v := range h.children {
		sb.WriteString("\t")
		v.print()
	}

	fmt.Println(sb.String())
}

func main() {
	h1 := &Human{"John", []*Human{}}
	h2 := &Human{"Bea", []*Human{}}
	h3 := &Human{"BV", []*Human{h1}}
	h4 := &Human{"Test", []*Human{h3, h2}}
	h4.print()
}
