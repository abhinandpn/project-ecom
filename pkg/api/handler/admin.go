package handler

import (
	"net/http"

	"github.com/abhinandpn/project-ecom/pkg/domain"
	"github.com/abhinandpn/project-ecom/pkg/util/res"
	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
}

// AdminSignUp godoc
// @summary api for admin to login
// @id AdminSignUp
// @tags Admin Login
// @Param input body domain.Admin{} true "inputs"
// @Router /admin/login [post]
// @Success 200 {object} res.Response{} "successfully logged in"
// @Failure 400 {object} res.Response{} "invalid input"
// @Failure 500 {object} res.Response{} "faild to generate jwt token"

// Admin SIgnup function
func (adm *AdminHandler) AdminSignUp(ctx *gin.Context) {

	// Create a model for admin table
	var admin domain.Admin

	if err := ctx.ShouldBindJSON(&admin); err != nil {
		response := res.ErrorResponse(400, "invlaid inputs", err.Error(), admin)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	// err:=adm
}
