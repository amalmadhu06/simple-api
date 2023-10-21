package interfaces

import "context"

type UserRepository interface {
	FindAll(ctx context.Context) (string, error)
}
