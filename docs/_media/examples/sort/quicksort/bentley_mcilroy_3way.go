package main

import (
	"fmt"
)

func main() {
	arr := []int32{1, 4, 2, 3, 4, 2, 3, 5, 6, 6, 8, 754, 23, 1}
	quicksort3way(arr)
	fmt.Printf("%+v", arr)
}

// Bentley-McIlroy 3-way partitioning
// https://www.cs.princeton.edu/~rs/talks/QuicksortIsOptimal.pdf
func quicksort3way(arr []int32) {
	quicksortRecursive(arr, 0, len(arr)-1)
}

func quicksortRecursive(arr []int32, from, to int) {
	if from >= to {
		return
	}

	j, i := partition(arr, from, to)
	quicksortRecursive(arr, from, j)
	quicksortRecursive(arr, i, to)

}

func partition(arr []int32, from, to int) (j, i int) {
	i = from - 1
	j = to

	p := from - 1
	q := to

	v := arr[to]

	for {
		i++
		for ; arr[i] < v; i++ {
		}
		j--
		for ; v < arr[j]; j-- {
			if j == from {
				break
			}
		}
		if i >= j {
			break
		}

		swap(arr, i, j)
		if arr[i] == v {
			p++
			swap(arr, p, i)
		}
		if v == arr[j] {
			q--
			swap(arr, j, q)
		}
	}
	swap(arr, i, to)
	j = i - 1
	i = i + 1

	for k := from; k < p; k, j = k+1, j-1 {
		swap(arr, k, j)
	}

	for k := to - 1; k > q; k, i = k-1, i+1 {
		swap(arr, i, k)
	}

	return j, i
}

func swap(arr []int32, f, t int) {
	tmp := arr[f]
	arr[f] = arr[t]
	arr[t] = tmp
}
