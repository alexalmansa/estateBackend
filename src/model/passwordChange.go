package model

import (
	"database/sql"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type PasswordChange struct {
	passwordHash    string `json:"-"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
	OldPassword     string `json:"password_old"`
}

func (u *PasswordChange) ChangePassword(conn *sql.DB, userId int) error {
	if len(u.Password) < 4 || len(u.PasswordConfirm) < 4 {
		return fmt.Errorf("password must be at least 4 characters long.")
	}

	if u.Password != u.PasswordConfirm {
		return fmt.Errorf("Passwords do not match.")
	}

	row := conn.QueryRow("SELECT password, email from user_account WHERE id = ?", userId)
	userLookup := User{}
	err := row.Scan(&userLookup.Password, &userLookup.Email)
	if err == sql.ErrNoRows {
		fmt.Println("not found user")
		fmt.Println(userLookup.Email)
		return fmt.Errorf("User not found " + err.Error())
	}

	fmt.Printf("DATABASE: " + userLookup.Email + " SENT pass : " + u.OldPassword)

	err = bcrypt.CompareHashAndPassword([]byte(userLookup.Password), []byte(u.OldPassword))
	if err != nil {
		return fmt.Errorf("Invalid old password")
	}

	pwdHash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("There was an error creating your account.")
	}
	u.passwordHash = string(pwdHash)

	_, err = conn.Exec("UPDATE user_account  SET password = ? WHERE id = ?", u.passwordHash, userId)

	return err
}
