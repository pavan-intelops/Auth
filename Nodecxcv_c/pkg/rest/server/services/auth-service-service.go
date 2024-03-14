package services

import (
	"github.com/pavan-intelops/Auth/nodecxcv_c/pkg/rest/server/daos"
	"github.com/pavan-intelops/Auth/nodecxcv_c/pkg/rest/server/models"
)

type Auth_serviceService struct {
	authServiceDao *daos.Auth_serviceDao
}

func NewAuth_serviceService() (*Auth_serviceService, error) {
	authServiceDao, err := daos.NewAuth_serviceDao()
	if err != nil {
		return nil, err
	}
	return &Auth_serviceService{
		authServiceDao: authServiceDao,
	}, nil
}

func (authServiceService *Auth_serviceService) GetAuth_service(id int64) (*models.Auth_service, error) {
	return authServiceService.authServiceDao.GetAuth_service(id)
}

func (authServiceService *Auth_serviceService) DeleteAuth_service(id int64) error {
	return authServiceService.authServiceDao.DeleteAuth_service(id)
}
