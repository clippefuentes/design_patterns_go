package main

import "fmt"

type Clan struct {
	name, region, allegiance string
}

func (c *Clan) DeepCopy() *Clan {
	return &Clan{
		c.name,
		c.region,
		c.allegiance,
	}
}

type Hero struct {
	name   string
	clan   *Clan
	skills []string
}

func (h *Hero) DeepCopy() *Hero {
	q := *h
	q.clan = h.clan.DeepCopy()
	copy(q.skills, h.skills)
	return &q
}

func main() {
	hero1 := &Hero{
		"Hero",
		&Clan{
			"Light", "North", "Good",
		},
		[]string{"Light"},
	}

	hero2 := hero1.DeepCopy()
	hero2.name = "Hero2"
	hero2.clan.allegiance = "IDK"
	hero2.skills = append(hero1.skills, "Hero v2")
	fmt.Println(hero1, hero1.clan)
	fmt.Println(hero2, hero2.clan)

}
