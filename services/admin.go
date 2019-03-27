package services

import "apitp/models"

type AdminService struct {
	BaseService
}

func (s *AdminService) AdminGetByName(loginName string) (*models.Admin, error) {
	return models.AdminGetByName(loginName)
}

func (s *AdminService) Update(user *models.Admin) error {
	return user.Update()
}
