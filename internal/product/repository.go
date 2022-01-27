package product

import (
	"context"
	"database/sql"
	"log"

	"github.com/dchaconcarde/storage/internal/domain"
)

const (
	getProductByName = "SELECT * FROM products where name = ?"
	createNewProduct = "INSERT INTO products(name, type, count, price, id_warehouse) VALUES( ?, ?, ?, ?, ?)"
	getAllProducts   = "SELECT * FROM products"
	updateProduct    = "UPDATE products SET name=?, type=?, count=?, price=?, id_warehouse=? WHERE id=?"
	getProductById   = "SELECT * FROM products WHERE id=?;"
)

type Repository interface {
	GetByName(ctx context.Context, name string) (domain.Product, error)
	Store(ctx context.Context, product domain.Product) (domain.Product, error)
	GetAll(ctx context.Context) ([]domain.Product, error)
	UpdateWithContext(ctx context.Context, product domain.Product) (domain.Product, error)
	GetById(ctx context.Context, id int) (domain.Product, error)
}
type repository struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetByName(ctx context.Context, name string) (domain.Product, error) {
	var product domain.Product

	rows, err := r.db.Query(getProductByName, name)
	if err != nil {
		log.Println(err)
		return domain.Product{}, err
	}
	for rows.Next() {
		if err := rows.Scan(&product.ID, &product.Name, &product.Type, &product.Count, &product.Price, &product.IdWarehouse); err != nil {
			log.Println(err.Error())
			return domain.Product{}, err
		}
	}
	return product, nil
}

func (r *repository) Store(ctx context.Context, product domain.Product) (domain.Product, error) {

	stmt, err := r.db.Prepare(createNewProduct) // se prepara el SQL
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() // se cierra la sentencia al terminar. Si quedan abiertas se genera consumos de memoria
	var result sql.Result
	result, err = stmt.Exec(product.Name, product.Type, product.Count, product.Price, product.IdWarehouse) // retorna un sql.Result y un error
	if err != nil {
		return domain.Product{}, err
	}
	insertedId, _ := result.LastInsertId() // del sql.Resul devuelto en la ejecución obtenemos el Id insertado
	product.ID = int(insertedId)
	return product, nil
}

func (r *repository) GetAll(ctx context.Context) ([]domain.Product, error) {
	var products []domain.Product
	rows, err := r.db.Query(getAllProducts)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	// se recorren todas las filas
	for rows.Next() {
		// por cada fila se obtiene un objeto del tipo Product
		var product domain.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Type, &product.Count, &product.Price, &product.IdWarehouse); err != nil {
			log.Fatal(err)
			return nil, err
		}
		//se añade el objeto obtenido al slice products
		products = append(products, product)
	}
	return products, nil
}

func (r *repository) UpdateWithContext(ctx context.Context, product domain.Product) (domain.Product, error) {
	query := updateProduct

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return domain.Product{}, err
	}

	res, err := stmt.ExecContext(ctx, product.Name, product.Type, product.Count, product.Price, product.IdWarehouse, product.ID)
	if err != nil {
		return domain.Product{}, err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return domain.Product{}, err
	}

	return product, nil
}

func (r *repository) GetById(ctx context.Context, id int) (domain.Product, error) {
	query := getProductById
	row := r.db.QueryRow(query, id)
	product := domain.Product{}
	err := row.Scan(&product.ID, &product.Name, &product.Type, &product.Count, &product.Price, &product.IdWarehouse)
	if err != nil {
		return domain.Product{}, err
	}

	return product, nil
}
