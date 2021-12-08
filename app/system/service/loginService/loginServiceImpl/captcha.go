package loginServiceImpl

import (
	"baize/app/system/models/systemModels"
	"image/color"

	"github.com/mojocn/base64Captcha"
)

var store = base64Captcha.DefaultMemStore
var rgba = color.RGBA{0, 0, 0, 0}
var fonts = []string{"wqy-microhei.ttc"}

//生成driver，      高，宽，背景文字的干扰，画线条数，背景颜色的指针，字体
var driver = base64Captcha.NewDriverMath(38, 106, 0, 0, &rgba, fonts)

func GenerateCode() (m *systemModels.CaptchaVo, err error) {
	captcha := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := captcha.Generate()
	if err != nil {
		return m, err
	}
	m = new(systemModels.CaptchaVo)
	m.Id = id
	m.Img = b64s
	return m, nil
}

func VerityCaptcha(id, base64 string) bool {
	return store.Verify(id, base64, true)
}
