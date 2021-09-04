package main

import "fmt"

func average(sf ...float64) float64 {
	total := 0.0
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}

func main() {
	n := average(43, 56, 87, 12, 45, 57)
	fmt.Println(n)
}

// package main

// import "fmt"

// func average(sf ...float64) float64 {
// 	total := 0.0
// 	for _, v := range sf {
// 		total += v
// 	}
// 	return total / float64(len(sf))
// }

// func main() {
// 	n := average(43, 56, 87, 12, 45, 57)
// 	fmt.Println(n)
// }

// package main

// import "fmt"

// func max(numbers ...int) int {
// 	var largest int
// 	for _, v := range numbers {
// 		if v > largest {
// 			largest = v
// 		}
// 	}
// 	return largest
// }

// func main() {
// 	greatest := max(4, 7, 9, 123, 543, 23, 435, 53, 125)
// 	fmt.Println(greatest)
// }
