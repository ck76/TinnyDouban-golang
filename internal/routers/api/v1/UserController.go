package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"tinnydouban/internal/service"
	"tinnydouban/pkg/app"
	"tinnydouban/pkg/errcode"
)

//这块的逻辑主要是校验及获取入参后，绑定并获取到的 app_key 和 app_secrect 进行数据库查询，
//检查认证信息是否存在，若存在则进行 Token 的生成并返回。

/**
纯纯参数校验和转发
*/

//curl -v -X POST 'http://127.0.0.1:8000/auth?app_key=eddycjy&app_secret=go-programming-tour-book'
func Login(c *gin.Context) {
	fmt.Print("login cccccccccccccccccccccccccccccccccccccccc")
	param := service.UserAuthRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		//global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.Login(&param)
	if err != nil {
		//global.Logger.Errorf("svc.CheckAuth err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
		return
	}

	token, err := app.GenerateToken(param.Name, param.Password)
	if err != nil {
		//global.Logger.Errorf("app.GenerateToken err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		return
	}
	response.ToResponse(gin.H{
		"token":    token,
		"name":     param.Name,
		"password": param.Password,
	})
}

//curl -v -X POST 'http://127.0.0.1:8000/auth?app_key=eddycjy&app_secret=go-programming-tour-book'
func Register(c *gin.Context) {
	param := service.UserAuthRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		//global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.Register(&param)
	if err != nil {
		//global.Logger.Errorf("svc.CheckAuth err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
		return
	}

	token, err := app.GenerateToken(param.Name, param.Password)
	if err != nil {
		//global.Logger.Errorf("app.GenerateToken err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		return
	}

	response.ToResponse(gin.H{
		"token":    token,
		"name":     param.Name,
		"password": param.Password,
	})
}

func GetUserList(c *gin.Context) {
	response := app.NewResponse(c)
	svc := service.New(c.Request.Context())
	us, err := svc.GetUserList()
	if err != nil {
		response.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
		return
	}

	response.ToResponse(gin.H{
		"users": us,
	})
}

func UpadtePassword(c *gin.Context) {
	param := service.UserAuthRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		//global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.UpdatePassword(&param)
	if err != nil {
		//global.Logger.Errorf("svc.CheckAuth err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
		return
	}

	response.ToResponse(gin.H{
		"name":     param.Name,
		"password": param.Password,
	})
}

func GetUserByToken(c *gin.Context) {
	param := service.UserAuthRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		//global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	user, err := svc.GetUserByToken(&param)
	if err != nil {
		//global.Logger.Errorf("svc.CheckAuth err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
		return
	}

	response.ToResponse(gin.H{
		"user": user,
	})
}
