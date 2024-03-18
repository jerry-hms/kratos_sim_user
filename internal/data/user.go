package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"kratos_sim/app/user/service/internal/biz"
)

type User struct {
	gorm.Model
	Id       int64  `gorm:"column:id" json:"id"`
	Username string `gorm:"column:username;size:20" json:"username"`
	Password string `gorm:"column:password;size:128" json:"password"`
	Nickname string `gorm:"column:nickname;size:20" json:"nickname"`
	Avatar   string `gorm:"column:avatar" json:"avatar"`
	Mobile   string `gorm:"column:mobile;size:11" json:"mobile"`
}

type UserRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) *UserRepo {
	return &UserRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/order")),
	}
}

func (u *UserRepo) CreateUser(ctx context.Context, b *biz.User) (*biz.User, error) {
	user := User{Username: b.Username, Password: b.Password, Nickname: b.Nickname, Avatar: b.Avatar, Mobile: b.Mobile}
	result := u.data.db.WithContext(ctx).Create(&user)
	return &biz.User{
		Id:        user.Id,
		Nickname:  user.Nickname,
		Avatar:    user.Avatar,
		Mobile:    user.Mobile,
		CreatedAt: user.CreatedAt,
	}, result.Error
}

func (u *UserRepo) GetUser(ctx context.Context, id int64) (*biz.User, error) {
	user := User{}
	result := u.data.db.WithContext(ctx).First(&user, id)
	return &biz.User{
		Id:        user.Id,
		Username:  user.Username,
		Password:  user.Password,
		Nickname:  user.Nickname,
		Avatar:    user.Avatar,
		Mobile:    user.Mobile,
		CreatedAt: user.CreatedAt,
	}, result.Error
}

func (u *UserRepo) FindByUsername(ctx context.Context, username string) (*biz.User, error) {
	user := User{}

	result := u.data.db.WithContext(ctx).Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &biz.User{
		Id:        user.Id,
		Username:  user.Username,
		Nickname:  user.Nickname,
		Password:  user.Password,
		Avatar:    user.Avatar,
		CreatedAt: user.CreatedAt,
	}, nil
}
