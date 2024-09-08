package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = "secret" // ключ для HS256

// структура для хранения данных пользователя
type User struct {
	UserID            int    `json:"user_id"`
	UserAccountID     int    `json:"user_account_id"`
	UserRole          string `json:"user_role"`
	UserCompositeRole string `json:"user_composite_role"`
}

// Функция для генерации JWT токена
func generateJWT(user User) (string, error) {
	// создаем claims с необходимыми полями
	claims := jwt.MapClaims{
		"user_id":             user.UserID,
		"user_account_id":     user.UserAccountID,
		"user_role":           user.UserRole,
		"user_composite_role": user.UserCompositeRole,
		"exp":                 time.Now().Add(time.Hour * 1).Unix(), // срок действия токена 1 час
	}

	// создаем токен с методом подписи HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// подписываем токен нашим секретным ключом
	return token.SignedString([]byte(secretKey))
}

// Обработчик для /login
func loginHandler(w http.ResponseWriter, r *http.Request) {
	user := User{
		UserID:            123,
		UserAccountID:     456,
		UserRole:          "admin",
		UserCompositeRole: "admin, user",
	}

	// генерируем JWT токен
	token, err := generateJWT(user)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	// отправляем токен в ответе
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(fmt.Sprintf(`{"token": "%s"}`, token)))
}

func main() {
	http.HandleFunc("/login", loginHandler)
	fmt.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
