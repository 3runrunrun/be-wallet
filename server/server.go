package server

import (
	"github.com/3runrunrun/be-wallet/account"
	"github.com/3runrunrun/be-wallet/account/service"
	"github.com/3runrunrun/be-wallet/account/service/repository"
	"github.com/3runrunrun/be-wallet/database"
	"github.com/gin-gonic/gin"
)

var (
	accountDbuHandler account.HandlerDbu
)

// Server to run web server
func Server() *gin.Engine {
	server := gin.Default()

	// api route group
	account := server.Group("/api/account")
	{
		account.GET("/init", accountDbuHandler.InitModule())
	}

	return server
}

// set db connection
// put db_connection as a parameter for creating repository instance
// then, put repo_instance as a param for creating service instance
// then, put service_instance as a param for creating handler instance
func init() {
	sql := database.ConnectSQL()

	accountDbuRepo := repository.NewDbuRepo(sql)
	accountDbuService := service.NewServiceDbu(accountDbuRepo)
	accountDbuHandler = account.NewHandlerDbu(accountDbuService)
}
