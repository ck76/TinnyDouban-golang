package model

type User struct {
	*Model
	Name        string `json:"name"`
	Password    string `json:"password"`
	Phone       string `json:"phone"`
	NewPassword string `json:"new_password"`
}

func (a User) TableName() string {
	return "ttmt_user"
}

//上述方法主要是用于服务端在获取客户端所传入的 app_key 和 app_secret 后，
//根据所传入的认证信息进行获取，以此判别是否真的存在这一条数据。
//func (a User) Login(db *gorm.DB) (User, error) {
//	var auth User
//	db = db.Where("name = ? AND password = ?", a.Name, a.Password)
//	err := db.First(&auth).Error
//	if err != nil && err != gorm.ErrRecordNotFound {
//		return auth, err
//	}
//
//	return auth, nil
//}
//
//func (a User) Register(db *gorm.DB) (User, error) {
//	var auth User
//	db = db.Create(&a)
//	db = db.Where("name = ? AND password = ?", a.Name, a.Password)
//	err := db.First(&auth).Error
//	if err != nil && err != gorm.ErrRecordNotFound {
//		return auth, err
//	}
//
//	return auth, nil
//}
//
//func (a User) GetUserList(db *gorm.DB) ([]User, error) {
//	var users []User
//	err := db.Find(&users).Error
//	if err != nil && err != gorm.ErrRecordNotFound {
//		return users, err
//	}
//
//	return users, nil
//}
//
////https://www.cnblogs.com/tortoise512/p/15316816.html
////https://zhuanlan.zhihu.com/p/453744310
//func (a User) UpdatePassword(db *gorm.DB) error {
//	err := db.
//		Model(&User{}).Where("name = ? AND password = ?", a.Name, a.Password).
//		Update("password", a.NewPassword).Error
//	return err
//}
//
//func (a User) GetUserByName(db *gorm.DB)(User,error){
//	//var user User = User{} //查询前是nil吗
//	var user User
//	err := db.Where("name = ?", a.Name).First(&user).Error
//	return user,err
//}
