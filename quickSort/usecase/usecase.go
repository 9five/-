package usecase

import (
	"algorithm/domain"
	"fmt"
	"math/rand"
	"time"
)

type QuickSortUsecase struct {
	ary []int
}

func NewQuickSortUsecase(ary []int) domain.SortAlgorithmUsecase {
	return &QuickSortUsecase{
		ary: ary,
	}
}

func (qs *QuickSortUsecase) quickSort(a []int) (result []int) {
	if len(a) <= 1 {
		return a
	}
	rand.Seed(time.Now().Unix())
	pivot := a[rand.Intn(len(a))]
	var leftA, rightA []int
	for _, v := range a {
		if v == pivot {
			continue
		}
		if v <= pivot {
			leftA = append(leftA, v)
		} else {
			rightA = append(rightA, v)
		}
	}
	fmt.Println("list:", a, "pivot:", pivot, "leftA:", leftA, "rightA", rightA)
	leftA = qs.quickSort(leftA)
	rightA = qs.quickSort(rightA)
	result = append(leftA, pivot)
	result = append(result, rightA...)

	return
}

func (qs *QuickSortUsecase) PrintAry() []int {
	return qs.ary
}

func (qs *QuickSortUsecase) Sort() error {
	qs.ary = qs.quickSort(qs.ary)

	return nil
}
