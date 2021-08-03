package main

import "fmt"

type Team struct {
	name, region string
}

func (t *Team) DeepCopy() *Team {
	return &Team{
		t.name,
		t.region,
	}
}

type Player struct {
	name      string
	age       int
	team      *Team
	playstyle []string
}

func (p *Player) DeepCopy() *Player {
	q := *p
	q.team = p.team.DeepCopy()
	copy(q.playstyle, p.playstyle)
	return &q
}

func main() {
	p1 := &Player{
		"Spica",
		19,
		&Team{
			"TSM", "NA",
		},
		[]string{},
	}

	p2 := p1.DeepCopy()
	p2.name = "Huni"
	p2.playstyle = append(p2.playstyle, "TOP")
	p1.playstyle = append(p1.playstyle, "JUNGLE")

	fmt.Println(p1, p1.team)
	fmt.Println(p2, p2.team)
}