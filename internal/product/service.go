package product

import (
	"context"
	"errors"

	"github.com/dchaconcarde/storage/internal/domain"
)

var (
	ErrNotFound    = errors.New("product not found")
	ErrInvalidName = errors.New("invalid name")
	ErrInvalidId   = errors.New("invalid id")
)

type Service interface {
	GetByName(ctx context.Context, name string) (domain.Product, error)
	Save(ctx context.Context, name, productType string, count int, price float64, idWarehouse int) (domain.Product, error)
	GetAll(ctx context.Context) ([]domain.Product, error)
	Update(ctx context.Context, product domain.Product) (domain.Product, error)
	GetById(ctx context.Context, id int) (domain.Product, error)
	Delete(ctx context.Context, id int) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}
func (s *service) GetByName(ctx context.Context, name string) (domain.Product, error) {
	product, err := s.repository.GetByName(ctx, name)
	if err != nil {
		return domain.Product{}, ErrNotFound
	}
	return product, nil
}

func (s *service) GetById(ctx context.Context, id int) (domain.Product, error) {
	product, err := s.repository.GetById(ctx, id)
	if err != nil {
		return domain.Product{}, ErrNotFound
	}
	return product, nil
}

func (s *service) Save(ctx context.Context, name, productType string, count int, price float64, idWarehouse int) (domain.Product, error) {
	p := domain.Product{}
	p.Name = name
	p.Type = productType
	p.Count = count
	p.Price = price
	p.IdWarehouse = idWarehouse

	return s.repository.Store(ctx, p)
}

func (s *service) GetAll(ctx context.Context) ([]domain.Product, error) {
	return s.repository.GetAll(ctx)
}

func (s *service) Update(ctx context.Context, p domain.Product) (domain.Product, error) {
	producto, err := s.repository.GetById(ctx, p.ID)

	if err != nil {
		return domain.Product{}, ErrNotFound
	}

	updatedProduct := fieldVerifier(producto, p)

	return s.repository.UpdateWithContext(ctx, updatedProduct)

}

func (s *service) Delete(ctx context.Context, id int) error {
	return s.repository.Delete(ctx, id)
}

func fieldVerifier(product, updated domain.Product) domain.Product {
	if updated.Name == "" {
		updated.Name = product.Name
	}
	if updated.Type == "" {
		updated.Type = product.Type
	}
	if updated.Count == 0 {
		updated.Count = product.Count
	}
	if updated.Price == 0 {
		updated.Price = product.Price
	}
	if updated.IdWarehouse == 0 {
		updated.IdWarehouse = product.IdWarehouse
	}
	return updated
}
