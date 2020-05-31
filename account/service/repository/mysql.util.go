/**
* implementations of DBUtils interface
**/

package repository

import (
	"github.com/3runrunrun/be-wallet/account/model"
	"github.com/jinzhu/gorm"
)

type dbuRepository struct {
	db *gorm.DB
}

// NewDbuRepo for creating repo instance
func NewDbuRepo(db *gorm.DB) DBUtils {
	return dbuRepository{db}
}

func (r dbuRepository) Migrate() error {

	isExist := r.db.HasTable(&model.Account{})

	if !isExist {
		r.db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.Account{})
	}

	return nil

}
