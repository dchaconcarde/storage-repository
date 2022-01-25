package handler

import (
	"github.com/dchaconcarde/storage/internal/product"
	"github.com/dchaconcarde/storage/pkg/web"
	"github.com/gin-gonic/gin"
)

func NewProduct(productService product.Service) *Product {
	return &Product{
		productService: productService,
	}
}

type Product struct {
	productService product.Service
}

type request struct {
	Name  string  `json:"nombre"`
	Type  string  `json:"tipo"`
	Count int     `json:"cantidad"`
	Price float64 `json:"precio"`
}

func (p *Product) GetByName() gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Param("name")
		if name == "" {
			web.Error(c, 400, product.ErrInvalidName.Error())
			return
		}
		found, err := p.productService.GetByName(c, name)
		if err != nil {
			web.Error(c, 404, product.ErrNotFound.Error())
			return
		}
		web.Success(c, 200, found)
	}
}

func (p *Product) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			web.Error(c, 422, err.Error())
			return
		}

		product, err := p.productService.Save(c, req.Name, req.Type, req.Count, req.Price)

		if err != nil {
			web.Error(c, 409, err.Error())
			return
		}
		web.Success(c, 201, product)

	}
}
