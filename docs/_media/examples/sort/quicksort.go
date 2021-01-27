// test it in https://play.golang.org/
package main

import (
	"fmt"
	"time"
)

var (
	swaps_counter int
)

func makeTimestamp() int64 { // play.golang.org has constant time.Now(), check timing in terminal
	return time.Now().UnixNano()
}

func main() {
	fmt.Printf("simple test\n")

	simpleTest := []int32{1, 4, 2, 9, 8, 7, 3, 5, 6, 10}

	testSort(lomutoQuickSort, "lomuto", simpleTest)

	testSort(hoareQuickSort, "hoare", simpleTest)

	testSort(iterativeLomutoQuickSort, "iter_lomuto", simpleTest)

	testSort(quicksort3way, "3-way sort", simpleTest)

	fmt.Printf("hard test\n")

	hardTest := []int32{1, 4, 2, 4, 9, 8, 4, 7, 4, 4, 3, 5, 6, 10}

	testSort(lomutoQuickSort, "lomuto", hardTest)

	testSort(hoareQuickSort, "hoare", hardTest)

	testSort(iterativeLomutoQuickSort, "iter_lomuto", hardTest)

	testSort(quicksort3way, "3-way sort", hardTest)

	fmt.Printf("sorted test\n")

	sortedTest := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	testSort(lomutoQuickSort, "lomuto", sortedTest)

	testSort(hoareQuickSort, "hoare", sortedTest)

	testSort(iterativeLomutoQuickSort, "iter_lomuto", sortedTest)

	testSort(quicksort3way, "3-way sort", sortedTest)

}

func testSort(callback func(arr []int32, from, to int), methodName string, testArrayOriginal []int32) {
	swaps_counter = 0
	testArray := make([]int32, len(testArrayOriginal))
	copy(testArray, testArrayOriginal)
	start := makeTimestamp()
	callback(testArray, 0, len(testArray)-1)
	t := makeTimestamp()
	elapsed := t - start
	fmt.Printf("%20s %+v - %d swaps time %d %d %d\n", methodName, testArray, swaps_counter, elapsed, t, start)
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
	i := low - 1
	j := high + 1
	for {

		// двигаем левый курсор вправо пока не найдем значение,
		// которое нужно свапнуть в правую часть
		i++
		for ; arr[i] < pivot; i++ {
		}

		// двигаем правый курсор влево пока не найдем значение,
		// которое нужно свапнуть в левую часть
		j--
		for ; arr[j] > pivot; j-- {
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

/// [iterativeLomutoQuickSort]

func iterativeLomutoQuickSort(arr []int32, low, high int) {
	// Создаем временный стек
	stack := make([]int, high-low+1)

	// Создаем указатель на верхушку стека
	top := -1

	// Пушим начальные значения low и high в стек
	top++
	stack[top] = low
	top++
	stack[top] = high

	// Продолжаем вытаскивать элементы из стека, пока он не будет пустым
	for top >= 0 {
		// Вытаскиваем low и high
		high = stack[top]
		top--
		low = stack[top]
		top--

		// Выставляем pivot в нужное место
		// и разделяем массив на две подчасти
		p := iterativeLomutoPartition(arr, low, high)

		// Если в левой части остались элементы
		// то пушим левую часть в стэк
		if p-1 > low {
			top++
			stack[top] = low
			top++
			stack[top] = p - 1
		}

		// Если в правой части остались элементы
		// то пушим правую часть в стэк
		if p+1 < high {
			top++
			stack[top] = p + 1
			top++
			stack[top] = high
		}
	}
}

// такой же как и lomutoPartition, магия в iterativeLomutoQuickSort
func iterativeLomutoPartition(arr []int32, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j <= high-1; j++ {
		if arr[j] <= pivot {
			i++
			if i != j {
				swap(arr, i, j)
			}
		}
	}
	swap(arr, i+1, high)
	return i + 1
}

/// [iterativeLomutoQuickSort]

// 3-way partition based quick sort
func quicksort3way(a []int32, l, r int) {
	if r <= l {
		return
	}

	i, j := partition3way(a, l, r)

	// Recursive
	quicksort3way(a, l, j)
	quicksort3way(a, i, r)
}

func partition3way(a []int32, l, r int) (i, j int) {
	i = l - 1
	j = r
	p := l - 1
	q := r
	v := a[r]

	for {
		// From left, find the first element greater than
		// or equal to v. This loop will definitely
		// terminate as v is last element
		i++
		for ; a[i] < v; i++ {
		}

		// From right, find the first element smaller than
		// or equal to v
		j--
		for ; v < a[j]; j-- {
			if j == l {
				break
			}
		}

		// If i and j cross, then we are done
		if i >= j {
			break
		}

		// Swap, so that smaller goes on left greater goes
		// on right
		swap(a, i, j)

		// Move all same left occurrence of pivot to
		// beginning of array and keep count using p
		if a[i] == v {
			p++
			swap(a, p, i)
		}

		// Move all same right occurrence of pivot to end of
		// array and keep count using q
		if a[j] == v {
			q--
			swap(a, j, q)
		}
	}

	// Move pivot element to its correct index
	swap(a, i, r)

	// Move all left same occurrences from beginning
	// to adjacent to arr[i]
	j = i - 1
	for k := l; k < p; k, j = k+1, j-1 {
		swap(a, k, j)
	}

	// Move all right same occurrences from end
	// to adjacent to arr[i]
	i = i + 1
	for k := r - 1; k > q; k, i = k-1, i+1 {
		swap(a, i, k)
	}
	return i, j
}

func swap(arr []int32, from, to int) {
	tmp := arr[to]
	arr[to] = arr[from]
	arr[from] = tmp
	swaps_counter++
}
