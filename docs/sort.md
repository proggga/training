# Сортировка

Сортировка самый первый вид алгоритмов, который встречаешь при подготовке к интервью, и тут начинается

Вот тут [полезная таблица (en)](https://en.wikipedia.org/wiki/Sorting_algorithm#Comparison_of_algorithms) сравнения сложностей алгоритмов

Еще полезный [cписок сортировок](https://ru.wikipedia.org/wiki/%D0%90%D0%BB%D0%B3%D0%BE%D1%80%D0%B8%D1%82%D0%BC_%D1%81%D0%BE%D1%80%D1%82%D0%B8%D1%80%D0%BE%D0%B2%D0%BA%D0%B8#%D0%A1%D0%BF%D0%B8%D1%81%D0%BE%D0%BA_%D0%B0%D0%BB%D0%B3%D0%BE%D1%80%D0%B8%D1%82%D0%BC%D0%BE%D0%B2_%D1%81%D0%BE%D1%80%D1%82%D0%B8%D1%80%D0%BE%D0%B2%D0%BA%D0%B8)

### [Характеристики алгоритмов](https://en.wikipedia.org/wiki/Sorting_algorithm#Classification):
- [computational complexity](https://en.wikipedia.org/wiki/Computational_complexity)
- [online](https://en.wikipedia.org/wiki/Online_algorithm) (онлайн) - может сортировать данные частично, по мере поступления
- [adaptive](https://en.wikipedia.org/wiki/Adaptive_sort) (адаптивный) - чем больше отсортирован, тем меньше работает
- [in-place](https://en.wikipedia.org/wiki/In-place_algorithm) (на месте) - умеет ли алгоритм сортировать массив без доп памяти
- [stable](https://en.wikipedia.org/wiki/Sorting_algorithm#Stability) (стабильный) - при сортировке оставляет одинаковые ключи в массив в том же порядке как и были в исходном

### [Пузырьковая сортировка (bubble sort)](sort/bubble_sort.md)

[filename](sort/short/bubble_sort.md ':include')

### [Cортировка вставками (insertion sort)](sort/insertion_sort.md)

[filename](sort/short/insertion_sort.md ':include')

### [Быстрая сортировка (quicksort)](sort/quicksort.md)

[filename](sort/short/quicksort.md ':include')

### [Сортировка слияниями (merge sort)](sort/merge_sort.md)

[filename](sort/short/merge_sort.md ':include')