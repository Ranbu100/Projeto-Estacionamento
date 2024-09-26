package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("qwer1234") //minha chave de autenticaçao

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Função para converter uma string para time.Time (necessita de um formato específico)
func ParseTime(input string) (time.Time, error) {
	layout := "02-01-2006 15:04:05" // Exemplo de layout, ajuste conforme o formato da sua string
	t, err := time.Parse(layout, input)
	if err != nil {
		return time.Time{}, errors.New("Formato de data/hora inválido")
	}
	return t, nil
}

func ParseOptionalTime(input *string) (*time.Time, error) {
	if input == nil || *input == "" {
		return nil, nil
	}
	parsedTime, err := ParseTime(*input) // Utiliza ParseTime para converter
	if err != nil {
		return nil, err
	}
	return &parsedTime, nil
}

// Claims é a estrutura que definirá as informações no JWT
type Claims struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

// Função para gerar um token JWT
func GenerateJWT(email string, role string) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &Claims{
		Email: email,
		Role:  role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
