package req

type AdminLoginStruct struct {
	UserName string `json:"id"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
