**[назад](sort)**

# Пузырковая сортировка [Bubble sort]

[filename](short/bubble_sort.md ':include')

#### сложность алгоритма:
Худший случай	O(n^2) сравнений, O(n^2) перемен местами
Лучший случай	O(n) сравнений, O(1) перемен местами
Средний случай	O(n^2) сравнений, O(n^2) перемен местами
Худший случай по памяти	O(1)

#### принцип работы:
двигаемся слева на право и сравниваем пары,
если один элемент в паре больше другого меняем их местами
далем так N-1 раз, т.к. в массиве N-1 пар
по завершение N-1 цикла массив отсортирован

![Gif algo example from wiki](https://upload.wikimedia.org/wikipedia/commons/c/c8/Bubble-sort-example-300px.gif)

#### пример реализации:

[filename](../_media/examples/sort/bubble_sort.go ':include :type=code :fragment=bubbleSort')

[ссылка на весь файл](https://github.com/proggga/training/blob/master/docs/_media/examples/sort/bubble_sort.go)