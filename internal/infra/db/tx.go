package db

import "context"

func RunTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	return fn(ctx)
}
