package auth_requests

type RegisterRequest struct {
	UserName string `validate:"required" json:"user_name"`
	FullName string `validate:"required" json:"full_name"`
	Address  string `validate:"required" json:"address"`
	Email    string `validate:"required" json:"email"`
	Password string `validate:"required" json:"password"`
}

type LoginRequest struct {
	Login    string `validate:"required" json:"login"`
	Password string `validate:"required" json:"password"`
}

type ResetPasswordRequest struct {
}

type ForgetPasswordRequest struct {
}

type LogoutRequest struct {
}
