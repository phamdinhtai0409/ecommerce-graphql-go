package graph

import (
	"context"
	"testing"

	"ecommerce-graphql-go/data"
	"ecommerce-graphql-go/graph/model"
	"ecommerce-graphql-go/loaders"
	"ecommerce-graphql-go/util"

	"github.com/stretchr/testify/assert"
)

func TestResolver(t *testing.T) {
	testData := data.NewData()
	resolver := &Resolver{
		Data: testData,
	}

	product := testData.GetProduct("1")
	assert.NotNil(t, product, "Test product should exist")

	user := testData.GetUser("2")
	assert.NotNil(t, user, "Test user should exist")

	ctx := context.WithValue(context.Background(), util.UserContextKey, user)
	ctx = context.WithValue(ctx, util.LoaderContextKey, loaders.NewLoaders(testData))

	t.Run("CreateProduct", func(t *testing.T) {
		input := model.ProductInput{
			Name:        "New Product",
			Price:       15.99,
			InStock:     10,
			Description: stringPtr("New Description"),
			Category:    "New",
		}

		result, err := resolver.Mutation().CreateProduct(ctx, input)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, input.Name, result.Name)
		assert.Equal(t, input.Price, result.Price)
		assert.Equal(t, input.InStock, result.InStock)
		assert.Equal(t, input.Description, result.Description)
		assert.Equal(t, input.Category, result.Category)
	})

	t.Run("UpdateProduct", func(t *testing.T) {
		input := model.ProductInput{
			Name:        "Updated Product",
			Price:       20.99,
			InStock:     15,
			Description: stringPtr("Updated Description"),
			Category:    "Updated",
		}

		result, err := resolver.Mutation().UpdateProduct(ctx, "1", input)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, input.Name, result.Name)
		assert.Equal(t, input.Price, result.Price)
		assert.Equal(t, input.InStock, result.InStock)
		assert.Equal(t, input.Description, result.Description)
		assert.Equal(t, input.Category, result.Category)

		// Update product to synchronize data in the cases below
		product = result
	})

	t.Run("GetProduct", func(t *testing.T) {
		result, err := resolver.Query().Product(ctx, "1")
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, product.ID, result.ID)
		assert.Equal(t, product.Name, result.Name)
		assert.Equal(t, product.Price, result.Price)
		assert.Equal(t, product.InStock, result.InStock)
		assert.Equal(t, product.Description, result.Description)
		assert.Equal(t, product.Category, result.Category)
	})

	t.Run("GetProducts", func(t *testing.T) {
		limit := int32(10)
		offset := int32(0)
		category := "Electronics"

		results, err := resolver.Query().Products(ctx, &limit, &offset, &category)
		assert.NoError(t, err)
		assert.NotNil(t, results)
		assert.GreaterOrEqual(t, len(results), 1)
	})

	t.Run("PlaceOrder", func(t *testing.T) {
		productIds := []string{"1", "2"}
		result, err := resolver.Mutation().PlaceOrder(ctx, productIds)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.NotEmpty(t, result.ID)
		assert.Equal(t, 2, len(result.Products))
		assert.Greater(t, result.Total, 0.0)
		assert.Equal(t, "PENDING", result.Status)
		assert.NotEmpty(t, result.CreatedAt)
		assert.Equal(t, user.ID, result.User.ID)
	})
}

func stringPtr(s string) *string {
	return &s
}
