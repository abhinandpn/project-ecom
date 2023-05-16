package req

// request for New User
type ReqUserDetails struct {
	UserName        string `json:"username"  binding:"required,min=3,max=15"`
	FName           string `json:"f_name"  binding:"required,min=2,max=50"`
	LName           string `json:"l_name"  binding:"required,min=1,max=50"`
	Email           string `json:"email" binding:"required,email"`
	Number          string `json:"phone" binding:"required,min=10,max=10"`
	Password        string `json:"password"  binding:"required,eqfield=ConfirmPassword"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
}

// request Edit User
type ReqEditUser struct {
	UserName        string `json:"username"  binding:"required,min=3,max=15"`
	FirstName       string `json:"f_name"  binding:"required,min=2,max=50"`
	LastName        string `json:"l_name"  binding:"required,min=1,max=50"`
	Email           string `json:"email" binding:"required,email"`
	Number          string `json:"phone" binding:"required,min=10,max=10"`
	Password        string `json:"password"  binding:"omitempty,eqfield=ConfirmPassword"`
	ConfirmPassword string `json:"confirm_password" binding:"omitempty"`
}

// request for Address
type ReqAddress struct {
	Name        string `json:"name" binding:"required,min=2,max=50"`
	PhoneNumber string `json:"phone_number" binding:"required,min=10,max=10"`
	House       string `json:"house" binding:"required"`
	Area        string `json:"area"`
	LandMark    string `json:"land_mark" binding:"required"`
	City        string `json:"city" binding:"required"`
	Pincode     uint   `json:"pincode" binding:"required,max=6"`
	CountryID   uint   `json:"country_id" binding:"required"`

	IsDefault *bool `json:"is_default"`
}

// request Edit address
type ReqEditAddress struct {
	ID          uint   `json:"address_id" binding:"required"`
	Name        string `json:"name" binding:"required,min=2,max=50"`
	PhoneNumber string `json:"phone_number" binding:"required,min=10,max=10"`
	House       string `json:"house" binding:"required"`
	Area        string `json:"area"`
	LandMark    string `json:"land_mark" binding:"required"`
	City        string `json:"city" binding:"required"`
	Pincode     uint   `json:"pincode" binding:"required,max=6"`
	CountryID   uint   `json:"country_id" binding:"required"`

	IsDefault *bool `json:"is_default"`
}

// Request Edit password
type ReqEditPassword struct {
	ID              uint   `json:"address_id" binding:"required"`
	OldPassword     string `json:"old_password" binding:"required"`
	NewPassword     string `json:"password"  binding:"required,eqfield=ConfirmPassword min=6,max=25"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
}
