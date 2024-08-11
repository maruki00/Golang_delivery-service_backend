package auth_domain_dto

import (
	"errors"
	"strconv"
)

type AuthDTO struct {
	email      string
	token      string
	user_id    int
	user_type  string
	user_level string
}

func AuthDTOError(msg string) error {
	return errors.New(msg)
}

func AuthDTOFromArray(items map[string]string) (*AuthDTO, error) {
	keys := []string{"email", "token", "user_id", "user_type", "user_level"}
	for _, key := range keys {
		if _, ok := items[key]; !ok {
			return nil, AuthDTOError("key " + key + " is missing")
		}
	}
	user_id, err := strconv.Atoi(items["user_id"])
	if err != nil {
		return nil, errors.New("user_id should be a valid int")
	}
	return &AuthDTO{
		email:      items["email"],
		token:      items["token"],
		user_id:    user_id,
		user_type:  items["user_type"],
		user_level: items["user_level"],
	}, nil
}

func (obj *AuthDTO) GetEmail() string {
	return obj.email
}
func (obj *AuthDTO) GetToken() string {
	return obj.token
}
func (obj *AuthDTO) GetUserId() int {
	return obj.user_id
}
func (obj *AuthDTO) GetUserType() string {
	return obj.user_type
}
func (obj *AuthDTO) GetUserLevel() string {
	return obj.user_level
}
