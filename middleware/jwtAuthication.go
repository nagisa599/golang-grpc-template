package middleware

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

func JwtAuthenticateMiddleware(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
	// TODO: ここのJWTからuser_idを取得する処理を実装
	return next(ctx)
}