package main

import (
	"container/heap"
)

func (b *board) depth() int { return len(b.path) }

type boards struct {
	boards []*board
}

//func (b *boards) seen

func (b *boards) Push(x interface{}) {
	xx := x.(*board)
	b.boards = append(b.boards, xx)
}

func (b *boards) Pop() interface{} {
	xx := b.boards[len(b.boards)-1]
	b.boards = b.boards[:len(b.boards)-1]
	return xx
}

func (b *boards) Len() int           { return len(b.boards) }
func (b *boards) Swap(i, j int)      { b.boards[i], b.boards[j] = b.boards[j], b.boards[i] }
func (b *boards) Less(i, j int) bool { return b.boards[i].depth() < b.boards[j].depth() }

func dijkstra(init *board) *board {
	unvisited := &boards{}
	heap.Push(unvisited, init)
	seen := map[[16]byte]struct{}{}
	for {
		if unvisited.Len() == 0 {
			return nil
		}
		node := heap.Pop(unvisited).(*board)
		if node.win() {
			return node
		}
		seen[node.hash()] = struct{}{}
		for _, neighbor := range node.neighbors() {
			if _, ok := seen[neighbor.hash()]; !ok {
				heap.Push(unvisited, neighbor)
			}
		}
	}
}
