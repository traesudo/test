package main

// first.c
import "C"

func main() {
	C.SayHello(C.CString("Hello, World\n"))
}
