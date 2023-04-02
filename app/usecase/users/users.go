package users

import (
	"courses/app/model"
	"courses/app/repository/users"
	"courses/helper/token"
	"time"
)

type UsersUsecase interface {
	Login(payload model.Login) (model.RespLogin, error)
	Register(payload *model.User) error
	Delete(id int) error
	Total() (int, error)
}

type userUsecase struct {
	users users.UsersRepo
}

func NewUsersUsecase(users users.UsersRepo) UsersUsecase {
	return &userUsecase{
		users: users,
	}
}

func (u *userUsecase) Login(payload model.Login) (model.RespLogin, error) {
	res := model.RespLogin{}
	users, err := u.users.Login(payload)
	if err != nil {
		return res, err
	}

	token, err := token.NewToken().GenerateTokenJWT(users)
	if err != nil {
		return res, err
	}

	res.Token = token
	res.Type = "bearer"

	return res, nil
}

func (u *userUsecase) Register(payload *model.User) error {
	payload.Type = model.TypeUserNonAdmin
	payload.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	return u.users.Register(payload)
}

func (u *userUsecase) Delete(id int) error {
	return u.users.Delete(id)
}

func (u *userUsecase) Total() (int, error) {
	return u.users.Total()
}
