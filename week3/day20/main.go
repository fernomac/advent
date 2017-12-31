package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func atoi(a string) int {
	i, e := strconv.Atoi(a)
	if e != nil {
		panic(e)
	}
	return i
}

type vector struct {
	x, y, z int
}

func vectorize(coords []string) vector {
	return vector{
		x: atoi(coords[0]),
		y: atoi(coords[1]),
		z: atoi(coords[2]),
	}
}

func (v vector) String() string {
	return fmt.Sprintf("(%v, %v, %v)", v.x, v.y, v.z)
}

type particle struct {
	pos, vel, acc vector
}

func (p *particle) Step() {
	p.vel.x += p.acc.x
	p.vel.y += p.acc.y
	p.vel.z += p.acc.z

	p.pos.x += p.vel.x
	p.pos.y += p.vel.y
	p.pos.z += p.vel.z
}

func (p *particle) String() string {
	return fmt.Sprintf("{pos: %v, vel: %v, acc: %v}", p.pos, p.vel, p.acc)
}

func parse(filename string) []*particle {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(bytes), "\n")
	repl := strings.NewReplacer("p=<", "", " v=<", "", " a=<", "", ">", "")
	particles := []*particle{}

	for _, line := range lines {
		parts := strings.Split(repl.Replace(line), ",")
		if len(parts) != 9 {
			panic("strange line: " + line)
		}

		pos := vectorize(parts[0:3])
		vel := vectorize(parts[3:6])
		acc := vectorize(parts[6:9])

		particles = append(particles, &particle{pos, vel, acc})
	}

	return particles
}

func bucketize(particles []*particle) map[vector][]*particle {
	result := map[vector][]*particle{}

	for i := 0; i < len(particles); i++ {
		pos := particles[i].pos
		result[pos] = append(result[pos], particles[i])
	}

	return result
}

func main() {
	particles := parse("input.txt")

	for {

		buckets := bucketize(particles)
		particles = nil

		for pos, list := range buckets {
			if len(list) == 1 {
				particles = append(particles, list[0])
			} else {
				fmt.Println("collision at ", pos, "removes", len(list), "particles")
			}
		}

		fmt.Println(len(particles), "particles remain")

		for _, p := range particles {
			p.Step()
		}
	}
}
