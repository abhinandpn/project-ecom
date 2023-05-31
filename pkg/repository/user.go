package repository

import (
	"context"
	"fmt"
	"time"

	domain "github.com/abhinandpn/project-ecom/pkg/domain"
	interfaces "github.com/abhinandpn/project-ecom/pkg/repository/interface"
	"github.com/abhinandpn/project-ecom/pkg/util/req"
	"github.com/abhinandpn/project-ecom/pkg/util/res"
	"gorm.io/gorm"
)

type userDatabase struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) interfaces.UserRepository {
	return &userDatabase{DB: DB}
}

func (usr *userDatabase) FindUser(ctx context.Context, user domain.Users) (domain.Users, error) {

	query := `select * from users where id = ? or email = ? or number = ? or user_name = ? `

	err := usr.DB.Raw(query, user.ID, user.Email, user.Number, user.UserName).Scan(&user).Error
	if err != nil {
		// return user, errors.New("faild to get user")
		return user, err
	}
	return user, nil
}

func (usr *userDatabase) FindUserByEmail(ctx context.Context, email string) (domain.Users, error) {

	var user domain.Users

	query := `SELECT * FROM users WHERE email = $1`

	err := usr.DB.Raw(query, email).Scan(&user).Error

	if err != nil {

		return user, fmt.Errorf("faild to find user with email %v", email)

	}

	fmt.Println("\n\ndb user detail", user)

	return user, nil
}

func (usr *userDatabase) FindUserByNumber(ctx context.Context, number string) (domain.Users, error) {

	var user domain.Users

	query := `select * from users where number = $1`

	err := usr.DB.Raw(query, number).Scan(&user).Error

	if err != nil {
		return user, fmt.Errorf("faild to find user with number %v", number)

	}

	fmt.Println("\n\ndb user detail", user)

	return user, nil
}

func (usr *userDatabase) FindUserById(ctx context.Context, id uint) (domain.Users, error) {

	var user domain.Users

	query := `select * from users where id = $1;`
	err := usr.DB.Raw(query, id).Scan(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (usr *userDatabase) FindUserByUserName(ctx context.Context, username string) (domain.Users, error) {

	var user domain.Users

	query := `select * from users where user_name = $1;`
	// fmt.Println("--xxxxxx----- > ", username)
	err := usr.DB.Raw(query, username).Scan(&user).Error
	// fmt.Println("--------------------------- >", usr.DB.Raw(query, username))
	if err != nil {
		return user, err
	}

	return user, nil
}

func (usr *userDatabase) SaveUser(ctx context.Context, user domain.Users) (UserId uint, err error) {

	query := `insert into users (user_name,f_name,l_name,email,number,password,created_at)
	Values ($1 ,$2 ,$3 ,$4 ,$5 ,$6 ,$7)`

	createdAt := time.Now()
	err = usr.DB.Raw(query, user.UserName, user.FName, user.LName,
		user.Email, user.Number, user.Password, createdAt).Scan(&user).Error

	if err != nil {

		return 0, fmt.Errorf("faild to save user %v", user.UserName)
	}

	return UserId, nil
}

func (usr *userDatabase) DeleteUser(ctx context.Context, id uint) error {

	// find if user exist or not
	user, err := usr.FindUserById(ctx, id)
	if err != nil {
		return err
	}
	// delete
	UserId := user.ID
	query := `delete from users where id=$1;`
	err = usr.DB.Raw(query, UserId).Error
	if err != nil {
		return err
	}
	// resonse
	return nil
}

func (usr *userDatabase) ListUsers(ctx context.Context, pagination req.PageNation) (res.ProductResponce, error) {

	limit := pagination.Count
	offset := (pagination.PageNumber - 1) * limit

	var res res.ProductResponce

	query := `select * from users order by users.id desc limit $1 offset $2`
	err := usr.DB.Raw(query, limit, offset).Scan(res).Error
	if err != nil {
		return res, err
	}
	return res, nil

}

func (usr *userDatabase) UpdateUser(ctx context.Context, info domain.Users) (domain.Users, error) {

	query := `update users set 
					user_name = $1,
					f_name = $2,
					l_name = $3,
					email = $4,
					number = $5,
				where id = $6`

	var user domain.Users
	err := usr.DB.Raw(query,
		info.UserName,
		info.FName,
		info.LName,
		info.Email,
		info.Number,
		info.ID).Scan(user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

// -----------------AddAddress-----------------

func (usr *userDatabase) AddAddress(ctx context.Context, Uid uint, addres req.ReqAddress) error {

	var body domain.Address

	query := `insert into addresses (user_id,
					house,
					phone_number,
					street,
					city,
					district,
					pincode,
		landmark)values ($1,$2,$3,$4,$5,$6,$7,$8);`
	err := usr.DB.Raw(query,
		Uid,
		addres.House,
		addres.PhoneNumber,
		addres.Street,
		addres.City,
		addres.District,
		addres.Pincode,
		// addres.IsDefault,
		addres.Landmark).Scan(&body).Error

	fmt.Println("address", body)
	if err != nil {
		return err
	}

	return nil
}

// -----------------DeletAddress-----------------

func (usr *userDatabase) UpdateAddress(ctx context.Context, Uid uint, address req.ReqAddress) error {
	var body req.ReqAddress
	query := `update addresses set house = ,
				phone_number= $1,
				street = $2,
				city = $3,
				district = $4,
				pincode = $5,
				landmark = $6 where id = $7;`
	err := usr.DB.Raw(query,
		body.House,
		body.PhoneNumber,
		body.Street,
		body.City,
		body.District,
		body.Pincode,
		body.Landmark, Uid).Scan(&body).Error

	if err != nil {
		return err
	}
	return nil
}

// -----------------ListAllAddress-----------------

func (usr *userDatabase) ListAllAddress(ctx context.Context, Uid uint) ([]res.ResAddress, error) {

	var body []res.ResAddress

	query := `select * from addresses where user_id = ?;`

	err := usr.DB.Raw(query, Uid).Scan(&body).Error

	if err != nil {
		return body, err
	}
	return body, nil
}
