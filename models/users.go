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

func GetAllUsers() ([]User, error) {

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

func GetUserByID(user_id string) (User, error) {
	conn, err := utils.DBConnect()
	if err != nil {
		return User{}, err
	}
	defer conn.Close()

	query := `SELECT id, username, email, password FROM users WHERE id = $1`
	id, _ := strconv.Atoi(user_id)

	rows, err := conn.Query(context.Background(), query, id)
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
	query := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3, $4))
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

func UpdateUser(id string, newdata User) error {
	// conn, err := utils.DBConnect()
	// if err != nil {
	// 	return err
	// }
	// defer conn.Conn().Close(context.Background)

	// olddata, err := GetUserByID(id)
	// if err != nil {
	// 	return err
	// }

	// if newdata.Username != "" {
	// 	olddata.Username = newdata.Username
	// }
	// if newdata.Email != "" {
	// 	olddata.Email = newdata.Email
	// }
	// if newdata.Password != "" {
	// 	olddata.Password = newdata.Password
	// }

	// _, err := conn.Exec(context.Background(), `UPDATE users SET username = $1, email = $2, password = $3 WHERE id = $4`, olddata.Username, olddata.Email, olddata.Password, id)

	// return err
	return nil
}
