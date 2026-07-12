package auth

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func GenerateToken(userID int) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenString string) (int, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return 0, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid claims")
	}

	value, exist := claims["user_id"].(float64)
	if !exist {
		return 0, errors.New("not exist")
	}
	id := int(value)
	return id, nil
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 1. Достаем заголовок Authorization из запроса
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "missing authorization header", http.StatusUnauthorized)
			return
		}
		// Заголовок должен быть в формате: "Bearer <токен>"
		// Разрезаем строку по пробелу
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "invalid authorization format", http.StatusUnauthorized)
			return
		}
		tokenString := parts[1] // Вот наш чистый токен

		// 2. Валидируем токен (вызываем твою функцию ValidateToken из задания №20)
		userID, err := ValidateToken(tokenString)
		if err != nil {
			fmt.Println("Ошибка валидации токена:", err)
			http.Error(w, "invalid or expired token", http.StatusUnauthorized)
			return
		}

		// 3. Кладем userID в контекст запроса с помощью нашего ключа UserIDKey
		ctx := context.WithValue(r.Context(), UserIDKey, userID)

		// 4. Нажимаем кнопку "Плей" — передаем управление следующему хендлеру (например, CreateTask)
		// Но передаем уже ОБНОВЛЕННЫЙ запрос r, внутри которого в контексте лежит наш userID
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
