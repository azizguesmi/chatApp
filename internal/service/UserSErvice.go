package service

import (
	"backend/internal/model"
	"backend/internal/repo"
	"fmt"
	"regexp"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func AddUser(username, email, password string) (int64, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return 0, fmt.Errorf("error while hashing the password in service")
	}
	user := model.User{
		UserName:       username,
		Email:          email,
		PasswordHashed: string(hash),
		CreatedAt:      time.Now(),
	}

	id, err := repo.AddUser(user)
	if err != nil {
		return 0, fmt.Errorf("error is saving user (service) %w", err)
	}
	return id, nil

}

func DeleteUser(id int) (bool, error) {
	test, err := repo.DeleteUserById(id)
	if err != nil {
		return false, fmt.Errorf("error in service %w", err)
	}
	return test, nil
}

func GetUserById(id int) (*model.User, error) {
	u, err := repo.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func UpdateUser(id int, username, email, password *string) (bool, error) {
	oldUser, err := GetUserById(id)
	if err != nil {
		return false, fmt.Errorf("user does not exists")
	}
	var newUser model.User
	if username == nil {
		newUser.UserName = oldUser.UserName
	} else {
		newUser.UserName = *username
	}
	if email == nil {
		newUser.Email = oldUser.Email
	} else {
		newUser.Email = *email
	}
	if password == nil {
		newUser.PasswordHashed = oldUser.PasswordHashed
	} else {
		hash, err := bcrypt.GenerateFromPassword([]byte(*password), bcrypt.DefaultCost)
		if err != nil {
			return false, fmt.Errorf("error while hashing the password in service")
		}
		newUser.PasswordHashed = string(hash)
	}

	newUser.ID = id

	test, err := repo.UpdateUser(newUser)
	if err != nil {
		return false, fmt.Errorf("error while updating user in service %w", err)
	}
	return test, err
}

func GetUserByEmail(email string) (*model.User, error) {
	users, err := repo.GetAllUsers()
	if err != nil {
		return nil, fmt.Errorf("error in getting users")
	}

	re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	if !re.MatchString(email) {
		return nil, fmt.Errorf("error in email structure")
	}
	for _, u := range users {
		if u.Email == email {
			return &u, nil
		}
	}
	return nil, fmt.Errorf("user not found")
}
