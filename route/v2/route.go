package v2

import codegen "github.com/dappster-io/DappsterOS-UserService/codegen/user_service"

type UserService struct{}

func NewUserService() codegen.ServerInterface {
	return &UserService{}
}
