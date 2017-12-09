package main

import (
	"fmt"
	"io/ioutil"
)

type group struct {
	Parent   *group
	Children []*group
}

func (g *group) String() string {
	return fmt.Sprint(g.Children)
}

func (g *group) Score(depth int) int {
	sum := depth
	for _, child := range g.Children {
		sum += child.Score(depth + 1)
	}
	return sum
}

func parse(file string) (*group, int) {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	str := string(bytes)
	var root *group
	var cur *group

	count := 0
	garbage := false
	bang := false

	for _, char := range str {
		if bang {
			// Bang escapes whatever comes after it.
			bang = false
		} else if garbage {
			// Garbage escapes everything up until a '>' except a '!'.
			if char == '!' {
				bang = true
			} else if char == '>' {
				garbage = false
			} else {
				count++
			}
		} else {
			switch char {
			case '!':
				bang = true

			case '<':
				garbage = true

			case '{':
				// Start of a new group.
				grp := &group{Parent: cur, Children: []*group{}}
				if cur == nil {
					root = grp
				} else {
					cur.Children = append(cur.Children, grp)
				}
				cur = grp

			case '}':
				// End of the current group.
				cur = cur.Parent

			case ',':
				// Ignored.

			default:
				panic(char)
			}
		}
	}

	return root, count
}

func main() {
	root, count := parse("input.txt")
	fmt.Println(root.Score(1))
	fmt.Println(count)
}
