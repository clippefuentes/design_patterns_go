package main

import "fmt"

type Vehicle struct {
	brand string
	year int
}

type WheelVehicle interface {
	DriveOnEarth()
}

type Car struct {
	Vehicle
	wheels int
	make  string
}

type Aircraft interface {
	fly()
}

type Airplane struct {
	Vehicle
	airline string
	wings int
}

func (car *Car) DriveOnEarth() {
	fmt.Printf("Drive the %v %v %v Car that has %d wheels \n", car.year, car.make, car.brand, car.wheels)
	fmt.Println("HAS WHEELS")
}

func (plane *Airplane) fly() {
	fmt.Printf("Fly %v of %v that has %v wings", plane.brand, plane.airline, plane.wings)
}


func changeWheel(wh WheelVehicle) {
	wh.DriveOnEarth()
}

func flyAircraft(ac Aircraft) {
	ac.fly()
}

func NewVehicleCompany(brand string, year int) *Vehicle {
	return &Vehicle{brand, year}
}

func (nv *Vehicle) MakePlane(airline string, wing int) *Airplane {
	return &Airplane{
		*nv, airline, wing,
	}
}

func (nv *Vehicle) MakeCar(wheels int, make string) *Car {
	return &Car {
		*nv, wheels, make,
	}
}


func main() {
	// car1 := &Car{
	// 	Vehicle{"Toyota", "Corolla", 2020},
	// 	4, "Blue",
	// }

	// fmt.Println(car1)
	// fmt.Println(&car1)
	// *&car1.brand = "Nissan"
	// *&car1.make = "GTR"
	// fmt.Println(car1)
	// fmt.Println(&car1)

	// plane1 := &Airplane{
	// 	Vehicle{"Air Canada Brand", "G1", 0},
	// 	"Air Canada Airline", 2,
	// }
	// changeWheel(car1)
	// flyAircraft(plane1)
	clyneCoorporation := NewVehicleCompany("Clyne", 2022)
	car2 := clyneCoorporation.MakeCar(4, "Excel Car")
	changeWheel(car2)
	plane2 := clyneCoorporation.MakePlane("Clippe Fly Coop", 3)
	flyAircraft(plane2)
}
