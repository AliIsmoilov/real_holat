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
	Report() repo.ReportI
}

type storage struct {
	infrastructureTypeRepo repo.InfrastructureTypeI
	infrastructureRepo     repo.InfrastructureI
	userRepo               repo.UserI
	verificationRepo       repo.VerificationStorageI
	reportRepo             repo.ReportI
}

func New(db *gorm.DB) StorageI {
	return &storage{
		infrastructureTypeRepo: postgres.NewInfrastructureType(db),
		infrastructureRepo:     postgres.NewInfrastructure(db),
		userRepo:               postgres.NewUser(db),
		verificationRepo:       postgres.NewVerification(db),
		reportRepo:             postgres.NewReport(db),
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

func (s *storage) Report() repo.ReportI {
	return s.reportRepo
}
