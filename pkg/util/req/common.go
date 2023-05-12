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
