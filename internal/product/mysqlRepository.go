package product

import (
	"database/sql"

	"github.com/bootcamp-go/consignas-go-db.git/internal/domain"
)

type MySQLRepository struct {
	database *sql.DB
}

func NewMySQLRepository(db *sql.DB) Repository {
	return &MySQLRepository{database: db}
}

func (repository *MySQLRepository) GetByID(id int) (domain.Product, error) {

	rows, err := repository.database.Query("SELECT id, name, quantity, code_value, price, expiration FROM products WHERE id=?", id)

	if err != nil {
		return domain.Product{}, err
	}

	var product domain.Product
	for rows.Next() {
		err := rows.Scan(&product.Id, &product.Name, &product.Quantity, &product.CodeValue, &product.Price, &product.Expiration)

		if err != nil {
			return domain.Product{}, err
		}
	}

	return product, nil
}

func (repository *MySQLRepository) Create(p domain.Product) (domain.Product, error) {

	stmt, err := repository.database.Prepare("INSERT INTO products (name, quantity, code_value, is_published, expiration, price) VALUES (?, ?, ?, ?, ?, ?)")

	if err != nil {
		return domain.Product{}, err
	}

	defer stmt.Close()

	var result sql.Result

	result, err = stmt.Exec(p.Name, p.Quantity, p.CodeValue, p.IsPublished, p.Expiration, p.Price)

	if err != nil {
		return domain.Product{}, err
	}

	insertId, _ := result.LastInsertId()

	p.Id = int(insertId)

	return p, nil
}
