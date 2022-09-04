package repository

import "belajar-unit-test/entity"

// agar mudah di test sebelum connect ke database perlu contract interface
type CategoryRepository interface {
	FindById(id string) *entity.Category
}
