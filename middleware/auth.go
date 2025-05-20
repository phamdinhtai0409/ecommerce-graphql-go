package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"ecommerce-graphql-go/graph/model"
	"ecommerce-graphql-go/util"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	ID    string     `json:"id"`
	Name  string     `json:"name"`
	Email string     `json:"email"`
	Role  model.Role `json:"role"`
	jwt.RegisteredClaims
}

func AuthMiddleware(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 {
			http.Error(w, "Invalid token format", http.StatusUnauthorized)
			return
		}

		tokenString := bearerToken[1]
		claims := &Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(util.GetJWTSecret()), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		user := &model.User{
			ID:    claims.ID,
			Name:  claims.Name,
			Email: claims.Email,
			Role:  claims.Role,
		}
		ctx := context.WithValue(r.Context(), util.UserContextKey, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GenerateToken(user *model.User) (string, error) {
	if user == nil {
		return "", fmt.Errorf("user is null")
	}

	expiration, err := time.ParseDuration(util.GetJWTExpiration())
	if err != nil {
		return "", fmt.Errorf("parsing jwt expires error")
	}

	claims := &Claims{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(util.GetJWTSecret()))
}
