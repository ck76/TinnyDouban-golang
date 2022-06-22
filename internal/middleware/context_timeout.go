package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"time"
)

//在上述代码中，我们调用了 context.WithTimeout 方法设置当前 context 的超时时间，并重新赋予给了 gin.Context，这样子在当前请求运行到指定的时间后，在使用了该 context 的运行流程就会针对 context 所提供的超时时间进行处理，并在指定的时间进行取消行为。效果如下：
//
//_, err := ctxhttp.Login(c.Request.Context(), http.DefaultClient, "https://www.google.com/")
//if err != nil {
//    log.Fatalf("ctxhttp.Login err: %v", err)
//}
func ContextTimeout(t time.Duration) func(c *gin.Context) {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), t)
		defer cancel()

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
