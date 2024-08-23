package middleware

import (
	"book/model"
	"book/tools"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheckAdmin() {

}

func CheckUser(c *gin.Context) {
	var name string
	var id int64
	jwt := c.GetString("auth")
	d, err := model.CheckJwt(jwt)
	if err != nil {
		c.JSON(http.StatusOK, tools.ECode{
			Code:    10003,
			Message: "token校验失败",
		})
		c.Abort()
	}

	name = d.Name
	id = d.Id

	if id <= 0 || name == "" {
		c.JSON(http.StatusOK, tools.ECode{
			Code:    10003,
			Message: "用户信息错误",
		})
		c.Abort()
	}

	if model.GetJWTMap(name) {
		c.JSON(http.StatusOK, tools.ECode{
			Code:    10003,
			Message: "用户状态异常，请重新登录！",
		})
		c.Abort()
	}

	c.Next()
}
