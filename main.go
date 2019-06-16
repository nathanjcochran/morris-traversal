package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var t Tree
	for _, arg := range os.Args[1:] {
		i, err := strconv.Atoi(arg)
		if err != nil {
			log.Fatalf("Invalid int: %s", arg)
		}
		t.Add(i)
	}

	results := t.Traverse()
	fmt.Println(results)
}

type Tree struct {
	root *node
}

type node struct {
	left  *node
	right *node
	val   int
}

func (t *Tree) Add(val int) {
	if t.root == nil {
		t.root = &node{val: val}
		return
	}

	current := t.root
	for {
		if val < current.val {
			if current.left == nil {
				current.left = &node{val: val}
				return
			} else {
				current = current.left
			}
		} else {
			if current.right == nil {
				current.right = &node{val: val}
				return
			} else {
				current = current.right
			}
		}
	}
}

func (t *Tree) Traverse() []int {
	var (
		results []int
		current = t.root
	)
	for current != nil {
		if current.left == nil {
			results = append(results, current.val)
			current = current.right
		} else {
			pre := current.left
			for pre.right != nil && pre.right != current {
				pre = pre.right
			}

			if pre.right == nil {
				pre.right = current
				current = current.left
			} else {
				pre.right = nil
				results = append(results, current.val)
				current = current.right
			}
		}
	}

	return results
}
