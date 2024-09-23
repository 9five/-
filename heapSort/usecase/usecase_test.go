package usecase_test

import (
	"testing"
	// "algorithm/domain/mocks"
	ucase "algorithm/heapSort/usecase"

	"github.com/stretchr/testify/assert"

	// "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type HeapSortUcaseSuite struct {
	suite.Suite
	ary []int
}

func TestStart(t *testing.T) {
	suite.Run(t, &HeapSortUcaseSuite{})
}

func (s *HeapSortUcaseSuite) SetupTest() {}

func (s *HeapSortUcaseSuite) TestSort_Success() {
	u := ucase.NewHeapSortUsecase([]int{34, 37, 89, 8, 92, 22, 41, 85, 70, 59})
	err := u.Sort()
	assert.NoError(s.Suite.T(), err)
}
