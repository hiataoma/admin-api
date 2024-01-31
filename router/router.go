package router

import (
	"admin-api/api/controller"
	"admin-api/common/config"
	"admin-api/docs"
	"admin-api/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// 初始化路由
func InitRouter() *gin.Engine {
	router := gin.New()
	// 跌机恢复
	router.Use(gin.Recovery())
	router.Use(middleware.Cors())
	router.StaticFS(config.Config.ImageSettings.UploadDir, http.Dir(config.Config.ImageSettings.UploadDir))
	router.Use(middleware.Logger())
	register(router)
	return router
}

func register(router *gin.Engine) {
	docs.SwaggerInfo.BasePath = "" // api/v1
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.GET("/api/captcha", controller.Captcha)
	router.POST("/api/login", controller.Login)
	jwt := router.Group("/api", middleware.AuthMiddleware(), middleware.LogMiddleware())
	{
		jwt.GET("/admin/list", controller.GetSysAdminList)                 // 分页获取用户列表
		jwt.GET("/admin/info", controller.GetSysAdminInfo)                 // 根据id查询用户详情
		jwt.POST("/admin/add", controller.CreateSysAdmin)                  // 新增用户
		jwt.PUT("/admin/update", controller.UpdateSysAdmin)                // 修改用户
		jwt.PUT("/admin/updatePassword", controller.ResetSysAdminPassword) // 重置密码
		jwt.GET("/dept/vo/list", controller.QuerySysDeptVoList)            // 获取部门列表
		jwt.GET("/post/vo/list", controller.QuerySysPostVoList)            // 获取岗位列表
		jwt.GET("/role/vo/list", controller.QuerySysRoleVoList)            // 获取角色列表
		// jwt.POST("/post/add", controller.CreateSysPost)
	}
}
