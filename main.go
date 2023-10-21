package main

import "fmt"

func main() {
	_array()
	_slice()
	_map()
	_struct()
}

func _array() {
	var x = [...]int{1, 2, 3, 4, 5}
	fmt.Println("this is array : ", x)
}

func _slice() {
	var x = []int{1, 2, 3, 4, 5}
	fmt.Println("this is slice : ", x)
	var y = []int{1, 5: 4, 6, 10: 100, 15}
	fmt.Println("this is also slice", y)
	x = append(x, 10, 2, 2, 2, 2, 3)
	fmt.Println("append to x slice", x)
	s := []int{1, 2, 3, 4}
	s_2 := x[:2]
	s1_ := x[1:]
	s1_3 := x[1:3]
	fmt.Println("slice's slice")
	fmt.Println("s:", s)
	fmt.Println("s_2:", s_2)
	fmt.Println("s1_", s1_)
	fmt.Println("s1_3", s1_3)
}

func _map() {
	var nilMap map[string]int
	fmt.Println("NilMap's len : ", len(nilMap))
	totalWins := map[string]int{}
	fmt.Println(totalWins == nil)
	totalWins["abc"] = 3
	fmt.Println("totalwins[abc] : ", totalWins["abc"])
	teams := map[string][]string{
		"a": {"John", "Bob"},
		"b": {"Mike July"},
		"c": {"Sindy", "Ann"},
	}
	fmt.Println(teams)
	v, ok := totalWins["abc"]
	fmt.Println(",Ok idiom", v, ok)
	delete(totalWins, "abc")
	v, ok = totalWins["abc"]
	fmt.Println("delete abc", v, ok)
}

func _struct() {
	type person struct {
		name string
		age  int
		pet  string
	}
	julia := person{
		age:  30,
		name: "sindy",
	}
	fmt.Println("julia's name : ", julia.name)
}
