package main

import "container/heap"

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
	// this is not dijkstra yet, like at all
	// (a) of all, it loops forever
	unvisited := &boards{}
	heap.Push(unvisited, init)
	for {
		if unvisited.Len() == 0 {
			return nil
		}
		node := heap.Pop(unvisited).(*board)
		if node.win() {
			return node
		}
		for _, neighbor := range node.neighbors() {
			heap.Push(unvisited, neighbor)
		}
	}
}
