package main

type Food interface {
	setColor(int)
	setPrice(int)
}

type Apple struct {
	price int
	color int
}

func (a Apple) setColor(cc int) {
	a.color = cc
}
func (a *Apple) setPrice(pp int) {
	a.price = pp
}
