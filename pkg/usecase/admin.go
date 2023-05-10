package usecase

import (
	interfaces "github.com/abhinandpn/project-ecom/pkg/repository/interface"
	services "github.com/abhinandpn/project-ecom/pkg/usecase/interfaces"
)

type AdminUseCase struct {
	adminRepo interfaces.AdminRepository
}

func NewAdminUseCase(repo interfaces.AdminRepository) services.AdminUseCase {

	return &AdminUseCase{adminRepo: repo}
}
