package entity

type stats struct {
	strength     int
	dexterity    int
	constitution int
	intelligence int
	magic        int
	spirit       int
}

type Player struct {
	id        int
	name      string
	hp        int
	mp        int
	level     int
	exp       int
	stats     stats
	abilities []Ability
}
