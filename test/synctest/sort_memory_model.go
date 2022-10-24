package main

func init() {
	println("hello,world")
}

func main() {
	done := make(chan string)

	go func() {
		println("你好, 世界")
		done <- "卧槽"
	}()
	println(<-done)
}
