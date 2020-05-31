/**
* The Model,
* lays on Enterprise Business Model (Entities)-layer
* it describes the most general requirements of the system
* Handler always depends on Service
* and it is 1:* dependency cardinalities
**/

package model

import (
	_ "github.com/go-sql-driver/mysql" //mysql dialect
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //mysql dialect
)

// Model contain gorm DB object
type Model struct {
	DB *gorm.DB
}

// Account table structure
type Account struct {
	gorm.Model
	AccountNo string `gorm:"type:varchar(6);NOT NULL" json:"account_no"`
	Fname     string `gorm:"type:varchar(25);NOT NULL" json:"first_name"`
	Lname     string `gorm:"type:varchar(25);NOT NULL" json:"last_name"`
}
