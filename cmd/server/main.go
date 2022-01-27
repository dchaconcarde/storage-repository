package main

import (
	"database/sql"
	"fmt"

	"github.com/dchaconcarde/storage/cmd/server/handler"
	"github.com/dchaconcarde/storage/internal/product"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "meli_sprint_user:Meli_Sprint#123@/storage")
	if err != nil {
		fmt.Println("Error de conexión en BD")
	} else {
		fmt.Println("Conexión exitosa con BD")
	}
	r := gin.Default()

	repo := product.NewRepo(db)
	service := product.NewService(repo)
	handler := handler.NewProduct(service)

	productsGroup := r.Group("/products")

	productsGroup.GET("/name/:name", handler.GetByName())
	productsGroup.GET("/:id", handler.GetById())
	productsGroup.POST("/", handler.Create())
	productsGroup.GET("/", handler.GetAll())
	productsGroup.PATCH("/:id", handler.Update())

	if err := r.Run(); err != nil {
		panic(err)
	}

}
