package interfaces

import "github.com/gin-gonic/gin"

type AdminHandler interface {

	// admin
	SudoAdminLogin(ctx *gin.Context) // admin mail,pass read form env (Sudo login)
	AdminLogin(ctx *gin.Context)     // admin login with admin table reading
	AdminHome(ctx *gin.Context)      // admin home (verification)

	// user
	Listuser(ctx *gin.Context)           // list all user
	BlockUser(ctx *gin.Context)          // block user
	FindUserByUserName(ctx *gin.Context) // user find with username
	FindUserWithEmail(ctx *gin.Context)  // user find with email
	FindUserWithNumber(ctx *gin.Context) // user find with number
	FindUserWithId(ctx *gin.Context)     // user find with id
	//
	//
}
