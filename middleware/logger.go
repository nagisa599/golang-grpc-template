package middleware

import (
	"context"
	"os"

	"log/slog"

	"github.com/99designs/gqlgen/graphql"
)

type LoggerHandler struct {
    slog.Handler
}

func NewLoggerHandler(h slog.Handler) *LoggerHandler {
    return &LoggerHandler{h}
}

func (lh *LoggerHandler) Handle(ctx context.Context, r slog.Record) error {
    // Add user ID from context to log attributes if present
    if userID := ctx.Value("user_id"); userID != nil {
        r.AddAttrs(slog.Attr{Key: "user_id", Value: slog.AnyValue(userID)})
    }
    return lh.Handler.Handle(ctx, r)
}

func LoggerMiddleware(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
    handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
        Level: slog.LevelInfo,
    })

    logger := slog.New(NewLoggerHandler(handler))

    oc := graphql.GetOperationContext(ctx)

    // Extracting and logging the variables passed in the operation
    variables := make(map[string]interface{})
    for key, value := range oc.Variables {
        variables[key] = value
    }

    logger.Info("GraphQL Operation executed",
        slog.String("query", oc.RawQuery),
        slog.Any("variables", variables), // Log the variables
    )
    return next(ctx)
}
