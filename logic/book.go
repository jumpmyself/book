package login

import (
	"book/model"
	"book/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"golang.org/x/net/context"
	"net/http"
	"strconv"
	"time"
)

func GetBook(c *gin.Context) {
	idStr := c.Query("id")
	if idStr == "" || idStr == "0" {
		c.JSON(http.StatusOK, tools.ParamErr)
		return
	}
	id, _ := strconv.ParseInt(idStr, 10, 64)
	ret, err := model.GetBook(id)
	if err != nil {
		c.JSON(http.StatusOK, tools.ECode{
			Code:    10001,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, tools.ECode{
		Data: ret,
	})
	return

}

func GetBooksFromRedisByCursor(c *gin.Context) {
	cursorStr := c.Query("cursor")
	fmt.Printf("cursor:%s\n", cursorStr)
	pageSizeStr := c.DefaultQuery("pageSize", "5")

	cursor, err := strconv.Atoi(cursorStr)
	if err != nil || cursor < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的游标参数"})
		return
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的每页大小参数"})
		return
	}

	// 创建具有超时的 Redis 操作上下文
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 构建 Redis 键名
	key := fmt.Sprintf("books_cursor_%d_size_%d", cursor, pageSize)

	// 检查 Redis 中是否存在数据
	exists, err := model.Rdb.HExists(ctx, "books", key).Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "检查 Redis 哈希表出错"})

	}

	if exists {
		// 数据在 Redis 中已存在，直接返回数据
		data, err := model.Rdb.HGet(ctx, "books", key).Result()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "从 Redis 中获取数据出错"})
			return
		}
		// 解码数据并返回
		var ret []*model.BookInfo
		if err := json.Unmarshal([]byte(data), &ret); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "解码 Redis 数据出错"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "data": ret})
		return
	}

	// 数据不在 Redis 中，从数据库获取
	ret, err := model.GetBooksByCursor(cursor, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "从数据库获取数据出错"})
		return
	}

	// 存储数据到 Redis，设置过期时间（可根据需求调整）
	jsonData, err := json.Marshal(ret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "编码 Redis 数据出错"})
		return
	}
	err = model.Rdb.HSet(ctx, "books", key, string(jsonData)).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "存储数据到 Redis 出错"})
		return
	}
	// 设置过期时间为5分钟
	err = model.Rdb.Expire(ctx, "books", 5*time.Minute).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "设置 Redis 过期时间出错"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "data": ret})
}

func AddBook(c *gin.Context) {
	var data model.Book
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusOK, tools.ParamErr)
		return
	}
	//TODO:增加参数校验
	//id, _ := strconv.ParseInt(idStr, 10, 64)
	err := model.CreateBook(&data)
	if err != nil {
		c.JSON(http.StatusOK, tools.ECode{
			Code:    10001,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, tools.OK)
	return
}

func DelBook(c *gin.Context) {
	idStr := c.Query("id")
	if idStr == "" || idStr == "0" {
		c.JSON(http.StatusOK, tools.ParamErr)
		return
	}
	id, _ := strconv.ParseInt(idStr, 10, 64)
	err := model.DeleteBook(id)
	if err != nil {
		c.JSON(http.StatusOK, tools.ECode{
			Code:    10001,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, tools.OK)
	return
}

func SaveBook(c *gin.Context) {
	var data model.Book
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusOK, tools.ParamErr)
		return
	}
	//TODO:增加参数校验
	err := model.SaveBook(&data)
	if err != nil {
		c.JSON(http.StatusOK, tools.ECode{
			Code:    10001,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, tools.ECode{
		Code: 0,
	})
	return
}
