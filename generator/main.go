package main

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"os"

	"github.com/lestrrat-go/jwx/jwk"
)

func main() {
	// Генерация RSA ключа
	privKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Printf("Ошибка при генерации RSA ключа: %v\n", err)
		return
	}

	// Преобразование приватного ключа в JWK
	jwkKey, err := jwk.New(privKey)
	if err != nil {
		fmt.Printf("Ошибка при создании JWK: %v\n", err)
		return
	}

	// Установка некоторых метаданных
	jwkKey.Set(jwk.KeyIDKey, "my-key-id") // Устанавливаем идентификатор ключа
	jwkKey.Set(jwk.AlgorithmKey, "RS256") // Указываем алгоритм

	// Преобразование в JSON формат
	jsonData, err := json.MarshalIndent(map[string]interface{}{
		"keys": []jwk.Key{jwkKey},
	}, "", "  ")
	if err != nil {
		fmt.Printf("Ошибка при сериализации JWK: %v\n", err)
		return
	}

	// Запись в файл jwk.txt
	err = os.WriteFile("jwk2.txt", jsonData, 0644)
	if err != nil {
		fmt.Printf("Ошибка при записи JWK в файл: %v\n", err)
		return
	}

	fmt.Println("JWK успешно сгенерирован и сохранен в файл jwk.txt")
}
