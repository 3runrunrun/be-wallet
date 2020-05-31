/**
* The Handler,
* lays on Interface Adapter-layer
* it responsibles as an interface to the Web
*
* Handler always depends on Service
* and it is 1:* dependency cardinalities
**/

package account

import (
	"log"
	"net/http"

	"github.com/3runrunrun/be-wallet/account/service"
	"github.com/gin-gonic/gin"
)

// handlerdbu struct
type handlerdbu struct {
	servicedbu service.DBUtils
}

// HandlerDbu interface definition
type HandlerDbu interface {
	InitModule() gin.HandlerFunc
}

// NewHandlerDbu to init handler object
func NewHandlerDbu(dbu service.DBUtils) HandlerDbu {
	return handlerdbu{dbu}
}

// InitModule to create new account table in database (if not exist)
func (h handlerdbu) InitModule() gin.HandlerFunc {

	err := h.servicedbu.CreateTable()

	if err != nil {
		log.Println("account handler.go: ", err)
		return func(c *gin.Context) {
			c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "status": "failed"})
		}
	}

	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "status": "success", "message": "account table created"})
	}
}
