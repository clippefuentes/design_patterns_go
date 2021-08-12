package main

import (
	"strings"
	"fmt"
)

type Todo struct {
	text string
}

type List struct {
	text string
	todos []TodoComponent
}

type TodoComponent interface {
	getText() string
}

func (l *List) insertTo(t *Todo) {
	l.todos = append(l.todos, t)
}

func (t *Todo) getText() string {
	return "<li>"+ t.text +"<li> \n"
}

func (l *List) getText() string {
	sb := strings.Builder{}
	sb.WriteString("<li> \n")
	sb.WriteString("\t" + l.text + "\n")
	sb.WriteString("\t<ul>\n")
	for _, v := range l.todos {
		sb.WriteString("\t\t"+v.getText())
	}
	sb.WriteString("\t</ul>\n")
	sb.WriteString("</li>")
	return sb.String()
}

func main() {
	t1 := &Todo{"Be Rich"}
	t2 := &Todo{"Make a lot of Money"}
	t1l := &Todo{"Help Mother Buy House"}
	t2l := &Todo{"Help my less fortune family"}
	l1 := &List{"Help", []TodoComponent{}}
	l2 := &List{"Must do", []TodoComponent{}}
	l3 := &List{"study", []TodoComponent{}}
	t31 := &Todo{"Study Go"}
	t32 := &Todo{"Crypto Trade"}
	l1.todos = append(l1.todos, t1l)
	l1.todos = append(l1.todos, t2l)
	l2.todos = append(l2.todos, t1)
	l2.todos = append(l2.todos, t2)
	l3.todos = append(l3.todos, t31)
	l3.todos = append(l3.todos, t32)
	l2.todos = append(l2.todos, l3)
	l1.todos = append(l1.todos, l2)

	fmt.Println(l1.getText())
}