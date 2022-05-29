package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"tinnydouban/global"
	"tinnydouban/pkg/app"
	"tinnydouban/pkg/email"
	"tinnydouban/pkg/errcode"
	"time"
)

////在前文中我们看到 gin 本身已经自带了一个 Recovery 中间件，但是在项目中，
//我们需要针对我们的公司内部情况或生态圈定制 Recovery 中间件，
//确保异常在被正常捕抓之余，要及时的被识别和处理，
//因此自定义一个 Recovery 中间件是非常有必要的，如下：
func Recovery() gin.HandlerFunc {
	defailtMailer := email.NewEmail(&email.SMTPInfo{
		Host:     global.EmailSetting.Host,
		Port:     global.EmailSetting.Port,
		IsSSL:    global.EmailSetting.IsSSL,
		UserName: global.EmailSetting.UserName,
		Password: global.EmailSetting.Password,
		From:     global.EmailSetting.From,
	})
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				//global.Logger.WithCallersFrames().Errorf("panic recover err: %v", err)

				err := defailtMailer.SendMail(
					global.EmailSetting.To,
					fmt.Sprintf("异常抛出，发生时间: %d", time.Now().Unix()),
					fmt.Sprintf("错误信息: %v", err),
				)
				if err != nil {
					//global.Logger.Panicf("mail.SendMail err: %v", err)
				}

				app.NewResponse(c).ToErrorResponse(errcode.ServerError)
				c.Abort()
			}
		}()
		c.Next()
	}
}
