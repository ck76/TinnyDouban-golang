package v1

import (
	"github.com/gin-gonic/gin"
	"tinnydouban/internal/service"
	"tinnydouban/pkg/app"
	"tinnydouban/pkg/errcode"
)

//这块的逻辑主要是校验及获取入参后，绑定并获取到的 app_key 和 app_secrect 进行数据库查询，
//检查认证信息是否存在，若存在则进行 Token 的生成并返回。

//curl -v -X POST 'http://127.0.0.1:8000/auth?app_key=eddycjy&app_secret=go-programming-tour-book'
func GetAuth(c *gin.Context) {
	param := service.AuthRequest{}
	response := app.NewResponse(c)
	//valid, errs := app.BindAndValid(c, &param)
	//if !valid {
	//	//global.Logger.Errorf("app.BindAndValid errs: %v", errs)
	//	response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
	//	return
	//}

	svc := service.New(c.Request.Context())
	err := svc.CheckAuth(&param)
	if err != nil {
		//global.Logger.Errorf("svc.CheckAuth err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
		return
	}

	token, err := app.GenerateToken(param.AppKey, param.AppSecret)
	if err != nil {
		//global.Logger.Errorf("app.GenerateToken err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		return
	}
	response.ToResponse(gin.H{
		"token": token,
	})
}
