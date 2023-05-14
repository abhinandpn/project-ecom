package req

type LoginStruct struct {
	UserName string `json:"username" binding:"required,min=3,max=15"`
	Password string `json:"password" binding:"required,min=6,max=25"`
}
type OtpLogin struct {
	Username    string `json:"username" binding:"required,min=3,max=15"`
	PhoneNumber uint   `json:"phone" binding:"required,min=10,max=10"`
}
type Otpverify struct {
	UserId uint `json:"user_id" binding:"required,numeric"`
	OTP    uint `json:"otp" binding:"required,min=4,max=6"`
}

/*
Page nation for if we have somany user we can show user by group and
eg: 10 users per page
each page have diffrents users
*/
type PageNation struct {
	Count      uint `json:"count"`
	PageNumber uint `json:"page_number"`
}
type BlockStruct struct {
	UserId uint `json:"user_id" binding:"required,numaric"`
}
type ReqPagination struct {
	Count      uint `json:"count"`
	PageNumber uint `json:"page_number"`
}
