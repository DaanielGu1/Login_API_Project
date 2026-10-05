package models

type node struct {
	data int
	next *node
}

func newNode(data int) *node {
	return &node{data: data}
}
