package main

import "fmt"
import "strconv"

func hashem(input string) [][]byte {
	result := make([][]byte, 0)
	for row := 0; row < 128; row++ {
		result = append(result, knothash(input+"-"+strconv.Itoa(row)))
	}
	return result
}

func bitmask(data []byte, offset int) bool {
	b := data[offset/8]
	mask := byte(1 << uint(7-(offset%8)))
	return ((b & mask) == mask)
}

func usage(hashes [][]byte) int {
	used := 0
	for row := 0; row < 128; row++ {
		hash := hashes[row]
		for column := 0; column < 128; column++ {
			if bitmask(hash, column) {
				used++
			}
		}
	}
	return used
}

type point struct {
	x, y int
}

func (p point) valid() bool {
	return (p.x >= 0 && p.x < 128 && p.y >= 0 && p.y < 128)
}

func find(hashes [][]byte) point {
	for row := 0; row < 128; row++ {
		for column := 0; column < 128; column++ {
			if bitmask(hashes[row], column) {
				return point{column, row}
			}
		}
	}
	return point{-1, -1}
}

func used(hashes [][]byte, pt point) bool {
	return bitmask(hashes[pt.y], pt.x)
}

func clear(hashes [][]byte, pt point) {
	data := hashes[pt.y]
	idx := pt.x / 8
	mask := ^byte(1 << uint(7-(pt.x%8)))
	data[idx] &= mask
}

func regions(hashes [][]byte) int {
	regions := 0
	for {
		pt := find(hashes)
		if !pt.valid() {
			break
		}

		queue := []point{pt}
		visited := make(map[point]struct{})
		visited[pt] = struct{}{}
		clear(hashes, pt)

		for len(queue) > 0 {
			pt = queue[0]
			queue = queue[1:]

			up := point{pt.x, pt.y - 1}
			down := point{pt.x, pt.y + 1}
			left := point{pt.x - 1, pt.y}
			right := point{pt.x + 1, pt.y}

			adjacent := []point{up, down, left, right}

			for _, pt := range adjacent {
				if pt.valid() && used(hashes, pt) {
					clear(hashes, pt)
					visited[pt] = struct{}{}
					queue = append(queue, pt)
				}
			}
		}

		regions++
	}
	return regions
}

func main() {
	flq := hashem("flqrgnkx")
	ljo := hashem("ljoxqyyw")

	fmt.Println(usage(flq))
	fmt.Println(usage(ljo))

	fmt.Println(regions(flq))
	fmt.Println(regions(ljo))
}
