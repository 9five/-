package usecase

import (
	"algorithm/domain"
	"fmt"
)

type BubbleSortUsecase struct {
	ary []int
}

func NewBubbleSortUsecase(ary []int) domain.SortAlgorithmUsecase {
	return &BubbleSortUsecase{
		ary: ary,
	}
}

func (bs *BubbleSortUsecase) PrintAry() []int {
	return bs.ary
}

func (bs *BubbleSortUsecase) Sort() error {

	var swapped bool
	for i := 0; i < len(bs.ary); i++ {
		swapped = false
		for j := 0; j < len(bs.ary)-1; j++ {
			fmt.Println(bs.ary[j], ":", bs.ary[j+1])
			if bs.ary[j] > bs.ary[j+1] {
				temp := bs.ary[j]
				bs.ary[j] = bs.ary[j+1]
				bs.ary[j+1] = temp
				swapped = true
				fmt.Println(bs.ary)
			}
		}
		if swapped == false {
			break
		}
	}

	return nil
}
