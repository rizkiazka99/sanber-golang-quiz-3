package repository

import (
	"database/sql"
	"fmt"
	"quiz3/config"
	"quiz3/models"
)

func CreateCategory(c models.Category) {
	var category models.Category

	sqlStatement := `
	INSERT INTO categories (id, name, created_at, created_by, modified_at, modified_by)
	VALUES ($1, $2, $3, $4, $5, $6)
	Returning *
	`

	config.Err = config.Db.QueryRow(
		sqlStatement,
		c.Id,
		c.Name,
		c.CreatedAt,
		c.CreatedBy,
		c.ModifiedAt,
		c.ModifiedBy,
	).Scan(
		&category.Id,
		&category.Name,
		&category.CreatedAt,
		&category.CreatedBy,
		&category.ModifiedAt,
		&category.ModifiedBy,
	)

	if config.Err != nil {
		panic(config.Err)
	} else {
		fmt.Printf("Category: %+v\n", category)
	}
}

func GetCategories() ([]models.Category, error) {
	var results []models.Category

	sqlStatement := `SELECT * from categories`

	rows, err := config.Db.Query(sqlStatement)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var category = models.Category{}

		err = rows.Scan(
			&category.Id,
			&category.Name,
			&category.CreatedAt,
			&category.CreatedBy,
			&category.ModifiedAt,
			&category.ModifiedBy,
		)

		if err != nil {
			return nil, err
		}

		results = append(results, category)
	}

	fmt.Println("Data:", results)
	return results, nil
}

func GetCategoryById(id int64) (models.Category, error) {
	var category models.Category

	sqlStatement := `SELECT * from categories WHERE id = $1`

	row := config.Db.QueryRow(sqlStatement, id)

	err := row.Scan(
		&category.Id,
		&category.Name,
		&category.CreatedAt,
		&category.CreatedBy,
		&category.ModifiedAt,
		&category.ModifiedBy,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return category, nil
		} else {
			return category, err
		}
	}

	return category, nil
}

func GetBooksByCategoryId(id int64) ([]models.Book, error) {
	var results []models.Book

	sqlStatement := `SELECT * from books WHERE category_id = $1`

	rows, err := config.Db.Query(sqlStatement, id)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var book = models.Book{}

		err = rows.Scan(
			&book.Id,
			&book.Title,
			&book.Description,
			&book.ImageUrl,
			&book.ReleaseYear,
			&book.Price,
			&book.TotalPage,
			&book.Thickness,
			&book.CategoryId,
			&book.CreatedAt,
			&book.CreatedBy,
			&book.ModifiedAt,
			&book.ModifiedBy,
		)

		if err != nil {
			return nil, err
		} else {
			results = append(results, book)
		}
	}

	fmt.Println("Books:", results)
	return results, nil
}

func UpdateCategory(id int64, category models.Category) (int64, error) {
	sqlStatement := `
	UPDATE categories
	SET name = $2, modified_by = $3, modified_at = $4
	WHERE id = $1;`

	res, err := config.Db.Exec(
		sqlStatement,
		id,
		category.Name,
		category.ModifiedBy,
		category.ModifiedAt,
	)
	if err != nil {
		return 0, err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return 0, err
	} else {
		return count, nil
	}
}

func DeleteCategory(id int64) (int64, error) {
	sqlStatement := `DELETE from categories WHERE id = $1`

	res, err := config.Db.Exec(sqlStatement, id)
	if err != nil {
		return 0, err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return 0, err
	} else {
		return count, nil
	}
}
