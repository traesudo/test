package main

import "fmt"

func main() {
	x := [3]int{1, 2, 3}
	//func ok()
	ok(x)
	fmt.Println(x)
}

func ok(arr [3]int) {
	arr[0] = 7
	fmt.Println(arr)

}
