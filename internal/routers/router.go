package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"time"
	_ "tinnydouban/docs"
	"tinnydouban/global"
	"tinnydouban/internal/middleware"
	v1 "tinnydouban/internal/routers/api/v1"
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

	if global.ServerSetting.RunMode == "debug" {
		engine.Use(gin.Logger())
		engine.Use(gin.Recovery())
	} else {
		engine.Use(middleware.Recovery())
		engine.Use(middleware.AccessLog())
	}
	engine.Use(middleware.ContextTimeout(60 * time.Second))
	engine.Use(middleware.AppInfo()) //?????用不用
	engine.Use(middleware.RateLimiter(methodLimiters))
	engine.Use(middleware.Translation())

	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	engine.POST("/auth", v1.GetAuth)
	engine.POST("/login", v1.Login)
	engine.POST("/register", v1.Register)

	//访问 $HOST/static 时，应用程序会读取到 blog-service/storage/uploads 下的文件
	engine.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))
	group := engine.Group("/api/v1")
	group.Use(middleware.JWT())
	{
		group.GET("/UserList", v1.GetUserList)
	}
	return engine
}
