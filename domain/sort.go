package domain

import ()

type SortAlgorithm struct {
}

type SortAlgorithmUsecase interface {
	PrintAry() []int
	Sort() error
}
