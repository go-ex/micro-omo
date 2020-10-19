package controllers

import (
	context "context"
	"github.com/go-ex/micro-omo/runtime/admin.proto"
)

type User struct {
}

func (u User) Login(ctx context.Context, request *admin.LoginRequest, response *admin.LoginResponse) error {
	response.Code = 0
	response.Msg = "test"

	return nil
}
