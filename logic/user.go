package login

import (
	"book/model"
	"book/tools"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func BorrowBook(c *gin.Context) {
	//获取用户信息
	uid := c.GetInt64("uid")
	//获取图书ID
	idStr := c.Query("id")
	if idStr == "" || idStr == "0" {
		c.JSON(http.StatusOK, tools.ParamErr)
		return
	}
	id, _ := strconv.ParseInt(idStr, 10, 64)
	//执行借书逻辑
	err := model.BorrowBook(uid, id)
	if err != nil {
		c.JSON(http.StatusOK, tools.ECode{
			Code:    10002,
			Message: err.Error(),
		})
		return
	}
	//返回成功
	c.JSON(http.StatusOK, tools.OK)
}

func ReturnBook(c *gin.Context) {
	//获取用户信息
	uid := c.GetInt64("uid")
	//获取图书ID
	idStr := c.Query("id")
	if idStr == "" || idStr == "0" {
		c.JSON(http.StatusOK, tools.ParamErr)
		return
	}
	id, _ := strconv.ParseInt(idStr, 10, 64)
	//执行借书逻辑
	err := model.ReturnBook(uid, id)
	if err != nil {
		c.JSON(http.StatusOK, tools.ECode{
			Code:    10002,
			Message: err.Error(),
		})
		return
	}
	//返回成功
	c.JSON(http.StatusOK, tools.OK)
}

func GetUserIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "uindex.tmpl", nil)
}
