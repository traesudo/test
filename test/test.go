package main

import (
	"fmt"
	"test/tools"
)

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
	var num = []int{1, 2, 3, 4}
	var queue tools.Queue
	for i := 0; i < len(num); i++ {
		queue.Push(num[i])
	}
	var stack tools.Stack
	stack.Push(num)
	fmt.Println("woshi stack", stack.Pop())
	fmt.Println(queue)
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())

}
