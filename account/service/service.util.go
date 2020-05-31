package service

import "github.com/3runrunrun/be-wallet/account/service/repository"

type servicedbu struct {
	dbutils repository.DBUtils
}

// DBUtils is a declaration for db utils function
type DBUtils interface {
	CreateTable() error
}

// NewServiceDbu for creating service-dbutilities object
func NewServiceDbu(dbu repository.DBUtils) DBUtils {
	return servicedbu{dbu}
}

// CreateTable account
func (dbu servicedbu) CreateTable() error {
	err := dbu.dbutils.Migrate()
	return err
}
