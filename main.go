package main

import (
	"fmt"
	"math/rand"
	"time"
	"unicode/utf8"
)

func main() {
	_1()
	_if()
	_for()
	_switch()
}

func _1() {
	x := 10
	if x > 5 {
		fmt.Println(x)
		x := 5
		fmt.Println(x)
	}
	fmt.Println(x)
}

func _if() {
	rand.Seed(time.Now().UnixNano())
	/*n := rand.Intn(10)
	if n == 0 {
		fmt.Println("too small")
	} else if n > 5 {
		fmt.Println("too big")
	} else {
		fmt.Println("nice")
	}*/
	if n := rand.Intn(10); n == 0 {
		fmt.Println("too small")
	} else if n > 5 {

	}
}

func _for() {
	/*for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println(i, "3でも5でも割り切れる")
			continue
		}
		if i%3 == 0 {
			fmt.Println(i, "3で割り切れる")
			continue
		}
		if i%5 == 0 {
			fmt.Println(i, "5で割り切れる")
			continue
		}
	}*/
	evenVals := []int{1, 2}
	for i := 0; i < len(evenVals)-1; i++ {
		fmt.Println(evenVals[i])
	}
}

func _switch() {
	words := []string{"山", "川", "sun", "タコの足", "Let's Go", "Golang is god"}
	for _, word := range words {
		switch size := utf8.RuneCountInString(word); size {
		case 1, 2, 3, 4:
			fmt.Printf("%s : %d  ~5\n", word, size)
		case 5, 6, 7, 8, 9, 10:
			fmt.Printf("%s : %d 5~10\n", word, size)
		default:
			fmt.Printf("%s : %d too long\n", word, size)

		}
	}
}
