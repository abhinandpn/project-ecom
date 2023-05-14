package repository

import (
	"context"
	"fmt"
	"time"

	domain "github.com/abhinandpn/project-ecom/pkg/domain"
	interfaces "github.com/abhinandpn/project-ecom/pkg/repository/interface"
	"gorm.io/gorm"
)

type userDatabase struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) interfaces.UserRepository {
	return &userDatabase{DB: DB}
}

// .....................................................
// Find user by details
func (usr *userDatabase) FindUser(ctx context.Context, user domain.Users) (domain.Users, error) {
	query := `select * from users where id = ? or email = ? or number = ? or user_name = ? `

	err := usr.DB.Raw(query, user.ID, user.Email, user.Number, user.UserName).Scan(&user).Error
	if err != nil {
		// return user, errors.New("faild to get user")
		return user, err
	}
	return user, nil
}

// Find user By mail id
func (usr *userDatabase) FindUserByEmail(ctx context.Context, email string) (user domain.Users, err error) {

	query := `SELECT * FROM users WHERE email = $1`

	err = usr.DB.Raw(query, email).Scan(&user).Error

	if err != nil {

		return user, fmt.Errorf("faild to find user with email %v", email)

	}

	fmt.Println("\n\ndb user detail", user)

	return user, nil
}

// find user by num
func (usr *userDatabase) FindUserByNumber(ctx context.Context, number uint) (user domain.Users, err error) {
	query := `select * from users where phone = $1`

	err = usr.DB.Raw(query, number).Scan(&user).Error

	if err != nil {
		return user, fmt.Errorf("faild to find user with number %v", number)

	}

	fmt.Println("\n\ndb user detail", user)

	return user, nil
}

// save user
func (usr *userDatabase) SaveUser(ctx context.Context, user domain.Users) (UserId uint, err error) {

	query := `insert into users (user_name,f_name,l_name,email,number,password,created_at)
	Values ($1 ,$2 ,$3 ,$4 ,$5 ,$6 ,$7)`

	createdAt := time.Now()
	err = usr.DB.Raw(query, user.UserName, user.FName, user.LName,
		user.Email, user.Number, user.Password, createdAt).Scan(&user).Error

	if err != nil {

		return 0, fmt.Errorf("faild to save user %s", user.UserName)
	}

	return UserId, nil
}

// .....................................................
