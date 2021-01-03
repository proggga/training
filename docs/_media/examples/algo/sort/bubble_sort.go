// test it in https://play.golang.org/
package main
import "fmt"

func main() {
	testArray := []int32{1, 9, 2, 4, 8, 7, 3, 5, 6}
	fmt.Printf("%+v\n", bubbleSort(testArray))
}

### [bubble_sort]
// сравниваем два соседних элемента и если один больше другого меняем местами
// повторяем N-1 раз
func bubbleSort(arr []int32) []int32 {
	length := len(arr)
	for i := 0; i < length - 1; i++ { // цикл повторяем N-1 раз чтоб убедится что все числа встали на места
		for j := i; j < length - 1; j++ {
			if arr[j] > arr[j+1] { // сравниваем текущее и следующие, если больше то меняем
				swap(arr, j, j+1)
			}
		}
	}
	return arr
}
### [bubble_sort]

func swap(arr []int32, from, to int) {
	to_value := arr[to]
	arr[to] = arr[from]
	arr[from] = to_value
}