package helper

import (
	"errors"
	"strconv"

	"github.com/abhinandpn/project-ecom/pkg/domain"
	"github.com/gin-gonic/gin"
)

func UserCheck(user, curUser domain.Users) (err error) {

	if curUser.Email == user.Email {
		err = errors.Join(err, errors.New("user alredy exist with this email"))
	}
	if curUser.Number == user.Number {
		err = errors.Join(err, errors.New("user alredy exist with this phone number"))
	}
	if curUser.UserName == user.UserName {
		err = errors.Join(err, errors.New("user alredy exist with this username"))
	}
	return err
}
func StringToUInt(str string) (uint, error) {
	val, err := strconv.Atoi(str)
	// fmt.Println("new ne w enenenenen--------------->", (str))
	return uint(val), err
}

func GetUserId(ctx *gin.Context) uint {

	UserId := ctx.GetString("userId") // string Type
	Id, err := strconv.Atoi(UserId)   // Int type
	if err != nil {
		return 0
	}
	return uint(Id) // current User Id

}
