package service

import (
	"context"
	v1 "realworld_demo/api/realworld/v1"
)

func (s *RealWorldService) Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginReply, error) {
	// Implement login logic here
	return &v1.LoginReply{
		User: &v1.LoginReply_User{
			Email:    "xxx@qq.com",
			Token:    "xxxeweqqeq",
			Username: "xxx_hsz",
			Bio:      "xxx",
			Image:    "xxx",
		},
	}, nil
}
