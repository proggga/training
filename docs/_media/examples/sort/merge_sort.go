// test it in https://play.golang.org/
package main

import "fmt"

func main() {
	testArray := []int32{1, 9, 2, 4, 8, 7, 3, 5, 6}
	sort(testArray)
	fmt.Printf("%+v\n", testArray)
}

func sort(arr []int32) {
	length := len(arr)

	mergeSort(arr, 0, length-1)
}

/// [mergeSortSplit]

func mergeSort(arr []int32, from, to int) {
	length := to - from + 1

	// пустой массив или массив из 1 элемента уже отсортирован,
	// также тут выйдет если курсоры выйдут за границы
	if length <= 1 {
		return
	}

	// если массив простой и состоит из 2 элементов, то сортируем его
	if length == 2 {
		if arr[from] > arr[to] {
			swap(arr, from, to)
		}
	}

	// дробим на наименьшие массивы
	part_size := length / 2

	leftB := from + part_size
	mergeSort(arr, from, leftB-1)
	mergeSort(arr, leftB, to)

	// после сортировки массивов склеиваем/мержим два массива
	merge(arr, length, from, leftB, to)
}

/// [mergeSortSplit]

/// [mergeSortMerge]

func merge(arr []int32, length, from, leftB, to int) {
	// я тут склеиваю в отдельный массив, который переложу в основной позже
	res := make([]int32, length)
	inserted := 0
	x := from
	y := leftB

	// берем наименьшие, и работаем пока не вышли за границы одного из массивов
	for x <= leftB-1 && y <= to {
		if arr[x] < arr[y] {
			res[inserted] = arr[x]
			x++
		} else {
			res[inserted] = arr[y]
			y++
		}
		inserted++
	}
	// если один массив уже закончился, а второй все еще нет,
	// то просто перекладываем хвост второго
	for x <= leftB-1 {
		res[inserted] = arr[x]
		x++
		inserted++
	}
	// докладываем так же если осталось
	for y <= to {
		res[inserted] = arr[y]
		y++
		inserted++
	}
	q := from
	for _, v := range res { // склеенное перекладываем в основной массив
		arr[q] = int32(v)
		q++
	}
}

/// [mergeSortMerge]

func swap(arr []int32, from, to int) {
	to_value := arr[to]
	arr[to] = arr[from]
	arr[from] = to_value
}
