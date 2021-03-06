package routers

import (
	"cblog/controllers/v1"
	_ "cblog/docs"
	"cblog/middleware"
	"cblog/pkg/setting"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"net/http"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	// swagger 文档  "http://ip:port/swagger/index.html"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//r.GET("/swagger/*any", ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "NAME_OF_ENV_VARIABLE"))

	// 全局中间件
	// Logger 中间件将日志写入 gin.DefaultWriter，即使你将 GIN_MODE 设置为 release。
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())
	// Recovery 中间件会 recover 任何 panic。如果有 panic 的话，会写入 500。
	r.Use(gin.Recovery())
	// 日志记录
	r.Use(middleware.LoggerToFile())

	// 静态资源路由
	r.StaticFS("/static/uploads", http.Dir(setting.UploadSetting.Path))

	// 动态路由
	apiv1 := r.Group("/api/v1")
	//apiv1.Use(gin.ErrorLogger())
	{
		apiv1.POST("/register", v1.Register)
		apiv1.POST("/login", v1.Login)
		apiv1.GET("/logout", v1.Logout)
		apiv1.POST("/upload", v1.UploadFile)

		apiv1Authorized := apiv1.Group("")
		apiv1Authorized.Use(middleware.JwtAuth())
		{
			apiv1Authorized.GET("/users/:id", v1.GetUser)
		}

		// 文章
		apiv1.GET("/articles/", v1.GetArticles)
		apiv1.GET("/articles/:id", v1.GetArticle)
		apiv1.POST("/articles", v1.CreateArticle)
		apiv1.PUT("/articles/:id", v1.UpdateArticle)
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
		// 文章标签
		apiv1.GET("/tags", v1.GetTags)
		apiv1.GET("/tags/:id", v1.GetTag)
		apiv1.POST("/tags", v1.CreateTag)
		apiv1.PUT("/tags/:id", v1.UpdateTag)
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
	}

	return r
}
