package usecase

import (
	"algorithm/domain"
	"context"
	"fmt"
)

type InsertionSortUsecase struct {
	ary []int
}

func NewInsertionSortUsecase(ary []int) domain.SortAlgorithmUsecase {
	return &InsertionSortUsecase{
		ary: ary,
	}
}

func (is *InsertionSortUsecase) PrintAry() []int {
	return is.ary
}

func (is *InsertionSortUsecase) Sort(ctx context.Context) error {

	for i := 0; i < len(is.ary); i++ {
		node := is.ary[i]
		j := i - 1
		fmt.Println("node:", node)
		for j >= 0 && is.ary[j] > node {
			is.ary[j+1] = is.ary[j]
			fmt.Println(is.ary)
			j--
		}
		is.ary[j+1] = node
		fmt.Println(is.ary)
	}
	return nil
}
