package util

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/oapi-codegen/runtime/types"
)

func ToPtr[T any](v T) *T {
	return &v
}

type contextKey string

const userIDKey contextKey = "userID"

func SetUserIDInContext(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, userIDKey, userID)
}

func GetUserIDFromContext(ctx context.Context) (types.UUID, error) {
	userID, ok := ctx.Value(userIDKey).(types.UUID)
	if !ok {
		return uuid.New(), errors.New("user Id not found in context")
	}
	return userID, nil
}
