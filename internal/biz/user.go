package biz

import (
	"context"
	"fmt"
	"realworld_demo/internal/conf"
	auth "realworld_demo/internal/pkg/middleware"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           uint
	Email        string
	Username     string
	Bio          string
	Image        string
	PasswordHash string
}

type UserLogin struct {
	Email    string
	Username string
	Token    string
	Bio      string
	Image    string
}

type UserUpdate struct {
	Email    string
	Username string
	Password string
	Bio      string
	Image    string
}

// hashPassword 使用 bcrypt 对密码进行哈希处理
func hashPassword(pwd string) string {
	b, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func verifyPassword(hashed, input string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(input)); err != nil {
		return false
	}
	return true
}

type UserRepo interface {
	CreateUser(ctx context.Context, user *User) error
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	GetUserByUsername(ctx context.Context, username string) (*User, error)
	GetUserByID(ctx context.Context, id uint) (*User, error)
	UpdateUser(ctx context.Context, user *User) (*User, error)
}

type ProfileRepo interface {
	GetProfile(ctx context.Context, username string) (*Profile, error)
	FollowUser(ctx context.Context, currentUserID uint, followingID uint) error
	UnfollowUser(ctx context.Context, currentUserID uint, followingID uint) error
	GetUserFollowingStatus(ctx context.Context, currentUserID uint, userIDs []uint) (following []bool, err error)
}

type UserUsecase struct {
	ur   UserRepo
	pr   ProfileRepo
	jwtc *conf.JWT

	log *log.Helper
}

type Profile struct {
	ID        uint
	Username  string
	Bio       string
	Image     string
	Following bool
}

func NewUserUsecase(ur UserRepo,
	pr ProfileRepo, logger log.Logger, jwtc *conf.JWT) *UserUsecase {
	return &UserUsecase{ur: ur, pr: pr, jwtc: jwtc, log: log.NewHelper(logger)}
}

func (uc *UserUsecase) generateToken(userID uint) string {
	return auth.GenerateToken(uc.jwtc.Secret, userID)
}

func (uc *UserUsecase) Register(ctx context.Context, username, email, password string) (*UserLogin, error) {
	uc.log.Infof("开始注册用户: email=%s, username=%s", email, username)

	// 检查邮箱是否已存在
	_, err := uc.ur.GetUserByEmail(ctx, email)
	if err == nil {
		uc.log.Errorf("注册失败: 邮箱 %s 已存在", email)
		return nil, errors.New(422, "email", "已被注册")
	} else if !errors.Is(err, errors.NotFound("user", "not found by email")) {
		uc.log.Errorf("查询用户邮箱时发生错误: %v", err)
		return nil, errors.InternalServer("user", "查询用户邮箱时发生错误")
	}
	// 创建用户
	u := &User{
		Email:        email,
		Username:     username,
		PasswordHash: hashPassword(password),
	}

	if err := uc.ur.CreateUser(ctx, u); err != nil {
		uc.log.Errorf("创建用户失败: %v", err)
		return nil, errors.InternalServer("user", fmt.Sprintf("创建用户失败: %v", err))
	}

	// 确保用户ID已被设置
	if u.ID == 0 {
		uc.log.Error("用户ID未被设置，无法生成token")
		return nil, errors.InternalServer("user", "创建用户后ID未设置")
	}

	// 生成token
	token := uc.generateToken(u.ID)
	if token == "" {
		uc.log.Error("生成token失败")
		return nil, errors.InternalServer("user", "生成token失败")
	}

	uc.log.Infof("用户注册成功: id=%d, email=%s, username=%s, token=%s", u.ID, email, username, token)

	// 返回用户信息
	return &UserLogin{
		Email:    email,
		Username: username,
		Token:    token,
		Bio:      u.Bio,
		Image:    u.Image,
	}, nil
}

func (uc *UserUsecase) Login(ctx context.Context, email, password string) (*UserLogin, error) {
	if len(email) == 0 {
		return nil, errors.New(422, "email", "cannot be empty")
	}
	u, err := uc.ur.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if !verifyPassword(u.PasswordHash, password) {
		return nil, errors.Unauthorized("user", "登录失败，账号或密码错误")
	}

	return &UserLogin{
		Email:    u.Email,
		Username: u.Username,
		Bio:      u.Bio,
		Image:    u.Image,
		Token:    uc.generateToken(u.ID),
	}, nil
}

func (uc *UserUsecase) GetCurrentUser(ctx context.Context) (*User, error) {
	cu := auth.FromContext(ctx)
	u, err := uc.ur.GetUserByID(ctx, cu.UserID)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (uc *UserUsecase) UpdateUser(ctx context.Context, uu *UserUpdate) (*UserLogin, error) {
	cu := auth.FromContext(ctx)
	u, err := uc.ur.GetUserByID(ctx, cu.UserID)
	if err != nil {
		return nil, err
	}
	u.Email = uu.Email
	u.Image = uu.Image
	u.PasswordHash = hashPassword(uu.Password)
	u.Bio = uu.Bio
	u, err = uc.ur.UpdateUser(ctx, u)
	if err != nil {
		return nil, err
	}
	return &UserLogin{
		Email:    u.Email,
		Username: u.Username,
		Bio:      u.Bio,
		Image:    u.Image,
		Token:    uc.generateToken(u.ID),
	}, nil
}
