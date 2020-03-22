package managers

import "github.com/dagem21/mydishdelivery/entity"

type ManagerRepository interface {
	Managers() ([]entity.Manager, []error)
	Manager(id uint) (*entity.Manager, []error)
	UpdateManager(manager *entity.Manager) (*entity.Manager, []error)
	DeleteManager(id uint) (*entity.Manager, []error)
	StoreManager(manager *entity.Manager) (*entity.Manager, []error)
}
