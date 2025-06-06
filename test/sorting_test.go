package test

import (
	"data-structures-algorithms/sorting"
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{5, 4, 3, 2, 1}
	sorting.BubbleSort(arr)
	fmt.Println(arr)
}

func TestInsertionSort(t *testing.T) {
	arr := []int{5, 4, 3, 2, 1}
	sorting.InsertionSort(arr)
	fmt.Println(arr)
}

func TestMergeSort(t *testing.T) {
	arr := []int{5, 4, 3, 2, 1}
	m := sorting.MergeSort(arr)
	fmt.Println(m)
}

func TestQuickSort(t *testing.T) {
	arr := []int{5, 4, 3, 2, 1}
	sorting.QuickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func TestSelectionSort(t *testing.T) {
	arr := []int{5, 4, 3, 2, 1}
	sorting.SelectionSort(arr)
	fmt.Println(arr)
}
