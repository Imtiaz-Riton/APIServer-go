package repository

import "github.com/Imtiaz-Riton/domain/entity"

type ProductRepository interface {
	GetAll() ([]*entity.Product, error)
	GetByID(id string) (*entity.Product, error)
	GetByTitle(title string) (*entity.Product, error)
	Create(product *entity.Product) (*entity.Product, error)
	Update(id string, product *entity.Product) (*entity.Product, error)
	Delete(id string) error
}
