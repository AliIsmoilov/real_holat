package storage

import (
	"real-holat/storage/postgres"
	"real-holat/storage/repo"

	"gorm.io/gorm"
)

type StorageI interface {
	InfrastructureType() repo.InfrastructureTypeI
	User() repo.UserI
}

type storage struct {
	infrastructureTypeRepo repo.InfrastructureTypeI
	userRepo               repo.UserI
}

func New(db *gorm.DB) StorageI {
	return &storage{
		infrastructureTypeRepo: postgres.NewInfrastructureType(db),
		userRepo:               postgres.NewUser(db),
	}
}

func (s *storage) InfrastructureType() repo.InfrastructureTypeI {
	return s.infrastructureTypeRepo
}

func (s *storage) User() repo.UserI {
	return s.userRepo
}
