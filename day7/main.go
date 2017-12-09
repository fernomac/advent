package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type node struct {
	Name     string
	Weight   int
	Children []string
}

func (n node) String() string {
	return fmt.Sprintf("%v (%v) %v", n.Name, n.Weight, n.Children)
}

func parse(file string) []node {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scan := bufio.NewScanner(f)
	nodes := make([]node, 0)

	for scan.Scan() {
		parts := strings.Split(scan.Text(), " ")

		name := parts[0]

		weightstr := parts[1]
		weight, err := strconv.Atoi(weightstr[1 : len(weightstr)-1])
		if err != nil {
			panic(err)
		}

		children := []string{}

		if len(parts) > 2 {
			children = parts[3:]
			for i := 0; i < len(children)-1; i++ {
				child := children[i]
				children[i] = child[:len(child)-1]
			}
		}

		nodes = append(nodes, node{
			Name:     name,
			Weight:   weight,
			Children: children,
		})
	}

	if scan.Err() != nil {
		panic(scan.Err())
	}

	return nodes
}

func main() {
	nodes := parse("input.txt")
	parents := map[string]string{}

	for _, node := range nodes {
		// fmt.Println(node)
		for _, child := range node.Children {
			parents[child] = node.Name
		}
	}

	for _, node := range nodes {
		if _, ok := parents[node.Name]; !ok {
			fmt.Println(node.Name)
			return
		}
	}
}
