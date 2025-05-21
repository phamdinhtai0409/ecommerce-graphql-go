package loaders

import (
	"context"
	"ecommerce-graphql-go/data"
	"ecommerce-graphql-go/graph/model"
	"ecommerce-graphql-go/util"
	"time"

	"github.com/vikstrous/dataloadgen"
)

type dataReader struct {
	data *data.Data
}

func (d *dataReader) getProducts(ctx context.Context, productIDs []string) ([]*model.Product, []error) {
	products := make([]*model.Product, len(productIDs))
	errors := make([]error, len(productIDs))

	for i, id := range productIDs {
		product := d.data.GetProduct(id)
		if product == nil {
			errors[i] = nil
			continue
		}
		products[i] = &model.Product{
			ID:          product.ID,
			Name:        product.Name,
			Price:       product.Price,
			InStock:     product.InStock,
			Description: product.Description,
			Category:    product.Category,
		}
	}
	return products, errors
}

func (d *dataReader) getUsers(ctx context.Context, userIDs []string) ([]*model.User, []error) {
	users := make([]*model.User, len(userIDs))
	errors := make([]error, len(userIDs))

	for i, id := range userIDs {
		user := d.data.GetUser(id)
		if user == nil {
			errors[i] = nil
			continue
		}
		users[i] = &model.User{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
			Role:  user.Role,
		}
	}
	return users, errors
}

type Loaders struct {
	ProductLoader *dataloadgen.Loader[string, *model.Product]
	UserLoader    *dataloadgen.Loader[string, *model.User]
}

func NewLoaders(data *data.Data) *Loaders {
	dr := &dataReader{data: data}
	return &Loaders{
		ProductLoader: dataloadgen.NewLoader(dr.getProducts, dataloadgen.WithWait(time.Millisecond)),
		UserLoader:    dataloadgen.NewLoader(dr.getUsers, dataloadgen.WithWait(time.Millisecond)),
	}
}

func For(ctx context.Context) *Loaders {
	return ctx.Value(util.LoaderContextKey).(*Loaders)
}

func GetProduct(ctx context.Context, productID string) (*model.Product, error) {
	loaders := For(ctx)
	return loaders.ProductLoader.Load(ctx, productID)
}

func GetProducts(ctx context.Context, productIDs []string) ([]*model.Product, error) {
	loaders := For(ctx)
	return loaders.ProductLoader.LoadAll(ctx, productIDs)
}

func GetUser(ctx context.Context, userID string) (*model.User, error) {
	loaders := For(ctx)
	return loaders.UserLoader.Load(ctx, userID)
}

func GetUsers(ctx context.Context, userIDs []string) ([]*model.User, error) {
	loaders := For(ctx)
	return loaders.UserLoader.LoadAll(ctx, userIDs)
}
