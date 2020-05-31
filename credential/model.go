package credential

import (
	_ "github.com/go-sql-driver/mysql" //mysql dialect
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //mysql dialect
)

// Model object
type Model struct {
	DB *gorm.DB
}

// Credential contain accounts' credential information
type Credential struct {
	gorm.Model
	Pin       string `gorm:"type:varchar(6);NOT NULL" json:"pin"`
	IsActive  string `gorm:"type:varchar(1)" json:"is_active"`
	AccountID uint   `gorm:"column:id_account;type:int unsigned;NOT NULL" json:"id_account"`
}
