package usecase

import (
	"algorithm/domain"
	"errors"
	"strings"
)

type LinearSearchUsecase struct {
	ary []string
}

func NewLinearSearchUsecase(ary []string) domain.SearchAlgorithmUsecase {
	return &LinearSearchUsecase{
		ary: ary,
	}
}

func (ls *LinearSearchUsecase) Search(keyword string) (int, error) {
	keyword = strings.ToLower(keyword)
	for k, v := range ls.ary {
		if strings.ToLower(v) == keyword {
			return k, nil
		}
	}

	return 0, errors.New("not found")
}
