package service

import (
	"context"
	"fmt"
	v1 "realworld_demo/api/realworld/v1"
	"realworld_demo/internal/biz"
)

// Login 方法
func (s *RealWorldService) Login(ctx context.Context, req *v1.LoginRequest) (reply *v1.LoginReply, err error) {
	rv, err := s.uc.Login(ctx, req.User.Email, req.User.Password)
	if err != nil {
		return nil, err
	}
	return &v1.LoginReply{
		User: &v1.LoginReply_User{
			Username: rv.Username,
			Token:    rv.Token,
		},
	}, nil
}

// Register 方法
func (s *RealWorldService) Register(ctx context.Context, req *v1.RegisterRequest) (reply *v1.UserReply, err error) {
	fmt.Printf("注册用户: email=%s, username=%s\n", req.User.Email, req.User.Username)
	// 调用业务层注册用户
	u, err := s.uc.Register(ctx, req.User.Username, req.User.Email, req.User.Password)
	if err != nil {
		fmt.Printf("注册失败: %v\n", err)
		return nil, err
	}
	if u == nil {
		fmt.Println("注册失败: 业务层返回的用户对象为空")
		return nil, fmt.Errorf("注册失败: 业务层返回的用户对象为空")
	}

	fmt.Printf("注册成功: email=%s, username=%s, token=%s\n", u.Email, u.Username, u.Token)

	// 构建响应
	reply = &v1.UserReply{
		User: &v1.UserReply_User{
			Email:    u.Email,
			Username: u.Username,
			Token:    u.Token,
			Bio:      u.Bio,
			Image:    u.Image,
		},
	}

	// 打印响应内容
	fmt.Printf("响应内容: %+v\n", reply)

	return reply, nil
}

// GetCurrentUser 方法
func (s *RealWorldService) GetCurrentUser(ctx context.Context, req *v1.GetCurrentRequest) (reply *v1.UserReply, err error) {
	u, err := s.uc.GetCurrentUser(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.UserReply{
		User: &v1.UserReply_User{
			Username: u.Username,
			Image:    u.Image,
			Bio:      u.Bio,
		},
	}, nil
}

// UpdateUser 方法
func (s *RealWorldService) UpdateUser(ctx context.Context, req *v1.UpdateUserRequest) (reply *v1.UserReply, err error) {
	u, err := s.uc.UpdateUser(ctx, &biz.UserUpdate{
		Email:    req.User.GetEmail(),
		Username: req.User.GetUsername(),
		Password: req.User.GetPassword(),
		Bio:      req.User.GetBio(),
		Image:    req.User.GetImage(),
	})
	if err != nil {
		return nil, err
	}
	return &v1.UserReply{
		User: &v1.UserReply_User{
			Username: u.Username,
			Email:    u.Email,
			Image:    u.Image,
			Bio:      u.Bio,
		},
	}, nil
}
