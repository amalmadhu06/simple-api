package db

import (
	"github.com/amalmadhu06/simple-api/pkg/config"
	"gorm.io/gorm"
)

func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {
	return &gorm.DB{}, nil
}
