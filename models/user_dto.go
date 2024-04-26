package dto

import "rest-api-2/db/dbstore"

type CreateUserDto struct {
	ID       int64  `json:"id"`
	Role     string `json:"role"`
	Username string `json:"username"`
	Password string `json:"password"`
}
type UserDto struct {
	Role     string `json:"role"`
	Username string `json:"username"`
}
type UpdateUserDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type LoginDto struct {
	Role     string `json:"role"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func DtoToUser(user *CreateUserDto) *dbstore.CreateUserParams {
	return &dbstore.CreateUserParams{
		Username: user.Username,
		Role:     user.Role,
		Password: user.Password,
	}
}

func UserToDto(user *dbstore.User) *UserDto {
	return &UserDto{
		Username: user.Username,
		Role:     user.Role,
	}
}
func UsersToDto(user *[]dbstore.User) []UserDto {
	var result []UserDto

	for _, users := range *user {
		userdto := UserToDto(&users)
		result = append(result, *userdto)
	}
	return result
}
func UpdateDtoToUser(user *UpdateUserDto) *dbstore.UpdateUserParams {
	return &dbstore.UpdateUserParams{
		Username: user.Username,
		Password: user.Password,
	}
}
func DtoToLogin(login *LoginDto) *dbstore.LoginRow {
	return &dbstore.LoginRow{
		Username: login.Username,
		Password: login.Password,
	}
}
func LoginToDto(login *dbstore.LoginRow) *LoginDto {
	return &LoginDto{
		Role:     login.Role,
		Password: login.Password,
	}
}
