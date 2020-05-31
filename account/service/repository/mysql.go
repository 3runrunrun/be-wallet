/**
* implementations of Repository interface
**/

package repository

import (
	"github.com/3runrunrun/be-wallet/account/model"
	"github.com/jinzhu/gorm"
)

type gormRepository struct {
	db *gorm.DB
}

// NewGormRepository for creating new DB repository
func NewGormRepository(db *gorm.DB) Repository {
	return gormRepository{db}
}

// Repository.ShowAccount
func (r gormRepository) ShowAccount() ([]model.Account, error) {
	var account []model.Account

	r.db.Find(&account)

	return account, nil
}

// Repository.AddAccount
func (r gormRepository) AddAccount(account model.Account) error {
	account = model.Account{AccountNo: "1", Fname: "fathir", Lname: "qisthi"}
	return nil
}
