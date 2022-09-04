package repository

import (
	"belajar-unit-test/entity"
	"fmt"

	"github.com/stretchr/testify/mock"
)

type CategoryRepositoryMock struct {
	Mock mock.Mock
}

func (repository *CategoryRepositoryMock) FindById(id string) *entity.Category {
	argument := repository.Mock.Called(id)
	if argument.Get(0) == nil {
		return nil
	} else {
		category := argument.Get(0).(entity.Category)
		fmt.Println(category)
		return &category
	}
}
