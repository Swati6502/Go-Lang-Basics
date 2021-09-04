package main

import "fmt"

func main() {

	const (
		monday = iota
		tuesday
		wednesday
		thursday
		friday
		saturday
		sunday
	)

	fmt.Println(monday, tuesday, wednesday, thursday, friday, saturday, saturday, sunday)
}

// package main

// import "fmt"

// func main(){

// 	const (
// 		monday = iota + 1
// 		tuesday
// 		wednesday
// 		thursday
// 		friday
// 		saturday
// 		sunday

// 	)

// fmt.Println(monday,tuesday,wednesday,thursday,friday,saturday,sunday)
// }

//strconv
// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func main() {

// 	s := strconv.Itoa(42)
// 	fmt.Println(s)

// }
