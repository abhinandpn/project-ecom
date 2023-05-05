package usecase

import (
	interfaces "github.com/abhinandpn/project-ecom/pkg/repository/interface"
)

type adminUseCase struct {
	adminRepo interfaces.AdminRepository
}

func NewAdminusecase(repo interfaces.AdminRepository) interfaces.admin
