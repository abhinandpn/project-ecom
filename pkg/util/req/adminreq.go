package req

type AdminLoginStruct struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
