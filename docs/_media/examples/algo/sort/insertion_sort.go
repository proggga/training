// test it in https://play.golang.org/
package main

import "fmt"

func main() {
	testArray := []int32{1, 9, 2, 4, 8, 7, 3, 5, 6}
	insertionSort(len(testArray), testArray)
	fmt.Printf("%+v\n", testArray)
}

/// [insertionSort]
func insertionSort(n int, arr []int32) {
	for i := 1; i < n; i++ {
		left := i - 1
		right := i
		target := arr[i]

		for right > 0 && arr[left] > target { // начинаем двигать влево, ищем нужное место
			arr[right] = arr[left]
			left--
			right--
		}
		// arr[left] < target - левое значение меньше чем arr[i] значит вставляем
		arr[right] = target
	}
}
/// [insertionSort]