package storage

import (
	"real-holat/storage/postgres"
	"real-holat/storage/repo"

	"gorm.io/gorm"
)

type StorageI interface {
	InfrastructureType() repo.InfrastructureTypeI
	Infrastructure() repo.InfrastructureI
	User() repo.UserI
	Verification() repo.VerificationStorageI
}

type storage struct {
	infrastructureTypeRepo repo.InfrastructureTypeI
	infrastructureRepo     repo.InfrastructureI
	userRepo               repo.UserI
	verificationRepo       repo.VerificationStorageI
}

func New(db *gorm.DB) StorageI {
	return &storage{
		infrastructureTypeRepo: postgres.NewInfrastructureType(db),
		infrastructureRepo:     postgres.NewInfrastructure(db),
		userRepo:               postgres.NewUser(db),
		verificationRepo:       postgres.NewVerification(db),
	}
}

func (s *storage) InfrastructureType() repo.InfrastructureTypeI {
	return s.infrastructureTypeRepo
}

func (s *storage) Infrastructure() repo.InfrastructureI {
	return s.infrastructureRepo
}

func (s *storage) User() repo.UserI {
	return s.userRepo
}

func (s *storage) Verification() repo.VerificationStorageI {
	return s.verificationRepo
}
