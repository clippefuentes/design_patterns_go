package main

import "fmt"

type Owner struct {
	name, work string
	salary     int
}

type Car struct {
	brand, make string
	owner       *Owner
	features    []string
}

func (o *Owner) DeepCopy() *Owner {
	return &Owner{
		o.name,
		o.work,
		o.salary,
	}
}

func (c *Car) DeepCopy() *Car {
	q := *c
	q.owner = c.owner.DeepCopy()
	copy(q.features, c.features)
	return &q
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
	car2.make = "Supra"
	car2.owner.work = "Trader"
	car2.features = append(car2.features, "Turbo")

	fmt.Println(car1, car1.owner)
	fmt.Println(car2, car2.owner)
}