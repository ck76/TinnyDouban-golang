package dao

import (
	"tinnydouban/internal/model"
)

func (d *Dao) Login(name, password string) (model.User, error) {
	auth := model.User{Name: name, Password: password}
	return auth.Login(d.engine)
}

func (d *Dao) Register(name, password string) (model.User, error) {
	auth := model.User{Name: name, Password: password}
	return auth.Register(d.engine)
}

func (d *Dao) GetUserList() ([]model.User, error) {
	auth := model.User{}
	return auth.GetUserList(d.engine)
}
