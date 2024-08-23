package tools

import (
	"fmt"
	"github.com/mojocn/base64Captcha"
	"github.com/redis/go-redis/v9"
	"golang.org/x/exp/rand"
	"golang.org/x/net/context"
	"strconv"
	"strings"
	"time"
)

type CaptchaData struct {
	CaptchaId string `json:"captcha_id"`
	Data      string `json:"data"`
}

type captchaStore struct {
	Data string `json:"data" form:"data" binding:"required"`
}
type driverString struct {
	Id            string
	CaptchaType   string
	VerifyValue   string
	DriverString  *base64Captcha.DriverString  //字符串
	DriverChinese *base64Captcha.DriverChinese //中文
	DriverMath    *base64Captcha.DriverMath    //数学
	DriverDigit   *base64Captcha.DriverDigit   //数字
}

// 数字驱动
var digitDriver = base64Captcha.DriverDigit{
	Height:   50,  //生成图片高度
	Width:    150, //生成图片宽度
	Length:   5,   //验证码长度
	MaxSkew:  1,   //文字的倾斜度，越大越倾斜。越不容易看懂
	DotCount: 1,   //背景的点数，越大字体越模糊
}

// Store 使用内存驱动，相关数据会存在内存空间里
var store = base64Captcha.DefaultMemStore

func CaptchaGenerate() (CaptchaData, error) {
	var ret CaptchaData

	//注意，这里直接使用digitDriver 会报错，必须传一个指针，具体原因参考接口实现课程中的内容
	c := base64Captcha.NewCaptcha(&digitDriver, store)
	id, b64s, _, err := c.Generate()
	if err != nil {
		return ret, err
	}

	ret.CaptchaId = id
	ret.Data = b64s
	return ret, nil
}

func CaptchaVerify(data CaptchaData) bool {
	return store.Verify(data.CaptchaId, data.Data, true)
}

// GenerateRandomCode 生成指定长度的随机验证码
func GenerateRandomCode(length int) string {
	rand.Seed(uint64(time.Now().UnixNano()))
	var builder strings.Builder
	for i := 0; i < length; i++ {
		digit := rand.Intn(10) // 生成 0 到 9 的随机数
		builder.WriteString(strconv.Itoa(digit))
	}
	return builder.String()
}

// ValidateVerificationCode 验证用户输入的验证码是否正确
func ValidateVerificationCode(inputCode, toStr string) bool {
	// 连接Redis
	ctx := context.Background()
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "192.168.30.38:6379", // Redis服务器地址和端口
		Password: "",                   // Redis密码，如果没有设置则为空
		DB:       0,                    // Redis数据库索引，默认为0
	})

	// 获取存储的验证码
	key := fmt.Sprintf("verification_code:%s", toStr) // 假设email是接收验证码的邮箱
	storedCode, err := redisClient.Get(ctx, key).Result()
	if err != nil {
		fmt.Printf("无法从Redis获取验证码：%v\n", err)
		return false
	}

	// 验证输入的验证码是否与存储的验证码相符
	return inputCode == storedCode
}
