package main

type catalog interface {
	shipping() float64
	tax() float64
}

type discount interface {
	offer() float64
}

type configurable struct {
	name       string
	price, qty float64
}

func (c *configurable) tax() float64 {
	return c.price * c.qty * 0.05
}

func (c *configurable) shipping() float64 {
	return c.qty * 5
}
