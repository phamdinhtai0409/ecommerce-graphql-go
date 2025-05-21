package middleware

import (
	"context"
	"ecommerce-graphql-go/data"
	"ecommerce-graphql-go/loaders"
	"ecommerce-graphql-go/util"
	"net/http"
)

func DataLoaderMiddleware(data *data.Data, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		loader := loaders.NewLoaders(data)
		r = r.WithContext(context.WithValue(r.Context(), util.LoaderContextKey, loader))
		next.ServeHTTP(w, r)
	})
}
