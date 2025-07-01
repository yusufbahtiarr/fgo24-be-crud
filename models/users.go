package models

import (
	"context"
	"encoding/json"
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
	results := utils.RedisClient.Exists(context.Background(), "all-users")
	if results.Val() == 0 {
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

		encoded, err := json.Marshal(users)
		if err != nil {
			return []User{}, err
		}

		utils.RedisClient.Set(context.Background(), "all-users", string(encoded), 0)

		return users, nil
	} else {
		data := utils.RedisClient.Get(context.Background(), "all-users")
		str := data.Val()

		var users []User
		err := json.Unmarshal([]byte(str), &users)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal users from redis: %w", err)
		}

		return users, nil
	}
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
