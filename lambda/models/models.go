package models

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

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

func CreateToken(user User) string {
	now := time.Now()
	validUntil := now.Add(time.Hour * 1).Unix()

	claims := jwt.MapClaims{
		"user":    user.Username,
		"expires": validUntil,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims, nil)

	secret := "mySecret"

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		// fmt.Errorf("signingString error %w", err)
		return tokenString
	}

	return tokenString
}
