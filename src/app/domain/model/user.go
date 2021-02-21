package model

type User struct {
	Id    int    `db:"user_id"`
	Name  string `db:"name"`
	Email string `db:"email"`
}
