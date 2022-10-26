package dao

import (
	"github.com/jinzhu/gorm"
	"tinnydouban/internal/model"
)

//操作数据库
func (db *Dao) Login(name, password string) (model.User, error) {
	var auth model.User
	db.Engine.Where("name = ? AND password = ?", name, password)
	err := db.Engine.First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return auth, err
	}

	return auth, nil
	//auth := model.User{Name: name, Password: password}
	//return auth.Login(d.Engine)
}

func (db *Dao) Register(name, password string) (model.User, error) {
	//auth := model.User{Name: name, Password: password}
	//return auth.Register(d.Engine)
	var auth model.User
	db.Engine.Create(&auth)
	db.Engine.Where("name = ? AND password = ?", name, password)
	err := db.Engine.First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return auth, err
	}

	return auth, nil
}

func (db *Dao) GetUserList() ([]model.User, error) {
	var users []model.User
	err := db.Engine.Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return users, err
	}

	return users, nil
	//auth := model.User{}
	//return auth.GetUserList(db.Engine)
}

func (db *Dao) UpdatePassword(name string, pass string, newPass string) error {
	err := db.Engine.Model(&model.User{}).Where("name = ? AND password = ?", name, pass).Update("password", newPass).Error
	return err
}

func (db *Dao) GetUserByName(name string) (model.User, error) {
	//auth := model.User{}
	//return auth.GetUserByName(d.Engine)
	var user model.User
	err := db.Engine.Where("name = ?", name).First(&user).Error
	return user, err
}
