package model

import "github.com/jinzhu/gorm"

type User struct {
	*Model
	Name     string `json:"name"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

func (a User) TableName() string {
	return "ttmt_user"
}

//上述方法主要是用于服务端在获取客户端所传入的 app_key 和 app_secret 后，
//根据所传入的认证信息进行获取，以此判别是否真的存在这一条数据。
func (a User) Login(db *gorm.DB) (User, error) {
	var auth User
	db = db.Where("name = ? AND password = ?", a.Name, a.Password)
	err := db.First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return auth, err
	}

	return auth, nil
}

func (a User) Register(db *gorm.DB) (User, error) {
	var auth User
	db = db.Create(&a)
	db = db.Where("name = ? AND password = ?", a.Name, a.Password)
	err := db.First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return auth, err
	}

	return auth, nil
}

func (a User) GetUserList(db *gorm.DB) ([]User, error) {
	var users []User
	err := db.Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return users, err
	}

	return users, nil
}
