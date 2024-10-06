package database

import (
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// Функция для генерации случайной строки фиксированной длины
func generateRandomCode(length int) string {
	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[r.Intn(len(charset))]
	}
	return string(result)
}
