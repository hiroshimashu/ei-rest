package domain

type User struct {
	ID       string `json:"id"`
	Gender   string `json:"gender"`
	Age      int    `json:"age"`
	Area     string `json:"area"`
	Email    string `json:"email"`
	Password string `json:"password"`
	IsPaid   bool   `json:"ispaid"`
}
