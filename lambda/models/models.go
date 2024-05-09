package models

import "golang.org/x/crypto/bcrypt"

type RegisterUser struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type User struct {
	Username     string `json:"username"`
	PasswordHash string `json:"password"`
}

func NewUser(registerUser RegisterUser) (User, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(registerUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return User{}, err
	}
	return User{
		Username:     registerUser.Username,
		PasswordHash: string(passwordHash),
	}, nil
}

func ValidatePassword(hashedPassword, plainTextPasword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainTextPasword))
	return err == nil
}
