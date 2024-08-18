package shared_entities

type UserEntity interface {
	GetId() int
	GetUsername() string
	GetFullname() string
	GetEmail() string
	GetAddress() string
	GetPassword() string
	GetUserType() string
	GetUserLevel() string
	GetIsOnline() string
	GetIsLocked() string
	GetLastSeen() string
	GetCreatedAt() string
	GetUpdatedAt() string
}
