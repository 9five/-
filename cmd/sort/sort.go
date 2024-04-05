package main

import (
	_bubbleSortUsecase "algorithm/bubbleSort/usecase"
	_heapSortUsecase "algorithm/heapSort/usecase"
	_insertionSortUsecase "algorithm/insertionSort/usecase"
	_mergeSortUsecase "algorithm/mergeSort/usecase"
	_quickSortUsecase "algorithm/quickSort/usecase"
	_selectionSortUsecase "algorithm/selectionSort/usecase"
	"fmt"
)

func main() {
	ary := []int{5, 9, 3, 1, 2, 8, 4, 7, 6}
	// bubbleSortFunc(ary)
	// selectionSortFunc(ary)
	// insertionSortFunc(ary)
	// heapSortFunc(ary)
	// mergeSortFunc(ary)
	quickSortFunc(ary)
}

func bubbleSortFunc(ary []int) {
	bs := _bubbleSortUsecase.NewBubbleSortUsecase(ary)
	fmt.Printf("before: %d\n", bs.PrintAry())
	if err := bs.Sort(); err != nil {
		fmt.Printf("err: %s\n", err.Error())
	}
	fmt.Printf("after: %d\n", bs.PrintAry())
}

func selectionSortFunc(ary []int) {
	ss := _selectionSortUsecase.NewSelectionSortUsecase(ary)
	fmt.Printf("before: %d\n", ss.PrintAry())
	if err := ss.Sort(); err != nil {
		fmt.Printf("err: %s\n", err.Error())
	}
	fmt.Printf("after: %d\n", ss.PrintAry())
}

func insertionSortFunc(ary []int) {
	is := _insertionSortUsecase.NewInsertionSortUsecase(ary)
	fmt.Printf("before: %d\n", is.PrintAry())
	if err := is.Sort(); err != nil {
		fmt.Printf("err: %s\n", err.Error())
	}
	fmt.Printf("after: %d\n", is.PrintAry())
}

func heapSortFunc(ary []int) {
	hs := _heapSortUsecase.NewHeapSortUsecase(ary)
	fmt.Printf("before: %d\n", hs.PrintAry())
	if err := hs.Sort(); err != nil {
		fmt.Printf("err: %s\n", err.Error())
	}
	fmt.Printf("after: %d\n", hs.PrintAry())
}

func mergeSortFunc(ary []int) {
	ms := _mergeSortUsecase.NewMergeSortUsecase(ary)
	fmt.Printf("before: %d\n", ms.PrintAry())
	if err := ms.Sort(); err != nil {
		fmt.Printf("err: %s\n", err.Error())
	}
	fmt.Printf("after: %d\n", ms.PrintAry())
}

func quickSortFunc(ary []int) {
	qs := _quickSortUsecase.NewQuickSortUsecase(ary)
	fmt.Printf("before: %d\n", qs.PrintAry())
	if err := qs.Sort(); err != nil {
		fmt.Printf("err: %s\n", err.Error())
	}
	fmt.Printf("after: %d\n", qs.PrintAry())
}
