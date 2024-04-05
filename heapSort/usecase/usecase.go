package usecase

import (
	"algorithm/domain"
	"fmt"
)

type HeapSortUsecase struct {
	ary []int
}

func NewHeapSortUsecase(ary []int) domain.SortAlgorithmUsecase {
	return &HeapSortUsecase{
		ary: ary,
	}
}

// formula of child nodes: parent * 2 + 1 && parent * 2 + 2
// formula of parent node: (child - 1) / 2

func (hs *HeapSortUsecase) heapify(nodePosition int, size int) {
	if nodePosition >= size {
		return
	}

	child1 := nodePosition*2 + 1
	child2 := nodePosition*2 + 2
	maximum := nodePosition
	if child1 < size && hs.ary[child1] > hs.ary[maximum] {
		maximum = child1
	}
	if child2 < size && hs.ary[child2] > hs.ary[maximum] {
		maximum = child2
	}
	if maximum != nodePosition {
		temp := hs.ary[nodePosition]
		hs.ary[nodePosition] = hs.ary[maximum]
		hs.ary[maximum] = temp
		hs.heapify(maximum, size)
	}
}

func (hs *HeapSortUsecase) buildHeap() {
	lastParent := (len(hs.ary) - 2) / 2 // minus two here its because size = maximum position + 1
	for i := lastParent; i >= 0; i-- {
		hs.heapify(i, len(hs.ary))
	}
}

func (hs *HeapSortUsecase) PrintAry() []int {
	return hs.ary
}

func (hs *HeapSortUsecase) Sort() error {
	hs.buildHeap()
	fmt.Printf("after buildHeap: \n%v\n", hs.ary)
	for i := len(hs.ary) - 1; i >= 0; i-- {
		temp := hs.ary[i]
		hs.ary[i] = hs.ary[0]
		hs.ary[0] = temp
		hs.heapify(0, i)

		fmt.Println(hs.ary)
	}

	return nil
}
