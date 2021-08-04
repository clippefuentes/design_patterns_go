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

func NewCarCopy(proto *Car, brand, make string) *Car {
	c := proto.DeepCopy()
	c.Brand = brand
	c.Make = make
	return c
}

var clyneEngineerCar = &Car{
	"", "", &Owner{"Clyne", "Engineer", 55000}, []string{"4 Wheels"},
} 

var clippeTraderCar = &Car{
	"", "", &Owner{"Clyne", "Trader", 100000}, []string{"Electric"},
} 

func NewClyneCar(brand, make string) *Car{
	return NewCarCopy(clyneEngineerCar, brand, make)
}

func NewClippeCar(brand, make string) *Car{
	return NewCarCopy(clippeTraderCar, brand, make)
}

func main() {
	car1 := NewClyneCar("Toyota", "Supra")
	car2 := NewClippeCar("Nissan", "GTR")
	fmt.Println(car1, car1.Owner)
	fmt.Println(car2, car2.Owner)
}