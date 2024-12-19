package v2

import codegen "github.com/dappsteros-io/DappsterOS-UserService/codegen/user_service"

type UserService struct{}

func NewUserService() codegen.ServerInterface {
	return &UserService{}
}
