package loaders

import (
	"context"
	"testing"

	"ecommerce-graphql-go/data"
	"ecommerce-graphql-go/graph/model"

	"github.com/stretchr/testify/assert"
)

func TestLoaders(t *testing.T) {
	testData := data.NewData()

	loaders := NewLoaders(testData)

	t.Run("ProductLoader", func(t *testing.T) {
		ctx := context.Background()

		product, err := loaders.ProductLoader.Load(ctx, "1")
		assert.NoError(t, err)
		assert.NotNil(t, product)
		assert.Equal(t, "1", product.ID)
		assert.Equal(t, "Laptop", product.Name)
		assert.Equal(t, 999.99, product.Price)
		assert.Equal(t, int32(10), product.InStock)
		assert.Equal(t, "Electronics", product.Category)

		product, err = loaders.ProductLoader.Load(ctx, "999")
		assert.NoError(t, err)
		assert.Nil(t, product)
	})

	t.Run("UserLoader", func(t *testing.T) {
		ctx := context.Background()

		user, err := loaders.UserLoader.Load(ctx, "1")
		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, "1", user.ID)
		assert.Equal(t, "Admin", user.Name)
		assert.Equal(t, "admin@example.com", user.Email)
		assert.Equal(t, model.RoleAdmin, user.Role)

		user, err = loaders.UserLoader.Load(ctx, "999")
		assert.NoError(t, err)
		assert.Nil(t, user)
	})
}
