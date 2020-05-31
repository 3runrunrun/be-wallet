/**
* The Service,
* as known as Use Case
* lays on Application Business Rule-layer
*
* Service always depends on Repository
* and it is 1:1 dependency cardinalities
**/

package service

import (
	"github.com/3runrunrun/be-wallet/account/model"
	"github.com/3runrunrun/be-wallet/account/service/repository"
)

// Service is an interface of service layer
type Service interface {
	AddAccount() (*model.Account, error)
	ShowAccount() ([]model.Account, error)
}

type service struct {
	repository repository.Repository
}

// NewService for making service interface
func NewService(r repository.Repository) Service {
	return service{r}
}

func (s service) ShowAccount() ([]model.Account, error) {
	var accounts []model.Account

	return accounts, nil
}

func (s service) AddAccount() (*model.Account, error) {
	account := &model.Account{AccountNo: "1", Fname: "fathir", Lname: "qisthi"}
	return account, nil
}
