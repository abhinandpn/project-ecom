package res

var ResponseMap map[string]string

type AdminLoginRes struct {
	ID       uint   `json:"id"`
	UserName string `json:"username"`
	Email    string `json:"email"`
}
