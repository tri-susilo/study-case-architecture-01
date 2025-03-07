package category_model

import (
	"go_crud_1/config"
	"go_crud_1/entities"
	"time"
)

func GetAll() []entities.Category {
	rows, err := config.MasterDB.Query(`SELECT * FROM categories`)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var categories []entities.Category

	for rows.Next() {
		var category entities.Category
		if err := rows.Scan(&category.Id, &category.Name, &category.CreatedAt, &category.UpdatedAt); err != nil {
			panic(err)
		}

		categories = append(categories, category)
	}

	return categories
}

func Create(category entities.Category) bool {
	currentTime := time.Now().Format("2006-01-02 15:04:05")

	result, err := config.MasterDB.Exec(`
		INSERT INTO categories (name, created_at, updated_at)
		VALUE(?, ?, ? )`,
		category.Name, currentTime, currentTime,
	)

	if err != nil {
		panic(err)
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	return lastInsertId > 0
}

func Detail(id int) entities.Category {
	row := config.MasterDB.QueryRow(`SELECT id, name FROM categories WHERE id = ?`, id)
	var category entities.Category
	if err := row.Scan(&category.Id, &category.Name); err != nil {
		panic(err.Error())
	}

	return category
}

func Update(id int, category entities.Category) bool {
	currentTime := time.Now().Format("2006-01-02 15:04:05")

	query, err := config.MasterDB.Exec(`UPDATE categories SET name = ?, Updated_at = ?  WHERE id = ?`, category.Name, currentTime, id)
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
	_, err := config.MasterDB.Exec(`DELETE FROM categories WHERE id = ?`, id)
	return err
}
