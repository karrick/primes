package main

import (
	"fmt"
	"sort"
	"testing"
)

func ensureIntSlicesMatch(tb testing.TB, actual, expected []int) {
	tb.Helper()

	results := make(map[int]int)

	for _, s := range expected {
		results[s] = 1
	}
	for _, s := range actual {
		results[s]--
	}

	keys := make([]int, 0, len(results))
	for k := range results {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, s := range keys {
		v, ok := results[s]
		if !ok {
			panic(fmt.Errorf("cannot find key: %v", s)) // panic because this function is broken
		}
		switch v {
		case -1:
			tb.Errorf("GOT: %v (extra)", s)
		case 0:
			// both slices have this key
		case 1:
			tb.Errorf("WANT: %v (missing)", s)
		default:
			panic(fmt.Errorf("key has invalid value: %v: %v", s, v)) // panic because this function is broken
		}
	}
}
