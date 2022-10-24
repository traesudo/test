package main

import "fmt"

type Person struct {
	name string
}

func (ok *Person) write() {
	fmt.Println("it is write ")
}
func (lx *Person) read() {
	fmt.Println("it is read ")
}

func main() {
	var ok Person
	ok.write()
	ok.read()

}
