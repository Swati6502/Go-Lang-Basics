package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("%#v\n", os.Args)
}

//more about os

// package main
// import (
// 	"fmt"
//      "os"
// )
// func main() {
//   fmt.Printf("%#v\n",os.Args)

//   fmt.Println("Path", os.Args[0])
//   fmt.Println("1st argument",os.Args[1])
//   fmt.Println("2nd argument",os.Args[2])
//   fmt.Println("3rd argument",os.Args[3])
//   fmt.Println("number of items inside os.Args:",len(os.Args))
// }

//os, strings
// package main

// import (
// 	"fmt"
// 	"os"
// 	"strings"
// )

// func main(){

// 	msg := os.Args[1]
// 	l := len(msg)
// 	s := msg + strings.Repeat("!",l)
// 	fmt.Println(s)

// }
