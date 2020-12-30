package main

import "container/heap"

func (b board) hash() string { return "ok" }
func (b board) depth() int   { return len(b.path) }

type boards struct {
	boards []board
	set    map[string]struct{}
}

func (b *boards) Push(x interface{}) {
	xx := x.(board)
	b.set[xx.hash()] = struct{}{}
	b.boards = append(b.boards, xx)
}

func (b *boards) Pop() interface{} {
	xx := b.boards[len(b.boards)-1]
	b.boards = b.boards[:len(b.boards)-1]
	delete(b.set, xx.hash())
	return xx
}

func (b *boards) Len() int           { return len(b.boards) }
func (b *boards) Swap(i, j int)      { b.boards[i], b.boards[j] = b.boards[j], b.boards[i] }
func (b *boards) Less(i, j int) bool { return b.boards[i].depth() < b.boards[j].depth() }

func dijkstra(init board) board {
	unvisited := &boards{set: make(map[string]struct{})}
	heap.Push(unvisited, init)
	for {

	}
	return init
}
