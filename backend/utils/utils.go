package utils

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

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
	layout := "2006-01-02 15:04:05" // Exemplo de layout, ajuste conforme o formato da sua string
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
