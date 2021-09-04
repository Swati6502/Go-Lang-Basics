package main

import (
	"fmt"
)

//function golang
func main() {
	local := 10
	show(local)

	incrWrong(local)
	fmt.Printf("main -> local = %d\n", local)
}

func show(n int) {
	fmt.Printf("show -> n = %d\n", n)
}

func incrWrong(n int) {

	// n:= main's local variable's value
	n++
}
