package imgyzm

import (
	"cmx/xcm"
	"encoding/base64"
	"fmt"
	"image/color"
	"io"
)

const (
	//CaptchaComplexLower complex level lower.
	captchaComplexLower = iota
	//CaptchaComplexMedium complex level medium.
	captchaComplexMedium
	//CaptchaComplexHigh complex level high.
	captchaComplexHigh
)

// CaptchaInterface captcha interface for captcha engine to to write staff
type captchaInterface interface {
	// BinaryEncoding covert to bytes
	BinaryEncoding() (bstrs []byte, err error)
	// WriteTo output captcha entity
	WriteTo(w io.Writer) (n int64, err error)
}

/*GetBase64ImgYzm 获取图片验证码和该验证码的Base64图片格式字符串
len:字符串长度
width:图片宽
height:图片高
yzm:验证码字符串，数字和大写的26个英文字母
base64ImgStr:Base64图片格式字符串
*/
func GetBase64ImgYzm(len, width, height int) (yzm, base64ImgStr string) {
	//config struct for Character
	//字符,公式,验证码配置
	var config = configCharacter{
		Height:             height,
		Width:              width,
		ComplexOfNoiseText: captchaComplexLower,
		ComplexOfNoiseDot:  captchaComplexLower,
		IsShowHollowLine:   true,
		IsShowNoiseDot:     true,
		IsShowNoiseText:    false,
		IsShowSlimeLine:    false,
		IsShowSineLine:     false,
		CaptchaLen:         len,
	}

	//创建验证码图片
	yzm = xcm.GetRandomString(int64(len), xcm.YZMSTR)
	capC := engineCharCreate(yzm, config)
	//以base64编码
	base64ImgStr = captchaWriteToBase64Encoding(capC)
	return
}

// captchaWriteToBase64Encoding converts captcha to base64 encoding string.
// mimeType is one of "audio/wav" "image/png".
func captchaWriteToBase64Encoding(cap captchaInterface) string {
	binaryData, _ := cap.BinaryEncoding()
	return fmt.Sprintf("data:%s;base64,%s", "image/png", base64.StdEncoding.EncodeToString(binaryData))
}

type point struct {
	X int
	Y int
}

//ConfigCharacter captcha config for captcha-engine-characters.
type configCharacter struct {
	// Height png height in pixel.
	// 图像验证码的高度像素.
	Height int
	// Width Captcha png width in pixel.
	// 图像验证码的宽度像素
	Width int
	//IsUseSimpleFont is use simply font(...base64Captcha/fonts/RitaSmith.ttf).
	IsUseSimpleFont bool
	//ComplexOfNoiseText text noise count.
	ComplexOfNoiseText int
	//ComplexOfNoiseDot dot noise count.
	ComplexOfNoiseDot int
	//IsShowHollowLine is show hollow line.
	IsShowHollowLine bool
	//IsShowNoiseDot is show noise dot.
	IsShowNoiseDot bool
	//IsShowNoiseText is show noise text.
	IsShowNoiseText bool
	//IsShowSlimeLine is show slime line.
	IsShowSlimeLine bool
	//IsShowSineLine is show sine line.
	IsShowSineLine bool
	// CaptchaLen Default number of digits in captcha solution.
	// 默认数字验证长度6.
	CaptchaLen int
	//BgColor captcha image background color (optional)
	//背景颜色
	BgColor *color.RGBA
}

//engineCharCreate create captcha with config struct.
func engineCharCreate(id string, config configCharacter) *captchaImageChar {
	var bgc color.RGBA
	if config.BgColor != nil {
		bgc = *config.BgColor
	} else {
		bgc = randLightColor()
	}
	captchaImage := newCaptchaImage(config.Width, config.Height, bgc)
	//背景有像素点干扰
	if config.IsShowNoiseDot {
		captchaImage.drawNoise(config.ComplexOfNoiseDot)
	}
	//波浪线-比较丑
	if config.IsShowHollowLine {
		captchaImage.drawHollowLine()
	}
	//背景有文字干扰
	if config.IsShowNoiseText {
		captchaImage.drawTextNoise(config.ComplexOfNoiseText, config.IsUseSimpleFont)
	}
	//画 细直线 (n 条)
	if config.IsShowSlimeLine {
		captchaImage.drawSlimLine(3)
	}
	//画 多个小波浪线
	if config.IsShowSineLine {
		captchaImage.drawSineLine()
	}
	//写入string
	captchaImage.drawText(id, config.IsUseSimpleFont)
	return captchaImage
}
