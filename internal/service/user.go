package service

import (
	"context"
	"fmt"
	v1 "realworld_demo/api/realworld/v1"
	"realworld_demo/internal/biz"

	"github.com/go-kratos/kratos/v2/transport"
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
	fmt.Println("\n\n=== 开始处理注册请求 ===")
	fmt.Printf("请求内容: %+v\n", req)

	// 打印传输信息
	if tr, ok := transport.FromServerContext(ctx); ok {
		fmt.Printf("请求类型: %s\n", tr.Kind())
		fmt.Printf("请求操作: %s\n", tr.Operation())
		fmt.Printf("请求头: %v\n", tr.RequestHeader())
	}

	// 参数验证
	if req == nil {
		fmt.Println("请求参数错误: 请求为空")
		return nil, fmt.Errorf("请求参数错误: 请求为空")
	}

	if req.User == nil {
		fmt.Println("请求参数错误: user为空")
		return nil, fmt.Errorf("请求参数错误: user为空")
	}

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
