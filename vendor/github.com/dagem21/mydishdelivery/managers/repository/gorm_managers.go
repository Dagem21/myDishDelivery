package repository

import (
	"github.com/dagem21/mydishdelivery/managers"
	"github.com/dagem21/mydishdelivery/entity"
	"github.com/jinzhu/gorm"
)

type ManagerGormRepo struct {
	conn *gorm.DB
}

func NewManagerGormRepo(db *gorm.DB) managers.ManagerRepository {
	return &ManagerGormRepo{conn: db}
}

func (mngRepo *ManagerGormRepo) Managers() ([]entity.Manager, []error) {
	mngs := []entity.Manager{}
	errs := mngRepo.conn.Find(&mngs).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return mngs, errs
}

func (mngRepo *ManagerGormRepo) Manager(id uint) (*entity.Manager, []error) {
	mng := entity.Manager{}
	errs := mngRepo.conn.First(&mng, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &mng, errs
}

func (mngRepo *ManagerGormRepo) UpdateManager(manager *entity.Manager) (*entity.Manager, []error) {
	mng := manager
	errs := mngRepo.conn.Save(mng).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return mng, errs
}

func (mngRepo *ManagerGormRepo) DeleteManager(id uint) (*entity.Manager, []error) {
	mng, errs := mngRepo.Manager(id)

	if len(errs) > 0 {
		return nil, errs
	}

	errs = mngRepo.conn.Delete(mng, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return mng, errs
}

func (mngRepo *ManagerGormRepo) StoreManager(manager *entity.Manager) (*entity.Manager, []error) {
	mng := manager
	errs := mngRepo.conn.Create(mng).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return mng, errs
}
