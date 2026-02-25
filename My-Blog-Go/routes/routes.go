package routes

import (
	"my-blog-go/handlers"
	"my-blog-go/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRoutes 配置所有API路由
func SetupRoutes(r *gin.Engine) {
	// 添加全局中间件
	r.Use(middleware.RequestIDMiddleware()) // 请求 ID 追踪
	r.Use(middleware.CORSMiddleware())      // CORS
	r.Use(middleware.RateLimitMiddleware()) // 全局请求频率限制

	// ==================== 健康检查接口 ====================
	// 不经过限流，用于负载均衡器和监控
	r.GET("/health", handlers.HealthCheck)   // 健康检查
	r.GET("/ready", handlers.ReadinessCheck) // 就绪检查
	r.GET("/system", handlers.SystemInfo)    // 系统信息（仅开发环境）

	// API路由组
	api := r.Group("/api")
	{
		// ==================== 公开接口 ====================
		// 无需登录即可访问

		// 认证相关（添加登录频率限制）
		loginGroup := api.Group("")
		loginGroup.Use(middleware.LoginRateLimitMiddleware())
		{
			loginGroup.POST("/user/login", handlers.FrontLogin)                       // 前台登录
			loginGroup.POST("/user/registration", handlers.Register)                  // 用户注册
			loginGroup.POST("/user/updateForForgetPassword", handlers.ForgetPassword) // 忘记密码
			loginGroup.POST("/admin/user/login", handlers.AdminLogin)                 // 后台登录
			loginGroup.POST("/code", handlers.SendCode)                               // 发送验证码
		}

		// 用户相关（需要可选认证来验证 isAdmin 参数）
		api.POST("/user/list", middleware.OptionalAuthMiddleware(), handlers.ListUsers) // 用户列表（isAdmin参数区分视图）

		// 文章相关（需要可选认证来验证 viewAll/isAdmin 参数）
		api.POST("/article/list", middleware.OptionalAuthMiddleware(), handlers.ListArticles)       // 文章列表（viewAll参数区分）
		api.POST("/article/detail", middleware.OptionalAuthMiddleware(), handlers.GetArticleDetail) // 文章详情（isAdmin参数区分）
		api.POST("/summary", handlers.GetArticleSummary)                                            // 获取文章摘要

		// 分类标签
		api.POST("/webInfo/getSortInfo", handlers.GetSortInfo) // 获取分类信息

		// 网站信息（需要可选认证来验证 isAdmin 参数）
		api.POST("/webInfo/getWebInfo", middleware.OptionalAuthMiddleware(), handlers.GetWebInfo) // 网站信息（isAdmin参数区分）

		// 评论
		api.POST("/comment/listComment", handlers.ClientCommentList) // 评论列表

		// 表白墙/爱链（需要可选认证来验证 isAdmin 参数）
		api.POST("/family/list", middleware.OptionalAuthMiddleware(), handlers.ListFamily) // 表白墙列表（isAdmin参数区分）
		api.POST("/family/getAdminFamily", handlers.GetAdminFamily)                        // 获取管理员表白墙

		// 微言（使用可选认证，用于验证 all 参数）
		api.POST("/weiYan/listWeiYan", middleware.OptionalAuthMiddleware(), handlers.WeiYanList)   // 微言列表
		api.POST("/weiYan/listNews", middleware.OptionalAuthMiddleware(), handlers.WeiYanNewsList) // 文章进展列表

		// 资源（需要可选认证来验证 isAdmin 参数）
		api.POST("/resource/listResource", middleware.OptionalAuthMiddleware(), handlers.ListResource)        // 资源列表（isAdmin参数区分）
		api.POST("/webInfo/listResourcePath", middleware.OptionalAuthMiddleware(), handlers.ListResourcePath) // 资源路径列表（isAdmin参数区分）
		api.POST("/webInfo/getClassifyList", handlers.LoveList)                                               // 分类统计
		api.POST("/webInfo/listCollect", handlers.ListCollect)                                                // 收藏列表

		// 树洞
		api.POST("/admin/treeHole/boss/list", handlers.ListTreeHole) // 树洞列表

		// IP相关
		api.POST("/ip", handlers.GetIP)                  // 获取IP
		api.POST("/ip/location", handlers.GetIPLocation) // 获取IP定位和天气
		api.POST("/list/ip", handlers.ListIP)            // IP统计
		api.POST("/submit", middleware.OptionalAuthMiddleware(), handlers.RecordIP) // 记录IP（可选认证，登录用户自动关联）
		api.POST("/map", handlers.MapData)               // 地图数据
		api.POST("/user/map", handlers.UserMapData)      // 用户地图数据

		// ==================== 需要登录的接口 ====================
		// 仅登录用户可访问
		auth := api.Group("")
		auth.Use(middleware.AuthMiddleware())
		{
			// 用户信息
			auth.POST("/user/updateUserInfo", handlers.UpdateUserInfo) // 更新用户信息

			// 资源上传
			auth.POST("/resource/saveResource", handlers.SaveResource) // 保存资源记录
			auth.POST("/resource/updateImage", handlers.UploadImage)   // 上传图片

			// 评论
			auth.POST("/admin/comment/boss/addComment", handlers.AddComment) // 添加评论

			// 爱链
			auth.POST("/family/addFamily", handlers.AddFamily) // 添加爱链

			// 微言
			auth.POST("/weiYan/deleteWeiYan", handlers.DeleteWeiYan) // 删除微言
			auth.POST("/weiYan/save", handlers.SaveWeiYan)           // 保存微言/文章进展（source参数区分）

			// 资源路径
			auth.POST("/webInfo/saveResourcePath", handlers.SaveResourcePath) // 保存资源路径

			// 树洞
			auth.POST("/webInfo/saveTreeHole", handlers.SaveTreeHole) // 保存树洞

			// 文章点赞
			auth.POST("/article/articleLike", handlers.ArticleLike) // 文章点赞/取消
		}

		// ==================== 管理员接口 ====================
		// 需要管理员权限
		admin := api.Group("")
		admin.Use(middleware.AdminMiddleware())
		{
			// 网站信息管理
			admin.POST("/admin/webInfo/updateAdminWebInfo", handlers.UpdateWebInfo) // 更新网站信息

			// 用户管理
			admin.POST("/admin/user/updateAttribute", handlers.UpdateUserAttribute) // 修改用户属性（status/admire/type）

			// 分类标签管理
			admin.POST("/webInfo/saveOrUpdateSort", handlers.SaveOrUpdateSort)   // 保存或更新分类
			admin.POST("/webInfo/deleteSort", handlers.DeleteSort)               // 删除分类
			admin.POST("/webInfo/saveOrUpdateLabel", handlers.SaveOrUpdateLabel) // 保存或更新标签
			admin.POST("/webInfo/deleteLabel", handlers.DeleteLabel)             // 删除标签
			admin.POST("/webInfo/listSortAndLabel", handlers.ListSortAndLabel)   // 获取分类和标签

			// 文章管理
			admin.POST("/article/delete", handlers.DeleteArticle)             // 删除文章
			admin.POST("/article/save", handlers.SaveArticle)                 // 保存文章
			admin.POST("/article/update", handlers.UpdateArticle)             // 更新文章
			admin.POST("/article/changeStatus", handlers.ChangeArticleStatus) // 修改文章状态
			admin.POST("/delArticleImage", handlers.DeleteImage)              // 删除文章图片

			// 资源管理
			admin.POST("/resource/changeResourceStatus", handlers.ChangeResourceStatus) // 修改资源状态

			// 资源路径管理
			admin.POST("/webInfo/updateResourcePath", handlers.UpdateResourcePath) // 更新资源路径
			admin.POST("/webInfo/deleteResourcePath", handlers.DeleteResourcePath) // 删除资源路径

			// 树洞管理
			admin.POST("/webInfo/deleteTreeHole", handlers.DeleteTreeHole) // 删除树洞

			// 禁用词管理
			admin.POST("/prohibitedWords/list", handlers.ListWords)     // 禁用词列表
			admin.POST("/prohibitedWords/add", handlers.SaveWords)      // 添加禁用词
			admin.POST("/prohibitedWords/delete", handlers.DeleteWords) // 删除禁用词
			admin.POST("/prohibitedWords/update", handlers.UpdateWords) // 更新禁用词

			// 评论管理
			admin.POST("/admin/comment/boss/list", handlers.AdminCommentList)       // 评论列表（管理）
			admin.POST("/admin/comment/boss/deleteComment", handlers.DeleteComment) // 删除评论

			// 爱链管理
			admin.POST("/family/deleteFamily", handlers.DeleteFamily)           // 删除爱链
			admin.POST("/family/changeLoveStatus", handlers.ChangeFamilyStatus) // 修改爱链状态

			// 微言管理
			admin.POST("/weiYan/changePublic", handlers.ChangeWeiYanPublic) // 修改微言公开状态

			// 评论通知（管理员手动发送）
			admin.POST("/codeComment", handlers.SendCodeComment) // 发送评论通知
		}
	}
}
