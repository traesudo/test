package test

import "fmt"

type Animal struct {
	Dog
}
type Dog struct {
	index string
}

func (c *Dog) Prepare() {
	c.index = "ok"
}

func (c *Animal) Write() {
	fmt.Println(c.index)
}
