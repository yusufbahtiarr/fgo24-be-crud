package models

import (
	"context"
	"fgo24-be-crud/utils"
	"strconv"

	"github.com/jackc/pgx/v5"
)

type User struct {
	Id       int    `form:"id" db:"id"`
	Username string `form:"username" db:"username"`
	Email    string `form:"email" db:"email"`
	Password string `form:"password" db:"password"`
}

var Users []User

func AllUser(search string) ([]User, error) {

	conn, err := utils.DBConnect()
	if err != nil {
		return []User{}, err
	}
	defer conn.Close()

	query := `SELECT id, username, email, password FROM users`

	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		return []User{}, err
	}
	defer rows.Close()

	users, err := pgx.CollectRows[User](rows, pgx.RowToStructByName)
	if err != nil {
		return []User{}, err
	}

	return users, nil
}

func UserById(user_id string) ([]User, error) {
	conn, err := utils.DBConnect()
	if err != nil {
		return []User{}, err
	}
	defer conn.Close()

	query := `SELECT id, username, email, password FROM users WHERE id = $1`

	id, _ := strconv.Atoi(user_id)

	rows, err := conn.Query(context.Background(), query, id)
	if err != nil {
		return []User{}, err
	}
	defer rows.Close()

	users, err := pgx.CollectRows[User](rows, pgx.RowToStructByName)
	if err != nil {
		return []User{}, err
	}

	return users, nil

}

func CreateUser(user User) error {
	conn, err := utils.DBConnect()
	if err != nil {
		return err
	}
	defer conn.Close()

	query := `INSERT INTO users (username, email password) VALUES ($1, $2, $3, md5($4))
	`

	_, err = conn.Exec(context.Background(), query, user.Username, user.Email, user.Password)
	if err != nil {
		return err
	}

	return nil
}

func DeleteUser(id string) error {
	conn, err := utils.DBConnect()
	if err != nil {
		return err
	}
	defer conn.Close()

	query := `DELETE FROM users WHERE id = $1`

	idInt, _ := strconv.Atoi(id)

	_, err = conn.Exec(context.Background(), query, idInt)
	if err != nil {
		return err
	}

	return nil
}
