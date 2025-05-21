package util

import (
	"context"

	"ecommerce-graphql-go/graph/model"
)

type contextKey string

const (
	UserContextKey   contextKey = "user"
	LoaderContextKey contextKey = "dataloaders"
)

func GetUserFromContext(ctx context.Context) *model.User {
	return ctx.Value(UserContextKey).(*model.User)
}
