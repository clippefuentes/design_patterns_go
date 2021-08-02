package main

import "fmt"

type Hero struct {
	name, race, job string
}

// functional approach
func NewHeroFactory(race, job string) func(name string) *Hero {
	return func(name string) *Hero {
		return &Hero{name, race, job}
	}
}

type HeroFactory struct {
	race, job string
}

func NewHeroFactory2 (race, job string) *HeroFactory {
	return &HeroFactory{race, job}
}

func (hf *HeroFactory) Create (name string) *Hero {
	return &Hero{name, hf.race, hf.job}
}

func main() {
	HumanWarriorsHeroFactory := NewHeroFactory("Human", "Warrior")
	hw1 := HumanWarriorsHeroFactory("Jon Snow")
	fmt.Println(hw1)
	
	ElfArcherHeroFactory := NewHeroFactory2("Elf", "Archer")
	ea1 := ElfArcherHeroFactory.Create("Hawkeye")
	fmt.Println(ea1)
}