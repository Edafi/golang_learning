package main

import "fmt"

type List struct {
	start *Node
	end   *Node
}

type Node struct {
	value     int
	next_node *Node
	prev_node *Node
}

func (list *List) create(value int) {
	var new_node Node = Node{value: value, next_node: nil, prev_node: nil}
	list.start = &new_node
	list.end = &new_node
}

func (list *List) add(value int) {
	var new_node Node = Node{value: value, next_node: nil, prev_node: list.end}
	list.end.next_node = &new_node
	list.end = &new_node
}

func (list *List) delete() {
	var prev_node *Node = list.end.prev_node
	list.end.prev_node = nil
	prev_node.next_node = nil
	list.end = prev_node
}

func (list *List) print() {
	var current_node *Node = list.start
	for current_node != nil {
		fmt.Println(current_node.value)
		current_node = current_node.next_node
	}
}
