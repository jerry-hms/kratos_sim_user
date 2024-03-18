package service

import (
	"context"
	"kratos_sim/api/user/service/v1"
	"kratos_sim/app/user/service/internal/biz"
)

func (u *UserService) CreateUser(ctx context.Context, req *v1.CreateUserReq) (*v1.CreateUserReply, error) {
	user, err := u.uc.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	return &v1.CreateUserReply{
		Id:        user.Id,
		Nickname:  user.Nickname,
		Avatar:    user.Avatar,
		Mobile:    user.Mobile,
		CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}

func (u *UserService) GetUser(ctx context.Context, req *v1.GetUserReq) (*v1.GetUserReply, error) {
	user, err := u.uc.Get(ctx, &biz.User{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return &v1.GetUserReply{
		Id:        user.Id,
		Username:  user.Username,
		Nickname:  user.Nickname,
		Avatar:    user.Avatar,
		Mobile:    user.Mobile,
		CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}

func (u *UserService) GetUserByUsername(ctx context.Context, req *v1.GetUserByUsernameReq) (*v1.GetUserByUsernameReply, error) {
	user, err := u.uc.GetUserByUsername(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	return &v1.GetUserByUsernameReply{
		Id:        user.Id,
		Nickname:  user.Nickname,
		Password:  user.Password,
		Username:  user.Username,
		Avatar:    user.Avatar,
		CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}
