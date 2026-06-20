package repo

import (
	db "backend/internal/database"
	"backend/internal/model"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

func GetUserByID(id int) (*model.User, error) {
	connection, err := db.Connect()

	var user model.User

	if err != nil {
		return nil, fmt.Errorf("connection error in get user %w ", err)
	}

	defer connection.Close()

	err = connection.Conn.QueryRow(
		"SELECT id username email password_hashed created_at FROM users WHERE id=?",
		id,
	).Scan(&user.ID, &user.UserName, &user.Email, &user.PasswordHashed, &user.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, fmt.Errorf("error in data base reading %W", err)
	}

	return &user, nil
}

func DeleteUserById(id int) (bool, error) {
	conn, err := db.Connect()

	if err != nil {
		return false, fmt.Errorf("connection error in get user %w ", err)
	}

	defer conn.Close()

	res, err := conn.Conn.Exec("DELETE FROM users WHERE id=?", id)

	if err != nil {
		return false, fmt.Errorf("sql error while deleting user %w", err)
	}

	rowEffected, err := res.RowsAffected()

	if err != nil {
		return false, err
	}

	if rowEffected == 0 {
		return false, nil
	}

	return true, nil
}

func AddUser(user model.User) (int64, error) {
	conn, err := db.Connect()

	if err != nil {
		return 0, fmt.Errorf("connection error in add user %w ", err)
	}
	defer conn.Close()

	res, err := conn.Conn.Exec(
		"INSERT INTO users (username, email, password_hashed, created_at) VALUES (?,?,?,?)",
		user.UserName,
		user.Email,
		user.PasswordHashed,
		user.CreatedAt,
	)

	if err != nil {
		return 0, fmt.Errorf("sql error adding user %w", err)
	}

	last_id, err := res.LastInsertId()

	if err != nil {
		return 0, fmt.Errorf("error geting insert id %w", err)
	}

	return last_id, nil
}

func UpdateUser(user model.User) (bool, error) {
	conn, err := db.Connect()

	if err != nil {
		return false, fmt.Errorf("connection error in get user %w ", err)
	}
	defer conn.Close()

	res, err := conn.Conn.Exec(
		"UPDATE users  SET username=?, email=?, password_hashed=?",
		user.UserName,
		user.Email,
		user.PasswordHashed,
	)

	if err != nil {
		return false, fmt.Errorf("user update error %w", err)
	}
	rowAffected, err := res.RowsAffected()

	if err != nil {
		return false, err
	}

	if rowAffected == 0 {
		return false, nil
	}
	return true, nil
}

func GetAllUsers() ([]model.User, error) {
	conn, err := db.Connect()

	if err != nil {
		return nil, fmt.Errorf("error connection to db hile getting all users %w", err)
	}
	defer conn.Close()

	res, err := conn.Conn.Query(
		"SELECT id,username,email,password_hashed,created_at FROM user",
	)

	var users []model.User

	for res.Next() {
		var u model.User
		err = res.Scan(&u.ID, &u.Email, &u.PasswordHashed, &u.CreatedAt)
		if err != nil {
			log.Println("skipping corrupted user", err)
			continue
		}
		users = append(users, u)
	}
	if err = res.Err(); err != nil {
		return nil, err
	}
	return users, nil
}
