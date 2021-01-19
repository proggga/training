// test it in https://play.golang.org/
package main

import "fmt"

var (
	swaps_counter int
)

func main() {
	swaps_counter = 0
	testArray := []int32{1, 4, 2, 9, 8, 7, 3, 5, 6}
	lomutoQuickSort(testArray, 0, len(testArray)-1)
	fmt.Printf("lomuto\t%+v - %d swaps\n", testArray, swaps_counter)

	swaps_counter = 0
	testArray = []int32{1, 4, 2, 9, 8, 7, 3, 5, 6}
	hoareQuickSort(testArray, 0, len(testArray)-1)
	fmt.Printf("hoare\t%+v - %d swaps\n", testArray, swaps_counter)
}

/// [lomutoQuickSort]
func lomutoQuickSort(arr []int32, from, to int) {
	if from < to {
		el := lomutoPartition(arr, from, to)
		lomutoQuickSort(arr, from, el-1)
		lomutoQuickSort(arr, el+1, to)
	}
}

func lomutoPartition(arr []int32, low, high int) int {
	pivot := arr[high]
	i := low
	for j := low; j <= high-1; j++ {
		if arr[j] <= pivot {
			if i != j {
				swap(arr, i, j)
			}
			i++
		}
	}
	swap(arr, i, high)
	return i
}

/// [lomutoQuickSort]

/// [hoareQuickSort]
func hoareQuickSort(arr []int32, from, to int) {
	if from < to {
		el := hoarePartition(arr, from, to)
		hoareQuickSort(arr, from, el) // тут разница с lomuto, тут нет -1
		hoareQuickSort(arr, el+1, to)
	}
}

func hoarePartition(arr []int32, low, high int) int {
	pivot := (arr[low] + arr[high]) / 2
	i := low
	j := high
	for {

    // двигаем левый курсор вправо пока не найдем значение,
    // которое нужно свапнуть в правую часть
		for arr[i] < pivot {
			i++
		}

    // двигаем правый курсор влево пока не найдем значение,
    // которое нужно свапнуть в левую часть
    for arr[j] > pivot {
			j--
		}

		// курсоры сошлись, выходим - больше делать нечего
		if i >= j {
			return j
		}

		// если не сошлись, меняем местами и продолжаем работать
		swap(arr, i, j)
	}
}

/// [hoareQuickSort]

func swap(arr []int32, from, to int) {
	tmp := arr[to]
	arr[to] = arr[from]
	arr[from] = tmp
	swaps_counter++
}
