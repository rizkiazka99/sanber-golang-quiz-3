package repository

import (
	"database/sql"
	"fmt"
	"quiz3/config"
	"quiz3/models"
)

func CreateBook(b models.Book) {
	var book models.Book

	sqlStatement := `
	INSERT INTO books (id, title, description, image_url, release_year, price, total_page, thickness, category_id, created_at, created_by, modified_at, modified_by)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
	Returning *
	`

	config.Err = config.Db.QueryRow(
		sqlStatement,
		b.Id,
		b.Title,
		b.Description,
		b.ImageUrl,
		b.ReleaseYear,
		b.Price,
		b.TotalPage,
		b.Thickness,
		b.CategoryId,
		b.CreatedAt,
		b.CreatedBy,
		b.ModifiedAt,
		b.ModifiedBy,
	).Scan(
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

	if config.Err != nil {
		panic(config.Err)
	} else {
		fmt.Printf("Book: %+v\n", book)
	}
}

func GetBooks() ([]models.Book, error) {
	var results []models.Book

	sqlStatement := `SELECT * from books`

	rows, err := config.Db.Query(sqlStatement)

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
		}

		results = append(results, book)
	}

	fmt.Println("Data:", results)
	return results, nil
}

func GetBookById(id int64) (models.Book, error) {
	var book models.Book

	sqlStatement := `SELECT * from books WHERE id = $1`

	row := config.Db.QueryRow(sqlStatement, id)

	err := row.Scan(
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
		if err == sql.ErrNoRows {
			return book, nil
		} else {
			return book, err
		}
	}

	return book, nil
}

func UpdateBook(id int64, book models.Book) (int64, error) {
	sqlStatement := `
	UPDATE books
	SET title = $2, description = $3, image_url = $4, release_year = $5, price = $6, total_page = $7, thickness = $8, category_id = $9, modified_at = $10, modified_by = $11
	WHERE id = $1;`

	res, err := config.Db.Exec(
		sqlStatement,
		id,
		book.Title,
		book.Description,
		book.ImageUrl,
		book.ReleaseYear,
		book.Price,
		book.TotalPage,
		book.Thickness,
		book.CategoryId,
		book.ModifiedAt,
		book.ModifiedBy,
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

func DeleteBook(id int64) (int64, error) {
	sqlStatement := `DELETE from books WHERE id = $1`

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
