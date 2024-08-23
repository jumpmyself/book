package login

import (
	"book/model"
	"book/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AdminLogin(c *gin.Context) {
	var user User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusOK, tools.ECode{
			Code:    10001,
			Message: err.Error(), //这里有风险
		})
	}

	fmt.Printf("user:%+v\n", user)

	if !tools.CaptchaVerify(tools.CaptchaData{
		CaptchaId: user.CaptchaId,
		Data:      user.CaptchaValue,
	}) {
		c.JSON(http.StatusOK, tools.ECode{
			Code:    10002,
			Message: "验证码校验失败！", //这里有风险
		})
		return
	}

	ret, _ := model.GetAdmin(user.Name)
	if ret.ID < 1 || ret.Password != tools.EncryptV1(user.Password) {
		c.JSON(http.StatusOK, tools.ECode{
			Code:    10001,
			Message: "帐号密码错误！",
		})
		return
	}

	_ = model.SetSession(c, user.Name, ret.ID)
	c.JSON(http.StatusOK, tools.ECode{
		Message: "登录成功",
	})
}

func AdminLogout(c *gin.Context) {
	_ = model.FlushSession(c)
	c.JSON(http.StatusUnauthorized, tools.ECode{
		Code:    0,
		Message: "您已退出登录",
	})
}
