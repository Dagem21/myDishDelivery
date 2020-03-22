package service

import (
	"github.com/dagem21/mydishdelivery/managers"
	"github.com/dagem21/mydishdelivery/entity"
)

type ManagerService struct {
	managerRepo managers.ManagerRepository
}

func NewManagerService(mngRepo managers.ManagerRepository) managers.ManagerService {
	return &ManagerService{managerRepo: mngRepo}
}

func (ms *ManagerService) Managers() ([]entity.Manager, []error) {
	mngs, errs := ms.managerRepo.Managers()
	if len(errs) > 0 {
		return nil, errs
	}
	return mngs, errs
}

func (ms *ManagerService) Manager(id uint) (*entity.Manager, []error) {
	mng, errs := ms.managerRepo.Manager(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return mng, errs
}

func (ms *ManagerService) UpdateManager(manager *entity.Manager) (*entity.Manager, []error) {
	mng, errs := ms.managerRepo.UpdateManager(manager)
	if len(errs) > 0 {
		return nil, errs
	}
	return mng, errs
}

func (ms *ManagerService) DeleteManager(id uint) (*entity.Manager, []error) {
	mng, errs := ms.managerRepo.DeleteManager(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return mng, errs
}

func (ms *ManagerService) StoreManager(manager *entity.Manager) (*entity.Manager, []error) {
	mng, errs := ms.managerRepo.StoreManager(manager)
	if len(errs) > 0 {
		return nil, errs
	}
	return mng, errs
}
