package product

import (
	"database/sql"
	"net/http/httptest"
	"testing"

	"github.com/dchaconcarde/storage/internal/domain"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func TestCreateOK(t *testing.T) {
	db, _ := sql.Open("mysql", "meli_sprint_user:Meli_Sprint#123@/storage")
	newProduct := domain.Product{
		Name:  "cafe",
		Type:  "negro",
		Count: 1,
		Price: 14.20,
	}

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	repo := NewRepo(db)
	res, _ := repo.Store(ctx, newProduct)

	assert.Equal(t, newProduct.Name, res.Name)
	assert.Equal(t, newProduct.Type, res.Type)
	assert.Equal(t, newProduct.Count, res.Count)
	assert.Equal(t, newProduct.Price, res.Price)
}

func TestGetByName(t *testing.T) {
	db, _ := sql.Open("mysql", "meli_sprint_user:Meli_Sprint#123@/storage")
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	nombre := "cafe"

	repo := NewRepo(db)
	res, _ := repo.GetByName(ctx, nombre)

	assert.Equal(t, nombre, res.Name)
	assert.IsType(t, domain.Product{}, res)
}

func TestGetAll(t *testing.T) {
	db, _ := sql.Open("mysql", "meli_sprint_user:Meli_Sprint#123@/storage")
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	repo := NewRepo(db)
	res, _ := repo.GetAll(ctx)
	expectedResults := 1
	assert.True(t, len(res) >= expectedResults)
}

func TestUpdateOK(t *testing.T) {
	db, _ := sql.Open("mysql", "meli_sprint_user:Meli_Sprint#123@/storage")
	newProduct := domain.Product{
		ID:    10,
		Name:  "Azucar",
		Type:  "Morena",
		Count: 1,
		Price: 14.20,
	}

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	repo := NewRepo(db)
	res, _ := repo.UpdateWithContext(ctx, newProduct)

	assert.Equal(t, newProduct.ID, res.ID)
	assert.Equal(t, newProduct.Name, res.Name)
	assert.Equal(t, newProduct.Type, res.Type)
	assert.Equal(t, newProduct.Count, res.Count)
	assert.Equal(t, newProduct.Price, res.Price)
}

func TestGetByIdOK(t *testing.T) {
	db, _ := sql.Open("mysql", "meli_sprint_user:Meli_Sprint#123@/storage")
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	newProduct := domain.Product{
		ID:    10,
		Name:  "Azucar",
		Type:  "Morena",
		Count: 1,
		Price: 14.20,
	}

	repo := NewRepo(db)
	res, _ := repo.GetById(ctx, newProduct.ID)

	assert.Equal(t, newProduct.ID, res.ID)
	assert.Equal(t, newProduct.Name, res.Name)
	assert.IsType(t, domain.Product{}, res)
}
