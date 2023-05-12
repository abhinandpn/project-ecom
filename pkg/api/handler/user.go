package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"

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

// ----------------------User-----------------
// UserSignUp godoc
// @summary api for user to signup
// @security ApiKeyAuth
// @id UserSignUp
// @tags User Signup
// @Param input body req.ReqUserDetails{} true "Input Fields"
// @Router /signup [post]
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

	copier.Copy(&user, body)

	err = usr.userUseCase.SignUp(ctx, user)
	if err != nil {
		response := res.ErrorResponse(400, "faild to signup", err.Error(), body)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := res.SuccessResponse(200, "Successfully Created Account", body)
	ctx.JSON(200, response)
}

// ----------------------User End-------------
