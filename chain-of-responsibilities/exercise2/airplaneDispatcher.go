package main

import "fmt"

type airplane struct {
	code            int
	gas             int
	statusCheck     bool
	nextDestination string
}

type Modifier interface {
	Add(Modifier)
	Handle()
}

type PlaneModifer struct {
	plane *airplane
	next  Modifier
}

func (p *PlaneModifer) Add(m Modifier) {
	if p.next != nil {
		p.next.Add(m)
	} else {
		p.next = m
	}
}

func (p *PlaneModifer) Handle() {
	if p.next != nil {
		p.next.Handle()
	}
}

type GasModifier struct {
	PlaneModifer
}

func (g *GasModifier) Handle() {
	if g.plane.gas <= 10 {
		fmt.Println("Refilling Gas")
		g.plane.gas = 500
	}
	g.PlaneModifer.Handle()
}

func GasStation(p *airplane) *GasModifier {
	return &GasModifier{PlaneModifer{plane: p}}
}

type StatusHandler struct {
	PlaneModifer
}

func (s *StatusHandler) Handle() {
	if s.plane.statusCheck == false {
		fmt.Println("Check Engine and Plane parts")
		s.plane.statusCheck = true
	}
	s.PlaneModifer.Handle()
}

func MechanicStation(p *airplane) *StatusHandler {
	return &StatusHandler{PlaneModifer{plane: p}}
}

type DestinationHandle struct {
	PlaneModifer
	newDestination string
}

func (d *DestinationHandle) Handle() {
	fmt.Println("Change Destination to", d.newDestination)
	d.plane.nextDestination = d.newDestination
	d.PlaneModifer.Handle()
}

func DestinationManagment(p *airplane, destination string) *DestinationHandle {
	return &DestinationHandle{PlaneModifer{plane: p}, destination}
}

func main() {
	plane1 := &airplane{code: 111}
	fmt.Println(plane1)
	root := &PlaneModifer{plane: plane1}
	root.Add(GasStation(plane1))
	root.Add(MechanicStation(plane1))
	root.Add(DestinationManagment(plane1, "Philippines"))
	root.Handle()
	fmt.Println(plane1)
}