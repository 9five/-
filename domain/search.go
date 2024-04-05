package domain

import ()

type SearchAlgorithm struct {
}

type SearchAlgorithmUsecase interface {
	Search(keyword string) (int, error)
}
