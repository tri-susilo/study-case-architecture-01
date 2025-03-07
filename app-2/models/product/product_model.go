package product_model

import (
	"go_crud_1/config"
	"go_crud_1/entities"
	"time"
)

func GetAll() []entities.Product {
	rows, err := config.MasterDB.Query(`
		SELECT
			products.id,
			products.name,
			categories.name as category_name,
			products.stock,
			products.description,
			products.created_at,
			products.updated_at
		FROM products
		JOIN categories ON products.category_id = categories.id
	`)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var products []entities.Product
	for rows.Next() {
		var product entities.Product
		err := rows.Scan(
			&product.Id,
			&product.Name,
			&product.Category.Name,
			&product.Stock,
			&product.Description,
			&product.CreatedAt,
			&product.UpdatedAt,
		)

		if err != nil {
			panic(err)
		}

		products = append(products, product)
	}

	return products
}

func Create(product entities.Product) bool {
	currentTime := time.Now().Format("2006-01-02 15:04:05")

	result, err := config.MasterDB.Exec(`
		INSERT INTO products(
			name, category_id, stock, description, created_at, updated_at
		) VALUES (?,?,?,?,?,?)`,
		product.Name,
		product.Category.Id,
		product.Stock,
		product.Description,
		currentTime,
		currentTime,
	)

	if err != nil {
		panic(err)
	}

	LastInsertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	return LastInsertId > 0
}

func Detail(id int) entities.Product {
	row := config.MasterDB.QueryRow(`
		SELECT
			products.id,
			products.name,
			categories.name as category_name,
			products.stock,
			products.description,
			products.created_at,
			products.updated_at
		FROM products
		JOIN categories ON products.category_id = categories.id
		WHERE products.id = ?
	`, id)

	var product entities.Product

	err := row.Scan(
		&product.Id,
		&product.Name,
		&product.Category.Name,
		&product.Stock,
		&product.Description,
		&product.CreatedAt,
		&product.UpdatedAt,
	)

	if err != nil {
		panic(err)
	}

	return product
}

func Update(id int, product entities.Product) bool {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	query, err := config.MasterDB.Exec(`
		UPDATE products SET
			name = ?,
			category_id = ?,
			stock = ?,
			description = ?,
			updated_at = ?
		WHERE id = ?
	`,
		product.Name,
		product.Category.Id,
		product.Stock,
		product.Description,
		currentTime,
		id,
	)

	if err != nil {
		panic(err)
	}

	result, err := query.RowsAffected()
	if err != nil {
		panic(err)
	}

	return result > 0
}

func Delete(id int) error {
	_, err := config.MasterDB.Exec(`DELETE FROM products WHERE id = ?`, id)
	if err != nil {
		panic(err)
	}

	return err
}
