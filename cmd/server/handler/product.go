package handler

import (
	"strconv"

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
	Name        string  `json:"nombre"`
	Type        string  `json:"tipo"`
	Count       int     `json:"cantidad"`
	Price       float64 `json:"precio"`
	IdWarehouse int     `json:"id_warehouse"`
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

		product, err := p.productService.Save(c, req.Name, req.Type, req.Count, req.Price, req.IdWarehouse)

		if err != nil {
			web.Error(c, 409, err.Error())
			return
		}
		web.Success(c, 201, product)

	}
}

func (p *Product) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		products, err := p.productService.GetAll(c)

		if err != nil {
			web.Error(c, 400, "error al obtener lista de productos")
			return
		}

		if len(products) == 0 {
			web.Error(c, 404, product.ErrNotFound.Error())
			return
		}
		web.Success(c, 200, products)
	}
}

func (p *Product) GetById() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			web.Error(c, 400, product.ErrInvalidId.Error())
			return
		}
		found, err := p.productService.GetById(c, int(id))
		if err != nil {
			web.Error(c, 404, product.ErrNotFound.Error())
			return
		}
		web.Success(c, 200, found)
	}
}

func (p *Product) Update() gin.HandlerFunc {
	return func(c *gin.Context) {

		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			web.Error(c, 400, product.ErrInvalidId.Error())
			return
		}
		found, err := p.productService.GetById(c, int(id))
		if err != nil {
			web.Error(c, 404, product.ErrNotFound.Error())
			return
		}
		foundBeforeUpdate := found
		if err := c.ShouldBindJSON(&found); err != nil {
			web.Error(c, 400, err.Error())
			return
		}
		if foundBeforeUpdate == found {
			web.Error(c, 400, "no hubo ningún cambio en el producto")
			return
		}

		product, err := p.productService.Update(c, found)
		if err != nil {
			web.Error(c, 400, err.Error())
		}
		web.Success(c, 200, product)

	}
}

func (p *Product) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			web.Error(c, 400, product.ErrInvalidId.Error())
			return
		}
		err = p.productService.Delete(c, int(id))
		if err != nil {
			web.Error(c, 404, product.ErrNotFound.Error())
			return
		}
		web.Success(c, 204, "")

	}
}
