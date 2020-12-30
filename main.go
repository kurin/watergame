package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type tube struct {
	items []byte
}

func (t tube) hash(w io.Writer) {
	w.Write(t.items)
}

func (t tube) room(o tube) bool {
	if len(t.items) == 0 || len(o.items) >= 4 {
		return false
	}
	if len(o.items) == 0 {
		return true
	}
	return o.items[len(o.items)-1] == t.items[len(t.items)-1]
}

func (t tube) topN() int {
	if len(t.items) == 0 {
		return 0
	}
	v := t.items[len(t.items)-1]
	for i := 0; i < len(t.items); i++ {
		idx := len(t.items) - (i + 1)
		if t.items[idx] != v {
			return i
		}
	}
	return len(t.items)
}

func (t tube) pour(o tube) (tube, tube) {
	if !t.room(o) {
		return t, o
	}

	n := t.topN()
	if n > 4-len(o.items) {
		n = 4 - len(o.items)
	}
	i := len(t.items) - n

	newFrom := t.items[:i]
	newTo := make([]byte, len(o.items)+n)
	copy(newTo, o.items)
	copy(newTo[len(o.items):], t.items[i:])

	return tube{items: newFrom}, tube{items: newTo}
}

func (t tube) done() bool {
	if len(t.items) == 0 {
		return true
	}
	var v byte
	for i := range t.items {
		if i == 0 {
			v = t.items[i]
		}
		if v != t.items[i] {
			return false
		}
	}
	return true
}

type pour struct {
	from, to int
}

type board struct {
	path  []pour
	tubes []tube

	visited bool
	dist    int
}

func (b *board) win() bool {
	for _, t := range b.tubes {
		if !t.done() || (len(t.items) > 0 && len(t.items) < 3) {
			return false
		}
	}
	return true
}

func (b *board) pour(i, j int) (*board, bool) {
	if i == j {
		return nil, false
	}
	tl, tr := b.tubes[i], b.tubes[j]
	if !tl.room(tr) {
		return nil, false
	}
	nb := &board{
		tubes: make([]tube, len(b.tubes)),
		path:  make([]pour, len(b.path)),
	}
	copy(nb.tubes, b.tubes)
	copy(nb.path, b.path)
	nb.path = append(nb.path, pour{from: i, to: j})
	nb.tubes[i], nb.tubes[j] = tl.pour(tr)
	return nb, true
}

func (b *board) hash() [16]byte {
	h := md5.New()
	for i, t := range b.tubes {
		h.Write([]byte{byte(i)})
		t.hash(h)
	}
	var out [16]byte
	copy(out[:], h.Sum(nil))
	return out
}

func (b *board) neighbors() []*board {
	var out []*board
	for i := range b.tubes {
		for j := range b.tubes {
			if nb, ok := b.pour(i, j); ok {
				out = append(out, nb)
			}
		}
	}
	return out
}

func (b *board) String() string {
	var tubes []string
	for _, t := range b.tubes {
		var tube []string
		for i := 0; i < 4; i++ {
			v := " "
			if i < len(t.items) {
				v = strconv.Itoa(int(t.items[i]))
			}
			tube = append(tube, v)
		}
		tubes = append(tubes, "["+strings.Join(tube, " ")+"]")
	}
	return strings.Join(tubes, " ")
}

func (b *board) ppath() string {
	var steps []string
	for _, p := range b.path {
		steps = append(steps, fmt.Sprintf("%d -> %d", p.from+1, p.to+1))
	}
	return strings.Join(steps, ", ")
}

func main() {
	b := &board{
		tubes: []tube{
			tube{items: []byte{1, 1, 2, 2}},
			tube{items: []byte{2, 2, 1, 1}},
			tube{items: []byte{}},
		},
	}
	fmt.Println(b)
	out := dijkstra(b)
	if out == nil {
		fmt.Println("failure")
		return
	}
	fmt.Println(out.ppath(), out)
}
