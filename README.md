# Cheapest Flights Within K Stops

> NeetCode 150 · Advanced Graphs — Bellman-Ford's rounds ARE a hop budget: cap at k+1 rounds, relaxing against last round's snapshot so no round chains two flights. Go, stdlib only, tested.

## Problem

Cheapest fare from src to dst using at most k intermediate stops (so k+1 flights), or −1. The hop constraint breaks Dijkstra — a pricier-but-shorter route can beat a cheaper-but-longer one it would have discarded.

## Approach

Bellman-Ford (dsa-166) has the constraint built into its anatomy: after r rounds, distances are optimal *among routes of at most r edges*. So run exactly k+1 rounds. The one modification that matters: relax against a **snapshot** of the previous round, not the live array — otherwise a single round can ripple through a whole chain (flight A updates B, B's fresh value immediately updates C), silently exceeding the hop budget. The 3-flight-chain tests pin exactly that failure mode at k=0,1,2. No negative-cycle pass is needed; the round cap bounds everything.

**Complexity:** O(k · E) time, O(n) space.

## Run

    go test ./...   # verify
    go run .        # demo

---

Part of a 275-repo DSA series in Go — Striver's SDE Sheet (dsa-001–191) plus the NeetCode 150 / Blind 75 additions (dsa-192–275), stdlib only, every solution tested. Browse: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
