package main

import "fmt"

// There has got to be a more elegant way of doing this?
func find(target int) (x, y int) {
	n := 1
	x = 0
	y = 0

	distance := 1

	for {
		// Move `distance` units right.
		for i := 0; i < distance; i++ {
			x++
			n++
			// fmt.Printf("%v\t(%v,%v)\n", n, x, y)
			if n == target {
				return
			}
		}

		// Move `distance` units up.
		for i := 0; i < distance; i++ {
			y++
			n++
			// fmt.Printf("%v\t(%v,%v)\n", n, x, y)
			if n == target {
				return
			}
		}

		distance++

		// Move `distance` units left.
		for i := 0; i < distance; i++ {
			x--
			n++
			// fmt.Printf("%v\t(%v,%v)\n", n, x, y)
			if n == target {
				return
			}
		}

		// Move `distance` units down.
		for i := 0; i < distance; i++ {
			y--
			n++
			// fmt.Printf("%v\t(%v,%v)\n", n, x, y)
			if n == target {
				return
			}
		}

		distance++
	}
}

func main() {
	x, y := find(368078)
	if x < 0 {
		x = -x
	}
	if y < 0 {
		y = -y
	}
	fmt.Printf("moves: %v\n", x+y)
}
