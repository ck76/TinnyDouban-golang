package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"net/http"
	"time"
	_ "tinnydouban/docs"
	"tinnydouban/global"
	"tinnydouban/internal/middleware"
	"tinnydouban/internal/routers/api/v1"
	"tinnydouban/pkg/limiter"
)

var methodLimiters = limiter.NewMethodLimiter().AddBuckets(limiter.LimiterBucketRule{
	Key:          "/auth",
	FillInterval: time.Second,
	Capacity:     10,
	Quantum:      10,
})

func NewRouter() *gin.Engine {
	engine := gin.New()
	engine.Use(middleware.Cors()) //开启中间件 允许使用跨域请求
	//engine.Use(cors.Default())
	//if global.ServerSetting.RunMode == "debug" {
	//	engine.Use(gin.Logger())
	//	engine.Use(gin.Recovery())
	//} else {
	//	engine.Use(middleware.Recovery())
	//	engine.Use(middleware.AccessLog())
	//}
	//engine.Use(middleware.ContextTimeout(60 * time.Second))
	//engine.Use(middleware.AppInfo()) //?????用不用
	//engine.Use(middleware.RateLimiter(methodLimiters))
	//engine.Use(middleware.Translation())

	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//卧槽？必须放group里才可以cros？？？？？
	root := engine.Group("/")
	//root.Use(middleware.Cors())
	{
		root.POST("login", v1.Login)
		root.GET("user/list", v1.GetUserList)
		root.POST("auth", v1.GetAuth)
		root.POST("register", v1.Register)
		root.POST("password/update", v1.UpadtePassword)
	}
	//访问 $HOST/static 时，应用程序会读取到 blog-service/storage/uploads 下的文件
	engine.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))
	group := engine.Group("/api/v1")
	//group.Use(middleware.Cors())
	{
		group.GET("/user/list", v1.GetUserList)
		group.POST("/login", v1.Login)
		group.POST("/register", v1.Register)
	}
	return engine
}
