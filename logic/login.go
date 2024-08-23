package login

import (
	"book/model"
	"book/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Name         string `json:"name" form:"name"`
	Password     string `json:"password" form:"password"`
	CaptchaId    string `json:"captcha_id" form:"captcha_id"`
	CaptchaValue string `json:"captcha_value" form:"captcha_value"`
}

func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.tmpl", nil)
}
func DoLogin(c *gin.Context) {
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

	ret, _ := model.GetUser(user.Name)
	if ret.ID < 1 || ret.Password != tools.EncryptV1(user.Password) {
		c.JSON(http.StatusOK, tools.ECode{
			Code:    10001,
			Message: "帐号密码错误！",
		})
		return
	}

	//生成TOKEN
	token, _ := model.GetJwt(ret.ID, user.Name)
	c.JSON(http.StatusOK, tools.ECode{
		Message: "登录成功",
		Data:    token,
	})
	return
}

func Logout(c *gin.Context) {

}

func GetCaptcha(context *gin.Context) {
	captcha, err := tools.CaptchaGenerate()
	if err != nil {
		context.JSON(http.StatusOK, tools.ECode{
			Code:    10005,
			Message: err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, tools.ECode{
		Data: captcha,
	})
}
