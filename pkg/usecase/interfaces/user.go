package interfaces

import (
	"context"
)

type UserUseCase interface {
	FindAll(ctx context.Context) (string, error)
}
