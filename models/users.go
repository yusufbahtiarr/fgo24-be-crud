package models

import (
	"context"
	"fgo24-be-crud/utils"
	"fmt"
	"strconv"

	"github.com/jackc/pgx/v5"
)

type User struct {
	Id       int    `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

type CreateUserRequest struct {
	Username string `json:"username" example:"yusuf_bahtr"`
	Email    string `json:"email" example:"yusuf@gmail.com"`
	Password string `json:"password" example:"log123"`
}

type UpdateUserRequest struct {
	Username *string `json:"username" example:"yusuf_bahtr"`
	Email    *string `json:"email" example:"yusuf@gmail.com"`
	Password *string `json:"password" example:"log123"`
}

var Users []User

func FindAllUsers() ([]User, error) {
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

	users, err := pgx.CollectRows[User](rows, pgx.RowToStructByName)
	if err != nil {
		return []User{}, err
	}

	return users, nil
}

func FindUserByID(user_id int) (User, error) {
	conn, err := utils.DBConnect()
	if err != nil {
		return User{}, err
	}
	defer conn.Close()

	query := `SELECT id, username, email, password FROM users WHERE id = $1`
	rows, err := conn.Query(context.Background(), query, user_id)
	if err != nil {
		return User{}, err
	}
	defer rows.Close()

	user, err := pgx.CollectOneRow[User](rows, pgx.RowToStructByName)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func CreateUser(user User) error {
	conn, err := utils.DBConnect()
	if err != nil {
		return err
	}
	defer conn.Close()
	query := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3)`
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

func UpdateUser(id int, newdata User) error {
	conn, err := utils.DBConnect()
	if err != nil {
		return err
	}
	defer conn.Close()

	olddata, err := FindUserByID(id)
	if err != nil {
		return err
	}

	if newdata.Username == "" && newdata.Email == "" && newdata.Password == "" {
		return fmt.Errorf("input data must not be empty")
	}

	if newdata.Username != "" {
		olddata.Username = newdata.Username
	}
	if newdata.Email != "" {
		olddata.Email = newdata.Email
	}
	if newdata.Password != "" {
		olddata.Password = newdata.Password
	}

	_, err = conn.Exec(context.Background(), `UPDATE users SET username = $1, email = $2, password = $3	WHERE id = $4
	`, olddata.Username, olddata.Email, olddata.Password, id)

	return err
}
