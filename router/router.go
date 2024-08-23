package router

import (
	_ "book/logic"
	login "book/logic"
	"book/middleware"
	"github.com/gin-gonic/gin"
)

func New() {
	r := gin.Default()
	// 加载网页文件
	r.LoadHTMLGlob("view/*")
	// 设置静态文件目录，包括图片目录
	r.Static("/static", "images")
	user := r.Group("/user")
	user.GET("/login", login.Login)
	user.POST("/login", login.DoLogin)
	user.GET("/captcha", login.GetCaptcha)
	user.GET("/index", login.GetUserIndex)
	user.Use(middleware.CheckUser)
	user.POST("/book/borrow", login.BorrowBook)
	user.POST("/book/return", login.ReturnBook)

	admin := r.Group("/admin")
	admin.POST("/login", login.AdminLogin)
	admin.GET("/logout", login.AdminLogout)

	book := r.Group("/book")
	book.GET("", login.GetBook)
	book.GET("/list", login.GetBooksFromRedisByCursor)
	book.POST("", login.AddBook)
	book.PUT("", login.SaveBook)
	book.DELETE("", login.DelBook)

	if err := r.Run(":8087"); err != nil {
		panic(err)
	}
}
