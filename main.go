package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	fmt.Println("1 :")
	_1()
	fmt.Println("3 :")
	_3()
	fmt.Println("5 :")
	_5()
}

// ###############################################
// 2.1
func _1() {
	x := complex(2.5, 3.1)
	y := complex(10.2, 2)
	fmt.Println(x + y)
	fmt.Println(x - y)
	fmt.Println(x * y)
	fmt.Println(real(x))
	fmt.Println(imag(x))
	fmt.Println(cmplx.Abs(x))
}

// ###############################################
// 2.3
const x = 20
const (
	idKey   = "id"
	nameKey = "name"
)
const z = 20 * 10

func _3() {
	fmt.Println(x)
}

//###############################################
//2.5

func _5() {
	_0 := 0_0
	a := "hello"
	あ := "this is also available"
	fmt.Println(_0)
	fmt.Println(a)
	fmt.Println(あ)
}
