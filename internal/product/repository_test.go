package product

import (
	"context"
	"database/sql"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-txdb"
	"github.com/dchaconcarde/storage/internal/domain"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

const (
	sqlConnection = "meli_sprint_user:Meli_Sprint#123@/storage"
)

func init() {
	txdb.Register("txdb", "mysql", sqlConnection)
}

func InitDb() (*sql.DB, error) {
	db, err := sql.Open("txdb", uuid.New().String())
	if err == nil {
		return db, db.Ping()
	}
	defer db.Close()
	return db, err
}

func TestCreateOK(t *testing.T) {
	db, _ := sql.Open("mysql", sqlConnection)
	newProduct := domain.Product{
		Name:        "cafe",
		Type:        "negro",
		Count:       1,
		Price:       14.20,
		IdWarehouse: 1,
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
	db, _ := sql.Open("mysql", sqlConnection)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	nombre := "cafe"

	repo := NewRepo(db)
	res, _ := repo.GetByName(ctx, nombre)

	assert.Equal(t, nombre, res.Name)
	assert.IsType(t, domain.Product{}, res)
}

func TestGetAll(t *testing.T) {
	db, _ := sql.Open("mysql", sqlConnection)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	repo := NewRepo(db)
	res, _ := repo.GetAll(ctx)
	expectedResults := 1
	assert.True(t, len(res) >= expectedResults)
}

func TestUpdateOK(t *testing.T) {
	db, _ := sql.Open("mysql", sqlConnection)
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
	db, _ := sql.Open("mysql", sqlConnection)
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

func TestDeleteOk(t *testing.T) {
	db, _ := sql.Open("mysql", sqlConnection)
	s := NewRepo(db)
	err := s.Delete(context.Background(), 10)

	assert.NoError(t, err)
}
func Test_Store_Mocked(t *testing.T) {
	db, err := InitDb()
	assert.NoError(t, err)

	repository := NewRepo(db)
	ctx := context.TODO()
	product := domain.Product{
		Name:        "Refresco",
		Type:        "coca",
		Count:       5,
		Price:       4561,
		IdWarehouse: 1,
	}
	res, err := repository.Store(ctx, product)
	assert.NoError(t, err)

	assert.Equal(t, product.Name, res.Name)
	assert.Equal(t, product.Type, res.Type)
	assert.Equal(t, product.Count, res.Count)
	assert.Equal(t, product.Price, res.Price)
	assert.Equal(t, product.IdWarehouse, res.IdWarehouse)
}

func Test_GetByName_Mocked(t *testing.T) {
	db, err := InitDb()
	assert.NoError(t, err)
	ctx := context.TODO()
	nombre := "Azucar"

	repo := NewRepo(db)
	res, _ := repo.GetByName(ctx, nombre)
	assert.Equal(t, nombre, res.Name)
	assert.IsType(t, domain.Product{}, res)
}

func Test_GetById_Mocked(t *testing.T) {
	db, err := InitDb()
	assert.NoError(t, err)
	ctx := context.TODO()
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

func Test_Update_OK_Mocked(t *testing.T) {
	db, err := InitDb()
	assert.NoError(t, err)
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

func Test_Delete_OK_Mocked(t *testing.T) {
	db, errorDb := InitDb()
	assert.NoError(t, errorDb)
	s := NewRepo(db)
	err := s.Delete(context.Background(), 11)

	assert.NoError(t, err)
}
