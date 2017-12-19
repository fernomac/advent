package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

//           (0,2)
//     (-1,1)  |  (1,1)
// (-2,0) |  (0,0) |  (2,0)
//    (-1,-1)  |  (1,-1)
//           (0,-2)

func main() {
	bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	x, y := 0, 0
	steps := strings.Split(string(bytes), ",")

	for _, step := range steps {
		switch step {
		case "n":
			y += 2
		case "s":
			y -= 2
		case "nw":
			x--
			y++
		case "ne":
			x++
			y++
		case "sw":
			x--
			y--
		case "se":
			x++
			y--
		default:
			panic(step)
		}
	}

	count := 0
	for x != 0 || y != 0 {
		if x == 0 {
			// Just move in the y direction.
			if y > 0 {
				y -= 2
			} else {
				y += 2
			}
		} else {
			// Some of both.
			if x > 0 {
				x--
			} else {
				x++
			}
			if y > 0 {
				y--
			} else {
				y++
			}
		}
		count++
	}

	fmt.Println(count)
}
