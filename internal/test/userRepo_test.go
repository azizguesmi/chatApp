package test

import (
	"backend/internal/model"
	"backend/internal/repo"
	"testing"
	"time"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name       string
		user       model.User
		expectedId int64
	}{
		{
			"test1", model.User{
				UserName:       "test_user",
				Email:          "test3@gmail.com",
				PasswordHashed: "testpassword",
				CreatedAt:      time.Now(),
			}, 5,
		},
		{
			"test2", model.User{
				UserName:       "test_user",
				Email:          "test4@gmail.com",
				PasswordHashed: "testpassword",
				CreatedAt:      time.Now(),
			}, 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id, err := repo.AddUser(tt.user)
			if err != nil {
				t.Errorf("error in test %s error : %v", tt.name, err)
			}
			if id != tt.expectedId {
				t.Errorf("wrong id inserted for test : %s expected %d gets : %d", tt.name, tt.expectedId, id)
			}
		})
	}
}
