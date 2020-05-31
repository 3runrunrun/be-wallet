/**
* The Repository,
* lays on Interface Adapter-layer
* it responsibles as an interface to data source
**/

package repository

import (
	"github.com/3runrunrun/be-wallet/account/model"
)

// DBUtils is an interface for database utilities functionality
type DBUtils interface {
	Migrate() error
}

// Repository is an interface for repository layer
type Repository interface {
	AddAccount(account model.Account) error
	ShowAccount() ([]model.Account, error)
}
