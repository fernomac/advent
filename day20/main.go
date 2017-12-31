package main

import (
	"fmt"
	"io/ioutil"
	"math"
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

func abs(i int) int {
	if i >= 0 {
		return i
	}
	return -i
}

func (v vector) Magnitude() int {
	return abs(v.x) + abs(v.y) + abs(v.z)
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

func isStable(x, dx int) bool {
	if dx > 0 {
		return x > 0
	}
	if dx < 0 {
		return x < 0
	}
	return true
}

func (p *particle) IsVelStable() bool {
	return isStable(p.vel.x, p.acc.x) && isStable(p.vel.y, p.acc.y) && isStable(p.vel.z, p.acc.z)
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

func findMinAcc(particles []*particle) int {
	minAcc := particles[0].acc.Magnitude()

	for i := 1; i < len(particles); i++ {
		acc := particles[i].acc.Magnitude()
		if acc < minAcc {
			minAcc = acc
		}
	}

	return minAcc
}

func findParticlesWithAcc(particles []*particle, acc int) []int {
	result := []int{}

	for i, p := range particles {
		if p.acc.Magnitude() == acc {
			result = append(result, i)
		}
	}

	return result
}

func filter(particles []*particle, idxs []int) map[int]*particle {
	result := map[int]*particle{}

	for _, idx := range idxs {
		result[idx] = particles[idx]
	}

	return result
}

func findMinVel(candidates map[int]*particle) int {
	min := math.MaxInt32
	for _, p := range candidates {
		vel := p.vel.Magnitude()
		if vel < min {
			min = vel
		}
	}
	return min
}

func findParticlesWithVel(candidates map[int]*particle, vel int) []int {
	result := []int{}
	for idx, p := range candidates {
		if p.vel.Magnitude() == vel {
			result = append(result, idx)
		}
	}
	return result
}

func main() {
	particles := parse("input.txt")

	minAcc := findMinAcc(particles)
	candidateIdxs := findParticlesWithAcc(particles, minAcc)
	candidates := filter(particles, candidateIdxs)

	for {
		stable := true
		for _, c := range candidates {
			c.Step()
			stable = stable && c.IsVelStable()
		}
		if stable {
			break
		}
	}

	minVel := findMinVel(candidates)
	candidateIdxs = findParticlesWithVel(candidates, minVel)

	for _, idx := range candidateIdxs {
		fmt.Printf("%v: %v\n", idx, candidates[idx])
	}
}
