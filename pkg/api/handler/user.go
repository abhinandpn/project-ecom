package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"

	twillio "github.com/abhinandpn/project-ecom/pkg/Twillio"
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

// UserLogin godoc

func (usr *UserHandler) UserLogin(ctx *gin.Context) {

	fmt.Println("-------login function started")

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

	ctx.SetCookie("user-auth", tokenString["jwtToken"], 60*60, "", "", false, true)

	response := res.SuccessResponse(200, "successfully logged in", tokenString["jwtToken"])
	ctx.JSON(http.StatusOK, response)

}

// Otp login
func (usr *UserHandler) UserOtpLogin(ctx *gin.Context) {
	fmt.Println(".......................")
	var body req.OTPLoginStruct
	if err := ctx.ShouldBindJSON(&body); err != nil {
		response := res.ErrorResponse(400, "invalid input", err.Error(), body)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	//check all input field is empty
	if body.Email == "" && body.Phone == "" && body.UserName == "" {
		err := errors.New("enter atleast user_name or email or phone")
		response := res.ErrorResponse(400, "invalid input", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	var user domain.Users
	copier.Copy(&user, body)

	user, err := usr.userUseCase.OtpLogin(ctx, user)

	if err != nil {
		resopnse := res.ErrorResponse(400, "can't login", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, resopnse)
		return
	}

	// If we get noe error then sent the OTP
	_, err = twillio.TwillioOtpSent("+91" + user.Number)

	if err != nil {
		response := res.ErrorResponse(500, "faild to send otp", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}
	response := res.SuccessResponse(200, "successfully otp send to registered number", user.ID)
	ctx.JSON(http.StatusOK, response)
}

func (usr *UserHandler) UserLoginOtpVerify(ctx *gin.Context) {

	var body req.OTPVerifyStruct

	if err := ctx.ShouldBindJSON(&body); err != nil {
		response := res.ErrorResponse(400, "invalid login_otp", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	var user = domain.Users{
		ID: body.UserID,
	}

	// get the user using loginOtp useCase

	user, err := usr.userUseCase.OtpLogin(ctx, user)

	if err != nil {
		response := res.ErrorResponse(400, "faild to login", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	// then varify the otp
	err = twillio.TwilioVerifyOTP("+91"+user.Number, body.OTP)
	if err != nil {
		response := res.ErrorResponse(400, "faild to login", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	// if everyting ok then generate token
	tokenString, err := auth.GenerateJWT(user.ID)
	if err != nil {
		response := res.ErrorResponse(500, "faild to login", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	ctx.SetCookie("user-auth", tokenString["accessToken"], 50*60, "", "", false, true)
	response := res.SuccessResponse(200, "successfully logged in uing otp", tokenString["accessToken"])
	ctx.JSON(http.StatusOK, response)
}

func (a *UserHandler) UserHome(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"StatusCode": 200,
		"message":    "Welcome to User Home",
	})
}
