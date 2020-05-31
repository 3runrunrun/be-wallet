package accountlog

import (
	_ "github.com/go-sql-driver/mysql" //mysql dialect
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //mysql dialect
)

// Model object
type Model struct {
	DB *gorm.DB
}

// AccountLog contains changes log of accounts' status
type AccountLog struct {
	gorm.Model
	// AccountStatus should be contained with activated/blocked/closed value
	AccountStatus string `gorm:"type:varchar(1)" json:"account_status"`
	AccountID     uint   `gorm:"column:id_account;type:int unsigned;NOT NULL" json:"id_account"`
}
