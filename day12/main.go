package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type node struct {
	id    int
	edges []int
}

func (n *node) String() string {
	return fmt.Sprintf("%v <-> %v\n", n.id, n.edges)
}

func atoi(s string) int {
	i, e := strconv.Atoi(s)
	if e != nil {
		panic(e)
	}
	return i
}

func parse(filename string) map[int]*node {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scan := bufio.NewScanner(file)
	result := make(map[int]*node)

	for scan.Scan() {
		parts := strings.Split(scan.Text(), " ")
		id := atoi(parts[0])
		edges := make([]int, 0)

		for i := 2; i < len(parts); i++ {
			part := parts[i]
			if strings.HasSuffix(part, ",") {
				part = part[:len(part)-1]
			}
			edges = append(edges, atoi(part))
		}

		result[id] = &node{
			id:    id,
			edges: edges,
		}
	}

	if scan.Err() != nil {
		panic(scan.Err())
	}

	return result
}

func main() {
	graph := parse("input.txt")

	queue := make([]int, 0)
	queue = append(queue, 0)

	visited := make(map[int]struct{})
	visited[0] = struct{}{}

	for len(queue) > 0 {
		id := queue[0]
		queue = queue[1:]

		node := graph[id]
		for _, edge := range node.edges {
			if _, ok := visited[edge]; !ok {
				queue = append(queue, edge)
				visited[edge] = struct{}{}
			}
		}
	}

	fmt.Println(len(visited))
}
