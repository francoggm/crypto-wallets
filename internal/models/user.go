package models

import (
	"net/mail"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int       `json:"id" db:"id"`
	Username  string    `json:"username" db:"username"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"password,omitempty" db:"password"`
	Role      string    `json:"role" db:"role"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	LastLogin time.Time `json:"last_login" db:"last_login"`
}

func (u *User) HashPassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hash)
	return nil
}

func (u *User) ValidUsername() bool {
	return u.Username != "" && len(u.Username) > 3 && len(u.Username) < 20
}

func (u *User) ValidEmail() bool {
	_, err := mail.ParseAddress(u.Email)
	return err == nil
}

func (u *User) ValidPassword() bool {
	return len(u.Password) > 7 && len(u.Password) < 15
}

func (u *User) IsAdmin() bool {
	return u.Role == "admin"
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *UserLogin) ValidEmail() bool {
	_, err := mail.ParseAddress(u.Email)
	return err == nil
}

func (u *UserLogin) ValidPassword() bool {
	return len(u.Password) > 7 && len(u.Password) < 15
}
