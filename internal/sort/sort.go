package sort

// BubbleSort сортировка пузырьком, чисто по угару O (n^2)
// Проходим массив двумя циклами, свапаем элементы массива между собой
func BubbleSort(a []int) []int {
	length := len(a)
	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}

	return a
}

// InsertionSort сортировка вставкой, O (n^2)
// Проходим весь массив одним циклом, на каждом элементе
// проходим повторно массив и ищем его позицию,
// после чего вставляем
func InsertionSort(a []int) []int {
	length := len(a)
	for i := 0; i < length; i++ {
		item := a[i]
		j := i - 1
		for j >= 0 && a[j] > item {
			a[j+1] = a[j]
			j--
		}
		a[j+1] = item
	}
	return a
}

// SelectionSort сортировка выбором, O (n^2)
// Проходим циклом по всему массив, на каждом этапе
// запоминаем самый маленький элемент и вставляем его
// на необходимую позицию в массиве
func SelectionSort(a []int) []int {
	length := len(a)
	for i := 0; i < length; i++ {
		minIndex := i
		for j := i + 1; j < length; j++ {
			if a[j] < a[minIndex] {
				minIndex = j
			}
		}

		if a[minIndex] != a[i] {
			a[minIndex], a[i] = a[i], a[minIndex]
		}
	}
	return a
}

// QuickSort быстрая сортировка, O (n^2), средний случай O (n log n)
// Рекурсивный алгоритм, который выбирает опорный элемент и разделяет
// массив на две части, элемент меньше опорного и элемент больше опорного
// Реализуется в две функции, одна определяет деление,
// вторая сортирует элементы (quickPartiotion)
func QuickSort(a []int, low int, high int) {
	if low < high {
		pi := quickPartition(a, low, high)

		QuickSort(a, low, pi-1)
		QuickSort(a, pi+1, high)
	}
}

func quickPartition(a []int, low, high int) int {
	pivot := a[high]
	i := low - 1

	for j := low; j < high; j++ {
		if a[j] <= pivot {
			i++
			a[i], a[j] = a[j], a[i]
		}
	}

	a[i+1], a[high] = a[high], a[i+1]
	return i + 1
}

// MergeSort сортировка слияние, O (n log n)
// Рекурсивный алгоритм, в котором мы дробим массив на 2 или 1 элемент
// после чего начинаем рекурсивно складывать и сливать массив в один исходный
func MergeSort(a []int) []int {
	length := len(a)
	if length == 2 {
		if a[0] > a[1] {
			return []int{a[1], a[0]}
		} else {
			return []int{a[0], a[1]}
		}
	}

	if length == 1 {
		return []int{a[0]}
	}

	mid := len(a) / 2
	ans := make([]int, 0, length)
	right := MergeSort(a[mid:])
	left := MergeSort(a[:mid])

	curL, curR := 0, 0

	for curL < len(left) && curR < len(right) {
		if left[curL] > right[curR] {
			ans = append(ans, right[curR])
			curR++
		} else {
			ans = append(ans, left[curL])
			curL++
		}
	}

	ans = append(ans, left[curL:]...)
	ans = append(ans, right[curR:]...)

	return ans
}
