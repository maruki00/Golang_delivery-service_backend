package auth_infrastructure_models

import "gorm.io/gorm"

// impolements AuthEntity
type AuthModel struct {
	gorm.Model
	id         int    `json: id`
	email      string `json: email`
	token      string `json: token`
	user_id    int    `json: user_id`
	user_type  string `json: user_type`
	user_level string `json: user_level`
}

func (obj *AuthModel) SetId(id int) {
	obj.id = id
}
func (obj *AuthModel) SetEmail(email string) {
	obj.email = email
}
func (obj *AuthModel) SetToken(token string) {
	obj.token = token
}
func (obj *AuthModel) SetUserId(user_id int) {
	obj.user_id = user_id
}
func (obj *AuthModel) SetUserType(user_type string) {
	obj.user_type = user_type
}
func (obj *AuthModel) SetUserLevel(user_level string) {
	obj.user_level = user_level
}

func (obj *AuthModel) GetId() int {
	return obj.id
}
func (obj *AuthModel) GetEmail() string {
	return obj.email
}
func (obj *AuthModel) GetToken() string {
	return obj.token
}
func (obj *AuthModel) GetUserId() int {
	return obj.user_id
}
func (obj *AuthModel) GetUserType() string {
	return obj.user_type
}
func (obj *AuthModel) GetUserLevel() string {
	return obj.user_level
}
