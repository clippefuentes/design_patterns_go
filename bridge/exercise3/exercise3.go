package main

import "fmt"

// Abstraction: Car
// Implementation: Features

type CarFeatures interface {
	useFeatures(make string)
}

type TeslaFeatures struct {}

func (tf *TeslaFeatures) useFeatures(make string) {
	tf.useAutoPilot(make)
}

func (tf *TeslaFeatures) useAutoPilot(make string) {
	fmt.Println(make, "use autopilot drive")
}

type ToyotaFeatures struct {}

func (tf *ToyotaFeatures) useFeatures(make string) {
	tf.useTurbo(make)
}

func (tf *ToyotaFeatures) useTurbo(make string) {
	fmt.Println(make, "use Turbo")
}

type Car struct {
	carFeatures CarFeatures
	make string
}

func (c *Car) useCar() {
	c.carFeatures.useFeatures(c.make)
}

func main() {
	teslaF := &TeslaFeatures{}
	toyotaF := &ToyotaFeatures{}
	car1 := &Car{teslaF, "G1"}
	car2 := &Car{toyotaF, "Corolla"}
	car1.useCar()
	car2.useCar()
}