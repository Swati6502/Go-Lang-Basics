package main

import (
	"fmt"
	"reflect"
)

func main() {

	var name = "Sangam"
	var age = 25

	fmt.Println(reflect.TypeOf(name))
	fmt.Println(reflect.TypeOf(age))

	fmt.Printf("%s is %d years old\n", name, age)
}
