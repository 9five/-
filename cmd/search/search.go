package main

import (
	_binarySearchUsecase "algorithm/binarySearch/usecase"
	_linearSearchUsecase "algorithm/linearSearch/usecase"
	"fmt"
)

func main() {
	ary := []string{
		"Apple",
		"Banana",
		"Orange",
		"Strawberry",
		"Grape",
		"Pineapple",
		"Watermelon",
		"Mango",
		"Kiwi",
		"Blueberry",
	}

	// linearSearchFunc(ary, "kiwi")
	binarySearchFunc(ary, "apple")
}

func linearSearchFunc(ary []string, keyword string) {
	var exist bool
	ls := _linearSearchUsecase.NewLinearSearchUsecase(ary)
	p, err := ls.Search(keyword)
	if err == nil {
		exist = true
	}
	fmt.Printf("keyword: %s\nexist: %t\nposition: %d\n", keyword, exist, p)
}

func binarySearchFunc(ary []string, keyword string) {
	var exist bool
	bs := _binarySearchUsecase.NewBinarySearchUsecase(ary)
	p, err := bs.Search(keyword)
	if err == nil {
		exist = true
	}
	fmt.Printf("keyword: %s\nexist: %t\nposition: %d\n", keyword, exist, p)
}
