package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)


type Owner struct {
	Name, Work string
	Salary     int
}

type Car struct {
	Brand, Make string
	Owner       *Owner
	Features    []string
}

func (c *Car) DeepCopy() *Car {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	d := gob.NewDecoder(&b)
	result := Car{}
	e.Encode(c)
	d.Decode(&result)
	return &result
}

func main() {
	car1 := &Car{
		"Toyota", "86",
		&Owner{
			"Clyne", "SE", 65000,
		},
		[]string{"Blue", "Four Wheels"},
	}

	car2 := car1.DeepCopy()
	car2.Make = "Supra"
	car2.Owner.Work = "Trader"
	car2.Features = append(car2.Features, "Turbo")

	fmt.Println(car1, car1.Owner)
	fmt.Println(car2, car2.Owner)
}