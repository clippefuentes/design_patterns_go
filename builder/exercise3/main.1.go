package main

import "fmt"

type Car struct {
	// Car Desc
	make string
	year int
	// Car Origin
	country string
    brand string
}

type CarBuilder struct {
	car *Car
}

type CarDescriptionBuilder struct {
	CarBuilder
}

type CarOriginBuilder struct {
	CarBuilder
}

func NewCarBuilder() *CarBuilder {
	return &CarBuilder{&Car{}}
}

func (cb *CarBuilder) Description() *CarDescriptionBuilder {
	return &CarDescriptionBuilder{*cb}
}

func (cb *CarBuilder) Origin() *CarOriginBuilder {
	return &CarOriginBuilder{*cb}
}

// Description Builder Methods
func (cdb *CarDescriptionBuilder) Make(make string) *CarDescriptionBuilder {
	cdb.car.make = make
	return cdb
}

func (cdb *CarDescriptionBuilder) Year(year int) *CarDescriptionBuilder {
	cdb.car.year = year
	return cdb
}

// Origin Builder Methds
func (cob *CarOriginBuilder) Country(country string) *CarOriginBuilder {
	cob.car.country = country
	return cob
}

func (cob *CarOriginBuilder) Brand(brand string) *CarOriginBuilder {
	cob.car.brand = brand
	return cob
}

func (cb *CarBuilder) Build() *Car {
	return cb.car
}

func main() {
	cb := NewCarBuilder()

	cb.
		Description().
			Make("Corolla").
			Year(2020).
		Origin().
			Country("Canada").
			Brand("Toyoto")

	car1 := cb.Build()

	fmt.Println(car1)
}
