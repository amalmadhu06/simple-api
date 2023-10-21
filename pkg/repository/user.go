package repository

import (
	"context"

	"github.com/amalmadhu06/simple-api/pkg/repository/interfaces"
	"gorm.io/gorm"
)

type userDatabase struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) interfaces.UserRepository {
	return &userDatabase{DB}
}

func (ud *userDatabase) FindAll(ctx context.Context) (string, error) {
	users := "User1: Amal, User2: Madhu"
	return users, nil
}
