package main

import (
	"reflect"
	"testing"
)

func TestTopN(t *testing.T) {
	table := []struct {
		tube tube
		n    int
	}{
		{
			tube: tube{},
			n:    0,
		},
		{
			tube: tube{items: []byte{1, 2, 3, 3}},
			n:    2,
		},
		{
			tube: tube{items: []byte{1, 3, 3, 3}},
			n:    3,
		},
		{
			tube: tube{items: []byte{3, 3, 3, 3}},
			n:    4,
		},
		{
			tube: tube{items: []byte{3, 3, 3, 4}},
			n:    1,
		},
	}

	for _, ent := range table {
		if ent.tube.topN() != ent.n {
			t.Errorf("wrong tube count for %v: got %d, want %d", ent.tube, ent.tube.topN(), ent.n)
		}
	}
}

func TestPour(t *testing.T) {
	table := []struct {
		from, to tube
		wf, wt   tube
	}{
		{
			from: tube{items: []byte{1, 1, 2, 3}},
			to:   tube{items: []byte{2, 2, 3}},
			wf:   tube{items: []byte{1, 1, 2}},
			wt:   tube{items: []byte{2, 2, 3, 3}},
		},
		{
			from: tube{items: []byte{1, 1, 0, 0}},
			to:   tube{items: []byte{}},
			wf:   tube{items: []byte{1, 1}},
			wt:   tube{items: []byte{0, 0}},
		},
		{
			from: tube{items: []byte{1, 1, 1}},
			to:   tube{items: []byte{0, 0, 1}},
			wf:   tube{items: []byte{1, 1}},
			wt:   tube{items: []byte{0, 0, 1, 1}},
		},
	}

	for _, ent := range table {
		gf, gt := ent.from.pour(ent.to)
		if !reflect.DeepEqual(ent.wf, gf) || !reflect.DeepEqual(ent.wt, gt) {
			t.Errorf("pour failed: %v -> %v, got (%v, %v) want (%v, %v)", ent.from, ent.to, gf, gt, ent.wf, ent.wt)
		}
	}
}
