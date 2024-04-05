package usecase

import (
	"algorithm/domain"
	"fmt"
)

type MergeSortUsecase struct {
	ary []int
}

func NewMergeSortUsecase(ary []int) domain.SortAlgorithmUsecase {
	return &MergeSortUsecase{
		ary: ary,
	}
}

func (ms *MergeSortUsecase) mergeSort(a []int) []int {
	if len(a) <= 1 {
		return a
	}
	var leftA, rightA []int
	mid := len(a) / 2
	for i := 0; i < mid; i++ {
		leftA = append(leftA, a[i])
	}
	for i := mid; i < len(a); i++ {
		rightA = append(rightA, a[i])
	}
	leftA = ms.mergeSort(leftA)
	rightA = ms.mergeSort(rightA)
	return ms.merge(leftA, rightA)
}

func (ms *MergeSortUsecase) merge(leftA, rightA []int) (result []int) {
	fmt.Println("leftA:", leftA, "rightA:", rightA)
	for len(leftA) > 0 && len(rightA) > 0 {
		if leftA[0] <= rightA[0] {
			result = append(result, leftA[0])
			leftA = append([]int{}, leftA[1:]...)
		} else {
			result = append(result, rightA[0])
			rightA = append([]int{}, rightA[1:]...)
		}
	}
	for _, v := range leftA {
		result = append(result, v)
	}
	for _, v := range rightA {
		result = append(result, v)
	}
	fmt.Println("result:", result)

	return
}

func (ms *MergeSortUsecase) PrintAry() []int {
	return ms.ary
}

func (ms *MergeSortUsecase) Sort() error {
	ms.ary = ms.mergeSort(ms.ary)

	return nil
}
