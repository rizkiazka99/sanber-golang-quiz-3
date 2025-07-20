package models

import "time"

type User struct {
	Id         int       `json:"id"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	CreatedAt  time.Time `json:"created_at"`
	CreatedBy  string    `json:"created_by"`
	ModifiedAt time.Time `json:"modified_at"`
	ModifiedBy string    `json:"modified_by"`
}

type LoggedIn struct {
	Id                  int       `json:"id"`
	Username            string    `json:"username"`
	AccessToken         string    `json:"access_token"`
	TokenExpirationTime time.Time `json:"token_expiration_time"`
}
