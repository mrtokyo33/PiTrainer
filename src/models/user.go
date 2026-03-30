package models

import (
	"fmt"
	"time"
)

type User struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewUser(username, password string) (*User, error) {
	user := &User{
		Username:  username,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := user.IsValid()
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) IsValid() error {
	if len(u.Username) < 3 {
		return fmt.Errorf("name must be greated than 2. got: %d", len(u.Username))
	}
	if len(u.Username) > 25 {
		return fmt.Errorf("name must be in max 25 characters. got: %d", len(u.Username))
	}

	return nil
}
