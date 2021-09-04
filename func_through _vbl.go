package main

import "fmt"

func greeting() {

	fmt.Println("hello world")
}

func main() {

	greeting()
}

// package main

// import "fmt"

// func main() {

// 	greeting := func() {

// 		fmt.Println("hello world")
// 	}
// 	greeting()
// }

// package main

// import "fmt"

// func main() {

// half := func(n int ) ( int , bool) {
//      return n/2 , n%2 == 0
// }
//   fmt.Println(half(2))
// }

// package main

// import "fmt"

// func main() {

// 	add := func(x, y int) int {
// 		return x + y
// 	}
// 	fmt.Println(add(1, 4))
// }

//closure
// package main
// import "fmt"

// func main() {

// x := 0
//  increment := func() int {
//     x++
//     return x
//  }

//  fmt.Println(increment())
//  fmt.Println(increment())
// }
// scope of x is func main
// func incremet has access to x

// returning a function
// package main

// import "fmt"

// func makeEvenGenerator() func() int {

// 	i := 0
// 	return func() int {
// 		i += 2
// 		return i
// 	}
// }

// func main() {

// 	nextEven := makeEvenGenerator()
// 	fmt.Println(nextEven()) //2
// 	fmt.Println(nextEven()) //4
// 	fmt.Println(nextEven()) //6

// 	masEven := makeEvenGenerator()
// 	fmt.Println(masEven()) //2
// 	fmt.Println(masEven()) //4
// 	fmt.Println(masEven()) //6
// }

//callback

// package main

// import "fmt"

// func visit(numbers []int, callback func(int)) {

// 	for _, n := range numbers {

// 		callback(n)
// 	}
// }

// func main() {

// 	visit([]int{1, 2, 3, 4}, func(n int) {

// 		fmt.Println(n)
// 	})
// }
