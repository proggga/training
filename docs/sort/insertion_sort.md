**[назад](sort)**

# Сортировка вставками [Insertion sort]

[filename](short/insertion_sort.md ':include')

#### сложность алгоритма:
- Худший случай O(n^2) сравнений и перемен местами
- Лучший случай O(n) сравнений, O(1) перемен местами
- Средний случай O(n^2) сравнений и перемен местами
- Худший случай по памяти О(n), и O(1) вспомогательный

#### принцип работы:

итерируемся слева на право

на каждой итерации

запомнимаем число и передвигая влево, ищем нужное место

![Gif](../_media/img/sort/insertion_sort.gif)

#### пример реализации:

[filename](../_media/examples/sort/insertion_sort.go ':include :type=code :fragment=insertionSort')

[ссылка на весь файл](https://github.com/proggga/training/blob/master/docs/_media/examples/sort/insertion_sort.go)
