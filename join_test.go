package splaytree

import "testing"

func TestJoin(t *testing.T) {
	fir := NewSplayTree(func(l, r int) bool {
		return l < r
	})
	oak := NewSplayTree(func(l, r int) bool {
		return l < r
	})
	it1 := []int{6, 4, 2, 5, 3, 7, 0}
	it2 := []int{-1, 14, -2, 15, -3, 17, 10}
	fir.InsertAll(it1)
	fir.Join(nil)
	fir.Join(fir)
	oak.InsertAll(it1)
	fir.Join(oak)
	if fir.Count() != len(it1) {
		t.Errorf("join count !it1")
	}
	if oak.NonEmpty() {
		t.Errorf("join oak !0")
	}
	fir.Check()
	oak.InsertAll(it2)
	fir.Join(oak)
	if fir.Count() != len(it1)+len(it2) {
		t.Errorf("join count !it1+it2")
	}
	if oak.NonEmpty() {
		t.Errorf("join oak !0")
	}
	fir.Check()
	for _, item := range append(it1, it2...) {
		if _, ok := fir.Lookup(item); !ok {
			t.Errorf("join lookup !%v", item)
		}
	}
	fir.Check()
}
