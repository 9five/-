package usecase

import (
	"algorithm/domain"
	"fmt"
)

type SelectionSortUsecase struct {
	ary []int
}

func NewSelectionSortUsecase(ary []int) domain.SortAlgorithmUsecase {
	return &SelectionSortUsecase{
		ary: ary,
	}
}

func (ss *SelectionSortUsecase) PrintAry() []int {
	return ss.ary
}

func (ss *SelectionSortUsecase) Sort() error {

	var minimumPosition int
	for i := 0; i < len(ss.ary); i++ {
		for j := i + 1; j < len(ss.ary); j++ {
			if ss.ary[j] < ss.ary[minimumPosition] {
				minimumPosition = j
			}
		}
		fmt.Println("minimum:", ss.ary[minimumPosition], "position:", minimumPosition, "i:", i)
		temp := ss.ary[i]
		ss.ary[i] = ss.ary[minimumPosition]
		ss.ary[minimumPosition] = temp
		fmt.Println(ss.ary)

		minimumPosition = i + 1
	}

	return nil
}
