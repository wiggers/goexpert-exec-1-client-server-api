package entity

import (
	"context"
)

type RepositoryInterface interface {
	Save(ctx context.Context, exchange *Exchange) error
}
