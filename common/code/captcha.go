// Package captcha 处理图⽚验证码逻辑
package code

import (
	"github.com/mojocn/base64Captcha"
	"sync"
)

type Captcha struct {
	Base64Captcha *base64Captcha.Captcha
}

// once 确保 internalCaptcha 对象只初始化⼀次
var once sync.Once

// internalCaptcha 内部使⽤的 Captcha 对象
var internalCaptcha *Captcha

// NewCaptcha 单例模式获取
func NewCaptcha() *Captcha {
	once.Do(func() {
		// 初始化 Captcha 对象
		internalCaptcha = &Captcha{}
		// 使⽤全局 Redis 对象，并配置存储 Key 的前缀
		store := RedisStore{
			RedisClient: redisClient,
			KeyPrefix:   "fresh" + ":captcha:",
		}

		// 配置 base64Captcha 驱动信息
		driver := base64Captcha.NewDriverDigit(
			conf.Height,
			conf.Width,
			conf.Length,   // ⻓度
			conf.Maxskew,  // 数字的最⼤倾斜⻆度
			conf.Dotcount, // 图⽚背景⾥的混淆点数量
		)
		// 实例化 base64Captcha 并赋值给内部使⽤的 internalCaptcha 对象
		internalCaptcha.Base64Captcha = base64Captcha.NewCaptcha(driver, &store)
	})
	return internalCaptcha
}

// GenerateCaptcha ⽣成图⽚验证码
func (c *Captcha) GenerateCaptcha() (id string, b64s string, err error) {
	return c.Base64Captcha.Generate()
}

// VerifyCaptcha 验证验证码是否正确
func (c *Captcha) VerifyCaptcha(id string, answer string) (match bool) {
	// ⽅便本地和 API ⾃动测试
	//if !app.IsProduction() && id == config.GetString("captcha.testing_key") {
	//	return true
	//}
	// 第三个参数是验证后是否删除，我们选择 false
	// 这样⽅便⽤户多次提交，防⽌表单提交错误需要多次输⼊图⽚验证码
	return c.Base64Captcha.Verify(id, answer, false)
}
