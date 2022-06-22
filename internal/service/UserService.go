package service

import (
	"errors"
	"tinnydouban/internal/model"
)

type LoginRequest struct {
	Name     string `form:"name" binding:"required"`
	Password string `form:"password" binding:"required"`
}

//在上述代码中，我们声明了 AuthRequest 结构体用于接口入参的校验，Name 和 Password 都设置为了必填项，
//在 CheckAuth 方法中，我们使用客户端所传入的认证信息作为筛选条件获取数据行，
//以此根据是否取到认证信息 ID 来进行是否存在的判定。
func (svc *Service) Login(param *LoginRequest) error {
	auth, err := svc.dao.Login(param.Name, param.Password)
	if err != nil {
		return err
	}

	if auth.ID > 0 {
		return nil
	}

	return errors.New("auth info does not exist.")
}

func (svc *Service) Register(param *LoginRequest) error {
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
