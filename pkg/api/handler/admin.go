package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/abhinandpn/project-ecom/pkg/auth"
	"github.com/abhinandpn/project-ecom/pkg/domain"
	"github.com/abhinandpn/project-ecom/pkg/helper"
	services "github.com/abhinandpn/project-ecom/pkg/usecase/interfaces"
	"github.com/abhinandpn/project-ecom/pkg/util/req"
	"github.com/abhinandpn/project-ecom/pkg/util/res"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type AdminHandler struct {
	AdminUseCase services.AdminUseCase
}

func NewAdminHandler(AdminUseCase services.AdminUseCase) *AdminHandler {
	return &AdminHandler{AdminUseCase: AdminUseCase}
}

/*
	Admin Login with Env file reading
	Super admin function
*/
//----------SUPER ADMIN STARTED----------------
func (adm *AdminHandler) SudoAdminLogin(ctx *gin.Context) {

	var body req.AdminLoginStruct

	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		response := res.ErrorResponse(400, "invalid input message from = env admin login", err.Error(), body)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	// check the fields
	if body.Email == "" && body.UserName == "" {
		err := errors.New("enter email or user_name atleast message from = env admin login")
		response := res.ErrorResponse(400, "invalid input message from = env admin login", err.Error(), body)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	var admin domain.Admin
	copier.Copy(&admin, &body)
	admin, err = adm.AdminUseCase.SudoLogin(ctx, admin)

	if err != nil {
		response := res.ErrorResponse(400, "faild to login ------------", err.Error(), admin)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	tokenString, err := auth.GenerateJWT(admin.ID)
	if err != nil {
		response := res.ErrorResponse(500, "faild to generate jwt token", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	ctx.SetCookie("admin-auth", tokenString["jwtToken"], 60*60, "", "", false, true)

	response := res.SuccessResponse(200, "successfully logged in", nil)
	ctx.JSON(http.StatusOK, response)

}

//----------SUPER ADMIN FINISHED---------------

// AdminLogin godoc
// @summary api for admin to login
// @id AdminLogin
// @tags Admin Login
// @Param input body req.LoginStruct{} true "inputs"
// @Router /admin/login [post]
// @Success 200 {object} res.Response{} "successfully logged in"
// @Failure 400 {object} res.Response{} "invalid input"
// @Failure 500 {object} res.Response{} "faild to generate jwt token"
func (adm *AdminHandler) AdminLogin(ctx *gin.Context) {

	var body req.AdminLoginStruct

	// Admin Binding and error handling
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		response := res.ErrorResponse(400, "invalid input", err.Error(), body)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	// check the fields
	if body.Email == "" && body.UserName == "" {
		err := errors.New("enter email or user_name atleast")
		response := res.ErrorResponse(400, "invalid input", err.Error(), body)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	var admin domain.Admin
	copier.Copy(&admin, &body)
	admin, err = adm.AdminUseCase.Login(ctx, admin)

	if err != nil {
		response := res.ErrorResponse(400, "faild to login", err.Error(), admin)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	tokenString, err := auth.GenerateJWT(admin.ID)
	if err != nil {
		response := res.ErrorResponse(500, "faild to generate jwt token", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	ctx.SetCookie("admin-auth", tokenString["jwtToken"], 60*60, "", "", false, true)

	response := res.SuccessResponse(200, "successfully logged in", nil)
	ctx.JSON(http.StatusOK, response)
}

// Admin Home

// AdminHome godoc
// @summary api admin home
// @id AdminHome
// @tags Admin Home
// @Router /admin [get]
// @Success 200 {object} res.Response{} "successfully logged in"
func (a *AdminHandler) AdminHome(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"StatusCode": 200,
		"message":    "Welcome to Admin Home",
	})
}

// List User

// ListUsers godoc
// @summary api for admin to list users
// @id ListUsers
// @tags Admin User
// @Param page_number query int false "Page Number"
// @Param count query int false "Count Of Order"
// @Router /admin/users [get]
// @Success 200 {object} res.Response{} "successfully got all users"
// @Failure 500 {object} res.Response{} "faild to get all users"
func (adm *AdminHandler) Listuser(ctx *gin.Context) {
	fmt.Println("-------------")
	count, err1 := helper.StringToUInt(ctx.Query("count"))
	pageNumber, err2 := helper.StringToUInt(ctx.Query("page_number"))

	err1 = errors.Join(err1, err2)
	if err1 != nil {
		response := res.ErrorResponse(400, "invalid inputs", err1.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	pagination := req.PageNation{
		PageNumber: pageNumber,
		Count:      count,
	}

	user, err := adm.AdminUseCase.FindAllUser(ctx, pagination)

	if err != nil {
		respone := res.ErrorResponse(500, "faild to get all users", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, respone)
		return
	}

	if len(user) == 0 {
		response := res.SuccessResponse(200, "there is no users to show for this page", nil)
		ctx.JSON(http.StatusOK, response)
		return
	}

	response := res.SuccessResponse(200, "successfully got all users", user)
	ctx.JSON(http.StatusOK, response)

}

// Block User

// BlockUser godoc
// @summary api for admin to block or unblock user
// @id BlockUser
// @tags Admin User
// @Param input body req.BlockStruct{} true "inputs"
// @Router /admin/users/block [patch]
// @Success 200 {object} res.Response{} "Successfully changed user block_status"
// @Failure 400 {object} res.Response{} "invalid input"
func (adm *AdminHandler) BlockUser(ctx *gin.Context) {

	var body req.BlockStruct

	if err := ctx.ShouldBindJSON(&body); err != nil {
		response := res.ErrorResponse(400, "invalid input", err.Error(), body)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	err := adm.AdminUseCase.BlockUser(ctx, body.UserId)
	if err != nil {
		response := res.ErrorResponse(400, "faild to change user block_status", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := res.SuccessResponse(200, "Successfully changed user block_status", body)
	// if successfully blocked or unblock user then response 200
	ctx.JSON(http.StatusOK, response)
}
