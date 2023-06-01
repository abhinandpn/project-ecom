package interfaces

import "github.com/gin-gonic/gin"

type AdminHandler interface {
	SudoAdminLogin(ctx *gin.Context)
	AdminLogin(ctx *gin.Context)
	AdminHome(ctx *gin.Context)
	Listuser(ctx *gin.Context)
	BlockUser(ctx *gin.Context)
	FindUserByUserName(ctx *gin.Context)
	FindUserWithEmail(ctx *gin.Context)
	FindUserWithNumber(ctx *gin.Context)
	FindUserWithId(ctx *gin.Context)
}
