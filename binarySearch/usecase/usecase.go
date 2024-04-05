package usecase

import (
	"algorithm/domain"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type BinarySearchUsecase struct {
	ary []string
}

func NewBinarySearchUsecase(ary []string) domain.SearchAlgorithmUsecase {
	return &BinarySearchUsecase{
		ary: ary,
	}
}

func (bs *BinarySearchUsecase) quickSortByASCII(a []string) (result []string) {
	if len(a) <= 1 {
		return a
	}
	rand.Seed(time.Now().Unix())
	pivot := a[rand.Intn(len(a))]
	var leftA, rightA []string
	for _, v := range a {
		if v == pivot {
			continue
		}

		for i := 0; i < len(pivot); i++ {
			ascii1, _ := strconv.Atoi(fmt.Sprintf("%d", v[i]))
			ascii2, _ := strconv.Atoi(fmt.Sprintf("%d", pivot[i]))

			if i > len(v)-1 || ascii1 < ascii2 {
				leftA = append(leftA, v)
				break
			} else if (i == len(pivot)-1 && i < len(v)-1) || ascii1 > ascii2 {
				rightA = append(rightA, v)
				break
			}
		}
	}
	leftA = bs.quickSortByASCII(leftA)
	rightA = bs.quickSortByASCII(rightA)
	result = append(leftA, pivot)
	result = append(result, rightA...)

	return
}

func (bs *BinarySearchUsecase) Search(keyword string) (int, error) {
	keyword = strings.ToLower(keyword)
	bs.ary = bs.quickSortByASCII(bs.ary)
	fmt.Println(bs.ary)
	leftSign := 0
	rightSign := len(bs.ary) - 1
	for leftSign <= rightSign {
		mid := (leftSign + rightSign) / 2

		if strings.ToLower(string(bs.ary[mid])) == keyword {
			return mid, nil
		}
		for i := 0; i < len(bs.ary[mid]); i++ {
			keywordASCII, _ := strconv.Atoi(fmt.Sprintf("%d", keyword[i]))
			aryASCII, _ := strconv.Atoi(fmt.Sprintf("%d", strings.ToLower(bs.ary[mid])[i]))

			if keywordASCII < aryASCII {
				rightSign = mid - 1
				break
			} else if keywordASCII > aryASCII {
				leftSign = mid + 1
				break
			}
		}
	}

	return 0, errors.New("not found")
}
