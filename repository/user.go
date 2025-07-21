package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"quiz3/config"
	"quiz3/middleware"
	"quiz3/models"
)

func CreateUser(user models.User) string {
	var exists bool
	var userCredentials models.User

	existQuery := `SELECT EXISTS (SELECT 1 FROM users WHERE username = $1)`
	e := config.Db.QueryRow(existQuery, user.Username).Scan(&exists)
	if e != nil {
		return "Something went wrong"
	} else if exists {
		return "Username has been taken"
	} else {
		sqlStatement := `
		INSERT INTO users (id, username, password, created_at, created_by, modified_at, modified_by)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		Returning *
		`
		config.Err = config.Db.QueryRow(
			sqlStatement,
			user.Id,
			user.Username,
			user.Password,
			user.CreatedAt,
			user.CreatedBy,
			user.ModifiedAt,
			user.ModifiedBy,
		).Scan(
			&userCredentials.Id,
			&userCredentials.Username,
			&userCredentials.Password,
			&userCredentials.CreatedAt,
			&userCredentials.CreatedBy,
			&userCredentials.ModifiedAt,
			&userCredentials.ModifiedBy,
		)

		if config.Err != nil {
			panic(config.Err)
		} else {
			fmt.Printf("User: %+v\n", userCredentials)
		}

		return ""
	}
}

func Login(username string, password string) (*models.User, error) {
	var user models.User

	sqlStatement := `
	SELECT id, username, password, created_at, created_by, modified_at 
	FROM users 
	WHERE username = $1 LIMIT 1`

	config.Err = config.Db.QueryRow(sqlStatement, username).Scan(
		&user.Id,
		&user.Username,
		&user.Password,
		&user.CreatedAt,
		&user.CreatedBy,
		&user.ModifiedAt,
	)

	if config.Err != nil {
		if errors.Is(config.Err, sql.ErrNoRows) {
			return nil, errors.New("account doesn't exist")
		} else {
			return nil, config.Err
		}
	} else {
		if err := middleware.CheckPasswordHash(password, user.Password); !err {
			return nil, errors.New("incorrect username or password")
		} else {
			return &user, nil
		}
	}
}
