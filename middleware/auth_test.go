package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"ecommerce-graphql-go/data"
	"ecommerce-graphql-go/graph/model"
	"ecommerce-graphql-go/util"

	"github.com/stretchr/testify/assert"
)

func TestAuthMiddleware(t *testing.T) {
	testData := data.NewData()

	user := testData.GetUser("1")
	assert.NotNil(t, user)

	// Generate token
	token, err := GenerateToken(user)
	assert.NoError(t, err)

	t.Run("ValidToken", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/graphql", nil)
		req.Header.Set("Authorization", "Bearer "+token)

		rr := httptest.NewRecorder()

		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user := r.Context().Value(util.UserContextKey)
			assert.NotNil(t, user)
			assert.Equal(t, user.(*model.User).ID, "1")
		})

		middleware := AuthMiddleware(handler)
		middleware.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
	})

	t.Run("InvalidToken", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/graphql", nil)
		req.Header.Set("Authorization", "Bearer invalid-token")

		rr := httptest.NewRecorder()

		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			t.Error("Handler should not be called with invalid token")
		})

		middleware := AuthMiddleware(handler)
		middleware.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusUnauthorized, rr.Code)
	})

	t.Run("NoToken", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/graphql", nil)

		rr := httptest.NewRecorder()

		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			t.Error("Handler should not be called without token")
		})

		middleware := AuthMiddleware(handler)
		middleware.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusUnauthorized, rr.Code)
	})
}
