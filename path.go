package main

import (
	"container/heap"
	"fmt"
	"math/rand"
)

func (b *board) depth() int { return len(b.path) }

type boards struct {
	boards []*board
}

func (b *boards) Push(x interface{}) {
	xx := x.(*board)
	b.boards = append(b.boards, xx)
}

func (b *boards) Pop() interface{} {
	xx := b.boards[len(b.boards)-1]
	b.boards[len(b.boards)-1] = nil
	b.boards = b.boards[:len(b.boards)-1]
	return xx
}

func (b *boards) Len() int           { return len(b.boards) }
func (b *boards) Swap(i, j int)      { b.boards[i], b.boards[j] = b.boards[j], b.boards[i] }
func (b *boards) Less(i, j int) bool { return b.boards[i].depth() < b.boards[j].depth() }

func dijkstra(init *board) *board {
	unvisited := &boards{}
	heap.Push(unvisited, init)
	seen := map[[12]byte]struct{}{}
	var i int
	for {
		i++
		if i%10000 == 0 {
			fmt.Println(unvisited.Len(), len(seen))
		}
		if unvisited.Len() == 0 {
			return nil
		}
		node := heap.Pop(unvisited).(*board)
		if node.win() {
			return node
		}
		seen[node.hash()] = struct{}{}
		ns := node.neighbors()
		rand.Shuffle(len(ns), func(i, j int) { ns[i], ns[j] = ns[j], ns[i] })
		for _, neighbor := range ns {
			if _, ok := seen[neighbor.hash()]; !ok && unvisited.Len() < 60000 {
				heap.Push(unvisited, neighbor)
			}
		}
	}
}
