package main

import "testing"

func TestCheapestWithinKStops(t *testing.T) {
	flights := []Flight{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}
	cases := []struct {
		n, src, dst, k, want int
	}{
		{3, 0, 2, 1, 200}, // via 1
		{3, 0, 2, 0, 500}, // direct only
		{3, 0, 1, 0, 100},
		{3, 2, 0, 1, -1}, // no flights backwards
		{3, 0, 0, 0, 0},  // already there
	}
	for _, c := range cases {
		if got := CheapestWithinKStops(c.n, flights, c.src, c.dst, c.k); got != c.want {
			t.Errorf("k=%d %d->%d = %d, want %d", c.k, c.src, c.dst, got, c.want)
		}
	}

	// The snapshot matters: without it, one round could chain two flights.
	chain := []Flight{{0, 1, 1}, {1, 2, 1}, {2, 3, 1}}
	if got := CheapestWithinKStops(4, chain, 0, 3, 0); got != -1 {
		t.Errorf("k=0 over a 3-flight chain = %d, want -1", got)
	}
	if got := CheapestWithinKStops(4, chain, 0, 3, 1); got != -1 {
		t.Errorf("k=1 over a 3-flight chain = %d, want -1", got)
	}
	if got := CheapestWithinKStops(4, chain, 0, 3, 2); got != 3 {
		t.Errorf("k=2 over a 3-flight chain = %d, want 3", got)
	}
}
