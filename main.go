package main

import (
	"fmt"
	"go-algorithms/internal/sort"
)

func main() {
	a := []int{7, 2, 3, 5, 4, 1, 7, 2, 3, 5, 4, 1, 7, 2, 3, 5, 4, 1, 7, 2, 3, 5, 4, 1}
	b := []int{7, 2, 3, 5, 4, 1, 7, 2, 3, 5, 4, 1, 7, 2, 3, 5, 4, 1, 7, 2, 3, 5, 4, 1}
	c := []int{7, 2, 3, 5, 4, 1, 7, 2, 3, 5, 4, 1, 7, 2, 3, 5, 4, 1, 7, 2, 3, 5, 4, 1}
	d := []int{7, 2, 3, 5, 4, 1, 7, 2, 3, 5, 4, 1, 7, 2, 3, 5, 4, 1, 7, 2, 3, 5, 4, 1}
	f := []int{7, 2, 3, 5, 4, 1, 7, 2, 3, 5, 4, 1, 7, 2, 3, 5, 4, 1, 7, 2, 3, 5, 4, 1}
	//test := []int{4, 3, 2, 1}
	sort.QuickSort(d, 0, len(d)-1)
	fmt.Printf("Bubble Sort : \t %d \n", sort.BubbleSort(a))
	fmt.Printf("Insertion Sort : \t %d \n", sort.InsertionSort(b))
	fmt.Printf("Selection Sort : \t %d \n", sort.SelectionSort(c))
	fmt.Printf("Quick Sort : \t %d \n", d)
	fmt.Printf("Merge Sort : \t %d \n", sort.MergeSort(f))
}
