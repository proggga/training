// test it in https://play.golang.org/
package main

import "fmt"

func main() {
	testArray := []int32{1, 9, 2, 4, 8, 7, 3, 5, 6}
	bubbleSort(testArray)
	fmt.Printf("%+v\n", testArray)
}

/// [bubbleSort]
// сравниваем два соседних элемента и если один больше другого меняем местами
// повторяем N-1 раз
func bubbleSort(arr []int32) {
	length := len(arr)
	// цикл повторяем N-1 раз чтоб убедится что все числа встали на места
	for i := 0; i < length-1; i++ {
		for j := i; j < length-1; j++ {
			// сравниваем текущее и следующие, если больше то меняем
			if arr[j] > arr[j+1] {
				swap(arr, j, j+1)
			}
		}
	}
}
/// [bubbleSort]

func swap(arr []int32, from, to int) {
	to_value := arr[to]
	arr[to] = arr[from]
	arr[from] = to_value
}
