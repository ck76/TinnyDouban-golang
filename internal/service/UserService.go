package service

import (
	"errors"
	"tinnydouban/internal/model"
	"tinnydouban/pkg/app"
)

/**
业务逻辑
*/

// https://go-macaron.com/zh-cn/middlewares/binding
//type LoginRequest struct {
//	Name     string `form:"name" binding:"required"`
//	Password string `form:"password" binding:"required"`
//	NewPassword string `form:"new_password"`
//}
//
//type UpdatePasswordRequest struct {
//	Name        string `form:"name" binding:"required"`
//	Password    string `form:"password" binding:"required"`
//	NewPassword string `form:"new_password"`
//}

type UserAuthRequest struct {
	Name        string `form:"name" binding:"required"`
	Password    string `form:"password" binding:"required"`
	NewPassword string `form:"new_password"`
	Token       string `form:"token"`
}

//在上述代码中，我们声明了 AuthRequest 结构体用于接口入参的校验，Name 和 Password 都设置为了必填项，
//在 CheckAuth 方法中，我们使用客户端所传入的认证信息作为筛选条件获取数据行，
//以此根据是否取到认证信息 ID 来进行是否存在的判定。
func (svc *Service) Login(param *UserAuthRequest) error {
	auth, err := svc.dao.Login(param.Name, param.Password)
	if err != nil {
		return err
	}

	if auth.ID > 0 {
		return nil
	}

	return errors.New("auth info does not exist.")
}

func (svc *Service) Register(param *UserAuthRequest) error {
	user, err := svc.dao.GetUserByName(param.Name)
	if user.Name != "" {
		return errors.New("user name already exists")
	}
	auth, err := svc.dao.Register(param.Name, param.Password)
	if err != nil {
		return err
	}

	if auth.ID > 0 {
		return nil
	}

	return errors.New("auth info does not exist.")
}

func (svc *Service) GetUserList() ([]model.User, error) {
	return svc.dao.GetUserList()
}

func (svc *Service) UpdatePassword(param *UserAuthRequest) error {
	//后续做encoder
	//密码规则在前端校验,保证传入数据合法
	user, err := svc.dao.GetUserByName(param.Name)
	//所以说dao层不一定必要
	//model.User{}.GetUserByName(svc.dao.Engine)
	if user.Password != param.Password {
		return errors.New("wrong pass")
	}
	if err != nil {
		return err
	}
	return svc.dao.UpdatePassword(param.Name, param.Password, param.NewPassword)
}

//TODO uuuu
func (svc *Service) GetUserByToken(param *UserAuthRequest) (*model.User, error) {
	claims, err := app.ParseToken(param.Token)
	if err != nil {
		return nil, errors.New("err")
	}
	var user model.User
	//解析出名字，然后查user by name
	//svc.dao.GetUserByName(claims)
	if claims != nil {

	}
	return &user, err
}
