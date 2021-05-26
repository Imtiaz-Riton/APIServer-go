package in_memory

import (
	"github.com/Imtiaz-Riton/domain/entity"
	"github.com/Imtiaz-Riton/domain/exception"
	"github.com/Imtiaz-Riton/domain/repository"
)

type inMemoryProductRepository struct {
	data map[string]*entity.Product
}

func (i *inMemoryProductRepository) GetAll() ([]*entity.Product, error) {
	products := make([]*entity.Product, 0)
	for _, product := range i.data {
		products = append(products, product)
	}
	return products, nil
}

func (i *inMemoryProductRepository) GetByID(id string) (*entity.Product, error) {
	if product, ok := i.data[id]; ok {
		return product, nil
	}
	return nil, exception.ErrNotFound
}

func (i *inMemoryProductRepository) GetByTitle(title string) (*entity.Product, error) {
	if product, ok := i.data[title]; ok {
		return product, nil
	}
	return nil, exception.ErrNotFound
}

func (i *inMemoryProductRepository) Create(product *entity.Product) (*entity.Product, error) {
	if _, ok := i.data[product.ID]; ok {
		return nil, exception.ErrConflict
	}
	i.data[product.ID] = product
	return product, nil
}

func (i *inMemoryProductRepository) Update(id string, product *entity.Product) (*entity.Product, error) {
	if _, ok := i.data[id]; !ok {
		return nil, exception.ErrNotFound
	}
	delete(i.data, id)
	i.Create(product)
	return product, nil
}

func (i *inMemoryProductRepository) Delete(id string) error {
	if _, ok := i.data[id]; !ok {
		return exception.ErrNotFound
	}
	delete(i.data, id)
	return nil
}

func NewInMemoryProductRepository(initialData map[string]*entity.Product) repository.ProductRepository {
	return &inMemoryProductRepository{data: initialData}
}
