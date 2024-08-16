package auth_requests

// type RegisterRequest struct {
// 	Name     string `validate:"required" json:"name"`
// 	Email    string `validate:"required" json:"email"`
// 	Password string `validate:"required" json:"password"`
// }

type LoginRequest struct {
	Login    string // `validate:"required" json:"login"`
	Password string // `validate:"required" json:"password"`
}
