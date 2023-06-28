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

	return user, nil
}

func (usr *userDatabase) FindUserByNumber(ctx context.Context, number string) (domain.Users, error) {

	var user domain.Users

	query := `select * from users where number = $1`

	err := usr.DB.Raw(query, number).Scan(&user).Error

	if err != nil {
		return user, fmt.Errorf("faild to find user with number %v", number)

	}

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
	err := usr.DB.Raw(query, username).Scan(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (usr *userDatabase) SaveUser(ctx context.Context, user domain.Users) (UserId uint, err error) {

	query := `insert into users (user_name,f_name,l_name,email,number,password,created_at)
	Values ($1 ,$2 ,$3 ,$4 ,$5 ,$6 ,$7) returning id;`
	var body domain.Users
	createdAt := time.Now()
	err = usr.DB.Raw(query, user.UserName, user.FName, user.LName,
		user.Email, user.Number, user.Password, createdAt).Scan(&body).Error
	if err != nil {

		return 0, fmt.Errorf("faild to save user %v", user.UserName)
	}

	return body.ID, nil
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

	query := `select * from addresses where user_id = $1;`

	err := usr.DB.Raw(query, Uid).Scan(&body).Error

	if err != nil {
		return body, err
	}
	return body, nil
}

// wishlist

func (w *userDatabase) CreateWishList(id uint) error {

	var body domain.WishList
	query := `insert into wish_lists (user_id)values ($1);`

	err := w.DB.Raw(query, id).Scan(&body).Error
	if err != nil {
		return err
	}
	return nil

}

func (w *userDatabase) AddToWishlistItem(uid, pfid uint) error {

	var body domain.WishListItems
	query := `insert into wish_list_items (wish_list_id,product_info_id)values ($1,$2);`

	err := w.DB.Raw(query, uid, pfid).Scan(&body).Error
	if err != nil {
		return err
	}
	return nil

}

func (w *userDatabase) RemoveFromWishListItem(wid, pfid uint) error {

	fmt.Println("Repo Remove Item Wislist ID = ", wid)
	fmt.Println("Repo Remove Item Product Info  ID = ", pfid)

	query := `DELETE FROM wish_list_items
	WHERE wish_list_id = $1 AND product_info_id = $2;`
	err := w.DB.Exec(query, wid, pfid).Error
	if err != nil {
		return err
	}
	return nil

}

func (w *userDatabase) FindWishListItemByWId(id uint) (domain.WishListItems, error) {

	var body domain.WishListItems
	query := `select * from wish_list_items where wish_list_id =$1;`

	err := w.DB.Raw(query, id).Scan(&body).Error
	if err != nil {
		return body, err
	}
	return body, nil

}

func (w *userDatabase) FindWishListByUid(id uint) (domain.WishList, error) {

	var body domain.WishList
	query := `select * from wish_lists where user_id =$1;`

	err := w.DB.Raw(query, id).Scan(&body).Error
	if err != nil {
		return body, err
	}
	return body, nil
}

func (w *userDatabase) FindProductFromWIshListItem(Wid, pfid uint) (bool, error) {

	query := `select exists (select * from wish_list_items where wish_list_id = $1 and product_info_id = $2) AS body`

	var body bool
	err := w.DB.Raw(query, Wid, pfid).Scan(&body).Error
	if err != nil {
		return body, err
	}

	return body, nil
}

func (w *userDatabase) FindWishListItemByWid(id uint) (domain.WishListItems, error) {

	var body domain.WishListItems
	query := `select * from wish_list_items where wish_list_id = $1`

	err := w.DB.Raw(query, id).Scan(&body).Error
	if err != nil {
		return body, err
	}
	return body, nil
}

func (w *userDatabase) ViewWishList(uid uint, pagination req.PageNation) ([]res.ViewWishList, error) {

	var body []res.ViewWishList
	limit := pagination.Count
	offset := (pagination.PageNumber - 1) * limit

	query := `SELECT
				    p.product_name AS "ProductName",
				    pi.price AS "Price",
				    pi.colour AS "Colour",
				    b.brand_name AS "Brand",
				    c.category_name AS "Category"
				FROM
				    wish_lists wl
				    INNER JOIN wish_list_items wli ON wl.id = wli.wish_list_id
				    INNER JOIN product_infos pi ON wli.product_info_id = pi.id
				    INNER JOIN products p ON pi.product_id = p.id
				    LEFT JOIN brands b ON p.brand_id = b.id
				    LEFT JOIN categories c ON p.category_id = c.id
				WHERE
				    wl.user_id = $1
				LIMIT
					$2 OFFSET $3;`

	err := w.DB.Raw(query, uid, limit, offset).Scan(&body).Error
	if err != nil {
		return body, err
	}
	return body, nil
}

func (u *userDatabase) GetAddressByUid(uid uint) (domain.Address, error) {

	var body domain.Address
	query := `select * from addresses where user_id =$1`
	err := u.DB.Raw(query, uid).Scan(&body).Error
	if err != nil {
		return body, err
	}
	return body, nil
}
