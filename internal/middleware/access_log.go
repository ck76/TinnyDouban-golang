package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"tinnydouban/global"
	"tinnydouban/pkg/logger"
	"time"
)

type AccessLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w AccessLogWriter) Write(p []byte) (int, error) {
	if n, err := w.body.Write(p); err != nil {
		return n, err
	}
	return w.ResponseWriter.Write(p)
}

//在 AccessLog 方法中，我们初始化了 AccessLogWriter，
//将其赋予给当前的 Writer 写入流（可理解为替换原有），
//并且通过指定方法得到我们所需的日志属性，最终写入到我们的日志中去，其中涉及到了如下信息：
//
//method：当前的调用方法。
//request：当前的请求参数。
//response：当前的请求结果响应主体。
//status_code：当前的响应结果状态码。
//begin_time/end_time：调用方法的开始时间，调用方法结束的结束时间。
func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyWriter := &AccessLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyWriter

		beginTime := time.Now().Unix()
		c.Next()
		endTime := time.Now().Unix()

		fields := logger.Fields{
			"request":  c.Request.PostForm.Encode(),
			"response": bodyWriter.body.String(),
		}
		global.Logger.WithFields(fields).Infof("access log: method: %s, status_code: %d, begin_time: %d, end_time: %d",
			c.Request.Method,
			bodyWriter.Status(),
			beginTime,
			endTime,
		)
	}
}

////在前文中我们看到 gin 本身已经自带了一个 Recovery 中间件，但是在项目中，
//我们需要针对我们的公司内部情况或生态圈定制 Recovery 中间件，
//确保异常在被正常捕抓之余，要及时的被识别和处理，
//因此自定义一个 Recovery 中间件是非常有必要的，如下：
//func Recovery() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		defer func() {
//			if err := recover(); err != nil {
//				//global.Logger.WithCallersFrames().Errorf("panic recover err: %v", err)
//				app.NewResponse(c).ToErrorResponse(errcode.ServerError)
//				c.Abort()
//			}
//		}()
//		c.Next()
//	}
//}
