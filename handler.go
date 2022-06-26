package main

import (
	"context"
	"demo/kitex_gen/user"
)

// LoginImpl implements the last service interface defined in the IDL.
type LoginImpl struct{}

// CheckUser implements the LoginImpl interface.
func (s *LoginImpl) CheckUser(ctx context.Context, req *user.LoginRequest) (resp *user.LoginResponse, err error) {
	// TODO: Your code here...
	resp = &user.LoginResponse{Message: "success"}
	err = nil
	return
}
