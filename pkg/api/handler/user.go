package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"

	"github.com/abhinandpn/project-ecom/pkg/auth"
	"github.com/abhinandpn/project-ecom/pkg/domain"
	services "github.com/abhinandpn/project-ecom/pkg/usecase/interfaces"
	"github.com/abhinandpn/project-ecom/pkg/util/req"
	"github.com/abhinandpn/project-ecom/pkg/util/res"
)

type UserHandler struct {
	userUseCase services.UserUseCase
}

func NewUserHandler(usecase services.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: usecase,
	}
}

// ---------------------- User Function -----------------
// UserSignUp godoc
// @summary api for user to signup
// @security ApiKeyAuth
// @id UserSignUp
// @tags User Signup
// @Param input body req.ReqUserDetails{} true "Input Fields"
// @Router user/signup [post]
// @Success 200 "Successfully created account for user"
// @Failure 400 "invalid input"

func (usr *UserHandler) UserSignUp(ctx *gin.Context) {

	var body req.ReqUserDetails

	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		response := res.ErrorResponse(400, "invalid input", err.Error(), body)

		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	var user domain.Users

	copier.Copy(&user, &body)

	err = usr.userUseCase.SignUp(ctx, user)
	if err != nil {
		response := res.ErrorResponse(400, "faild to signup", err.Error(), body)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := res.SuccessResponse(200, "Successfully Created Account", body)
	ctx.JSON(200, response)
}
func (usr *UserHandler) UserLogin(ctx *gin.Context) {

	var body req.LoginStruct

	err := ctx.ShouldBindJSON(&body)

	if err != nil {

		response := res.ErrorResponse(400, "invalid input", err.Error(), nil)

		ctx.JSON(http.StatusBadRequest, response)

		return
	}

	// Check all input filed is empty

	if body.UserName == "" {

		err := errors.New("enter username")

		response := res.ErrorResponse(400, "invalid input", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	var user domain.Users
	copier.Copy(&user, &body)

	user, err = usr.userUseCase.Login(ctx, user)
	if err != nil {
		response := res.ErrorResponse(400, "faild to login", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	// generate JWT token

	tokenString, err := auth.GenerateJWT(user.ID)
	if err != nil {
		response := res.ErrorResponse(400, "faild to create JWT token", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	ctx.SetCookie("user-auth", tokenString["JWT_CODE"], 60*60, "", "", false, true)

	response := res.SuccessResponse(200, "successfully logged in", tokenString["JWT_CODE"])
	ctx.JSON(http.StatusOK, response)
}
