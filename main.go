package main

import "fmt"

// Flight is a directed priced edge.
type Flight struct {
	From, To, Price int
}

// CheapestWithinKStops returns the cheapest fare from src to dst using at
// most k intermediate stops (k+1 flights), or -1. Bellman-Ford capped at
// k+1 rounds, relaxing against a SNAPSHOT of the previous round so each
// round adds at most one flight to any route.
func CheapestWithinKStops(n int, flights []Flight, src, dst, k int) int {
	const inf = int(^uint(0) >> 1)
	dist := make([]int, n)
	for i := range dist {
		dist[i] = inf
	}
	dist[src] = 0
	for round := 0; round <= k; round++ {
		snapshot := append([]int{}, dist...)
		for _, f := range flights {
			if snapshot[f.From] != inf && snapshot[f.From]+f.Price < dist[f.To] {
				dist[f.To] = snapshot[f.From] + f.Price
			}
		}
	}
	if dist[dst] == inf {
		return -1
	}
	return dist[dst]
}

func main() {
	flights := []Flight{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}
	fmt.Println(CheapestWithinKStops(3, flights, 0, 2, 1))
	fmt.Println(CheapestWithinKStops(3, flights, 0, 2, 0))
}
