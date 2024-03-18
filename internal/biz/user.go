package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos_sim/api/user/service/v1"
	"time"
)

type User struct {
	Id        int64
	Username  string
	Password  string
	Nickname  string
	Avatar    string
	Mobile    string
	CreatedAt time.Time
}

type UserRepo interface {
	CreateUser(ctx context.Context, u *User) (*User, error)
	GetUser(ctx context.Context, id int64) (*User, error)
	FindByUsername(ctx context.Context, username string) (*User, error)
}

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usecase/server"))}
}

func (u *UserUseCase) Create(ctx context.Context, req *v1.CreateUserReq) (*User, error) {
	user := &User{
		Username: req.Username,
		Nickname: req.Nickname,
		Password: req.Password,
		Avatar:   req.Avatar,
		Mobile:   req.Mobile,
	}
	return u.repo.CreateUser(ctx, user)
}

func (u *UserUseCase) Get(ctx context.Context, user *User) (*User, error) {
	return u.repo.GetUser(ctx, user.Id)
}

func (u *UserUseCase) GetUserByUsername(ctx context.Context, username string) (*User, error) {
	return u.repo.FindByUsername(ctx, username)
}
