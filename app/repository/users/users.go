package users

import (
	"courses/app/model"
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

type UsersRepo interface {
	Register(payload *model.User) error
	Login(payload model.Login) (model.User, error)
	Total() (int, error)
	Delete(id int) error
}

type usersRepo struct {
	db *gorm.DB
}

func NewUsersRepo(db *gorm.DB) UsersRepo {
	return &usersRepo{
		db: db,
	}
}

func (u *usersRepo) Register(payload *model.User) error {
	pass := sha256.Sum256([]byte(payload.Password))
	query, err := u.db.DB().Exec("INSERT INTO users(username,password,type,created_at,is_deleted) VALUES (?,?,?,?,?)", payload.Username, fmt.Sprintf("%x", pass[:]), payload.Type, time.Now(), false)
	if err != nil {
		return err
	}

	id, err := query.LastInsertId()
	if err != nil {
		return err
	}

	payload.ID = int(id)

	return nil
}

func (u *usersRepo) Login(payload model.Login) (model.User, error) {
	var resp model.User
	pass := sha256.Sum256([]byte(payload.Password))
	query := u.db.Raw("SELECT id, username, type, created_at FROM users WHERE username = ? AND password = ? AND type = ?", payload.Username, fmt.Sprintf("%x", pass[:]), model.TypeUserAdmin).Scan(&resp)

	if query.Error != nil {
		return resp, query.Error
	}

	return resp, nil
}

func (u *usersRepo) Delete(id int) error {
	query := u.db.Exec("UPDATE users SET is_deleted = ? WHERE id = ? AND type = ?", true, id, model.TypeUserNonAdmin)
	if query.Error != nil {
		return query.Error
	}

	return nil
}

func (u *usersRepo) Total() (int, error) {
	total := 0
	err := u.db.DB().QueryRow("SELECT count(*) FROM users WHERE type = ?", model.TypeUserNonAdmin).Scan(&total)
	if err != nil {
		return total, err
	}

	return total, nil
}
