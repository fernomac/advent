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

func parse(file string) map[string]*node {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scan := bufio.NewScanner(f)
	nodes := make(map[string]*node)

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

		nodes[name] = &node{
			Name:     name,
			Weight:   weight,
			Children: children,
		}
	}

	if scan.Err() != nil {
		panic(scan.Err())
	}

	return nodes
}

func weight(nodes map[string]*node, root string, depth int) int {
	for i := 0; i < depth; i++ {
		fmt.Print("  ")
	}
	fmt.Println("->", root)

	node := nodes[root]

	if len(node.Children) < 2 {
		// They had better be balanced...
		sum := node.Weight
		for _, child := range node.Children {
			sum += weight(nodes, child, depth+1)
		}
		fmt.Printf("%v: (%v)\n", root, sum)
		return sum
	}

	sum := node.Weight
	weights := make([]int, len(node.Children))

	for idx, child := range node.Children {
		weight := weight(nodes, child, depth+1)
		weights[idx] = weight
		sum += weight
	}

	fmt.Printf("%v: (%v)\n", root, sum)

	if weights[0] == weights[1] {
		// Look for an unbalanced child further on.
		for i := 2; i < len(weights); i++ {
			if weights[i] != weights[0] {
				// Child i is wrong. Adjust it.
				offset := weights[0] - weights[i]
				fmt.Println(node.Children[i], "should weigh", nodes[node.Children[i]].Weight+offset)
				os.Exit(0)
			}
		}
	} else if weights[0] == weights[2] {
		// Child 1 is wrong.
		offset := weights[0] - weights[1]
		fmt.Println(node.Children[1], "should weigh", nodes[node.Children[1]].Weight+offset)
		os.Exit(0)
	} else {
		// Child 0 is wrong.
		offset := weights[1] - weights[0]
		fmt.Println(node.Children[0], "should weigh", nodes[node.Children[0]].Weight+offset)
		os.Exit(0)
	}

	return sum
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

	root := ""
	for _, node := range nodes {
		if _, ok := parents[node.Name]; !ok {
			root = node.Name
			break
		}
	}

	fmt.Println(root)

	weight(nodes, root, 0)
}
