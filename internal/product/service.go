package product

import (
	"context"
	"errors"

	"github.com/dchaconcarde/storage/internal/domain"
)

var (
	ErrNotFound    = errors.New("product not found")
	ErrInvalidName = errors.New("invalid name")
)

type Service interface {
	GetByName(ctx context.Context, name string) (domain.Product, error)
	Save(ctx context.Context, name, productType string, count int, price float64) (domain.Product, error)
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
		return domain.Product{}, errors.New("product doesn't exists")
	}
	return product, nil
}

func (s *service) Save(ctx context.Context, name, productType string, count int, price float64) (domain.Product, error) {
	p := domain.Product{}
	p.Name = name
	p.Type = productType
	p.Count = count
	p.Price = price

	return s.repository.Store(ctx, p)
}
