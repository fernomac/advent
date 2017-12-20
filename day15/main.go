package main

import "fmt"

func nextA(cur int) int {
	for {
		cur = (cur * 16807) % 2147483647
		if cur%4 == 0 {
			return cur
		}
	}
}

func nextB(cur int) int {
	for {
		cur = (cur * 48271) % 2147483647
		if cur%8 == 0 {
			return cur
		}
	}
}

func main() {
	// a := 65
	// b := 8921

	a := 883
	b := 879

	count := 0

	for i := 0; i < 5*1000*1000; i++ {
		a = nextA(a)
		b = nextB(b)
		if (a & 0xFFFF) == (b & 0xFFFF) {
			count++
		}
	}

	fmt.Println(count)
}
