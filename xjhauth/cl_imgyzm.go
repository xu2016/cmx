package xjhauth

import (
	"bytes"
	"cmxu/xcm"
	"encoding/base64"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

const (
	imageStringDpi = 72.0
	//MimeTypeCaptchaImage output base64 mine-type.
	MimeTypeCaptchaImage = "image/png"
	//FileExtCaptchaImage output file extension.
	FileExtCaptchaImage = "png"
)
const (
	//CaptchaComplexLower complex level lower.
	CaptchaComplexLower = iota
	//CaptchaComplexMedium complex level medium.
	CaptchaComplexMedium
	//CaptchaComplexHigh complex level high.
	CaptchaComplexHigh
)
const (
	//CaptchaModeNumber mode number.
	CaptchaModeNumber = iota
	//CaptchaModeAlphabet mode alphabet.
	CaptchaModeAlphabet
	//CaptchaModeArithmetic mode arithmetic.
	CaptchaModeArithmetic
	//CaptchaModeNumberAlphabet mode mix number and alphabet,this is also default mode.
	CaptchaModeNumberAlphabet
)

//GoTestOutputDir run go test command where the png and wav file output

const (
	// DefaultLen Default number of digits in captcha solution.
	// 默认数字验证长度.
	DefaultLen = 6
	// MaxSkew max absolute skew factor of a single digit.
	// 图像验证码的最大干扰洗漱.
	MaxSkew = 0.7
	// DotCount Number of background circles.
	// 图像验证码干扰圆点的数量.
	DotCount = 20
)
const (
	digitFontWidth     = 11
	digitFontHeight    = 18
	digitFontBlackChar = 1
)

// CaptchaInterface captcha interface for captcha engine to to write staff
type CaptchaInterface interface {
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
	var configC = ConfigCharacter{
		Height:             height,
		Width:              width,
		ComplexOfNoiseText: CaptchaComplexLower,
		ComplexOfNoiseDot:  CaptchaComplexLower,
		IsShowHollowLine:   true,
		IsShowNoiseDot:     true,
		IsShowNoiseText:    false,
		IsShowSlimeLine:    false,
		IsShowSineLine:     false,
		CaptchaLen:         len,
	}

	//create a characters captcha.
	yzm, capC := GenerateCaptcha(6, xcm.NSUSTR, configC)
	//以base64编码
	base64ImgStr = CaptchaWriteToBase64Encoding(capC)
	GYzm.Add(yzm, "")
	return
}

// CaptchaWriteToBase64Encoding converts captcha to base64 encoding string.
// mimeType is one of "audio/wav" "image/png".
func CaptchaWriteToBase64Encoding(cap CaptchaInterface) string {
	binaryData, _ := cap.BinaryEncoding()
	return fmt.Sprintf("data:%s;base64,%s", "image/png", base64.StdEncoding.EncodeToString(binaryData))
}

// CaptchaItem captcha basic information.
type CaptchaItem struct {
	// Content captcha entity content.
	Content string
	// VerifyValue captcha verify value.
	VerifyValue string
	// ImageWidth image width pixel.
	ImageWidth int
	// ImageHeight image height pixel.
	ImageHeight int
}

//GenerateCaptcha first parameter is empty string,so the package will generate a random uuid for you.
// 	idKeyC,capC := base64Captcha.GenerateCaptcha("",configC)
// 	//write to base64 string.
// 	base64stringC := base64Captcha.CaptchaWriteToBase64Encoding(capC)
func GenerateCaptcha(len, stp int, config ConfigCharacter) (id string, captchaInstance CaptchaInterface) {
	id = xcm.GetRandomString(int64(len), stp)
	captchaInstance = EngineCharCreate(id, config)
	return
}

var trueTypeFontFamilys = readFontsToSliceOfTrueTypeFonts()

//CaptchaImageChar captcha-engine-char return type.
type CaptchaImageChar struct {
	CaptchaItem
	nrgba   *image.NRGBA
	Complex int
}

//ConfigCharacter captcha config for captcha-engine-characters.
type ConfigCharacter struct {
	// Height png height in pixel.
	// 图像验证码的高度像素.
	Height int
	// Width Captcha png width in pixel.
	// 图像验证码的宽度像素
	Width int
	//Mode : base64captcha.CaptchaModeNumber=0, base64captcha.CaptchaModeAlphabet=1, base64captcha.CaptchaModeArithmetic=2, base64captcha.CaptchaModeNumberAlphabet=3.
	Mode int
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
type point struct {
	X int
	Y int
}

//newCaptchaImage new blank captchaImage context.
//新建一个图片对象.
func newCaptchaImage(width int, height int, bgColor color.RGBA) (cImage *CaptchaImageChar, err error) {
	m := image.NewNRGBA(image.Rect(-8, -5, width, height))
	draw.Draw(m, m.Bounds(), &image.Uniform{bgColor}, image.ZP, draw.Src)
	cImage = &CaptchaImageChar{}
	cImage.nrgba = m
	cImage.ImageHeight = height
	cImage.ImageWidth = width
	err = nil
	return
}

//drawHollowLine draw strong and bold white line.
//添加一个较粗的空白直线
func (captcha *CaptchaImageChar) drawHollowLine() *CaptchaImageChar {
	first := captcha.ImageWidth / 20
	end := first * 19
	lineColor := color.RGBA{R: 245, G: 250, B: 251, A: 255}
	x1 := float64(rand.Intn(first))
	x2 := float64(rand.Intn(first) + end)
	multiple := float64(rand.Intn(5)+3) / float64(5)
	if int(multiple*10)%3 == 0 {
		multiple = multiple * -1.0
	}
	w := captcha.ImageHeight / 20
	for ; x1 < x2; x1++ {
		y := math.Sin(x1*math.Pi*multiple/float64(captcha.ImageWidth)) * float64(captcha.ImageHeight/3)
		if multiple < 0 {
			y = y + float64(captcha.ImageHeight/2)
		}
		captcha.nrgba.Set(int(x1), int(y), lineColor)
		for i := 0; i <= w; i++ {
			captcha.nrgba.Set(int(x1), int(y)+i, lineColor)
		}
	}
	return captcha
}

//drawSineLine draw a sine line.
//画一条正弦曲线.
func (captcha *CaptchaImageChar) drawSineLine() *CaptchaImageChar {
	var py float64
	//振幅
	a := rand.Intn(captcha.ImageHeight / 2)
	//Y轴方向偏移量
	b := random(int64(-captcha.ImageHeight/4), int64(captcha.ImageHeight/4))
	//X轴方向偏移量
	f := random(int64(-captcha.ImageHeight/4), int64(captcha.ImageHeight/4))
	// 周期
	var t float64
	if captcha.ImageHeight > captcha.ImageWidth/2 {
		t = random(int64(captcha.ImageWidth/2), int64(captcha.ImageHeight))
	} else if captcha.ImageHeight == captcha.ImageWidth/2 {
		t = float64(captcha.ImageHeight)
	} else {
		t = random(int64(captcha.ImageHeight), int64(captcha.ImageWidth/2))
	}
	w := float64((2 * math.Pi) / t)
	// 曲线横坐标起始位置
	px1 := 0
	px2 := int(random(int64(float64(captcha.ImageWidth)*0.8), int64(captcha.ImageWidth)))

	c := color.RGBA{R: uint8(rand.Intn(150)), G: uint8(rand.Intn(150)), B: uint8(rand.Intn(150)), A: uint8(255)}

	for px := px1; px < px2; px++ {
		if w != 0 {
			py = float64(a)*math.Sin(w*float64(px)+f) + b + (float64(captcha.ImageWidth) / float64(5))
			i := captcha.ImageHeight / 5
			for i > 0 {
				captcha.nrgba.Set(px+i, int(py), c)
				i--
			}
		}
	}
	return captcha
}

//drawSlimLine draw n slim-random-color lines.
//画n条随机颜色的细线
func (captcha *CaptchaImageChar) drawSlimLine(num int) *CaptchaImageChar {
	first := captcha.ImageWidth / 10
	end := first * 9
	y := captcha.ImageHeight / 3
	for i := 0; i < num; i++ {
		point1 := point{X: rand.Intn(first), Y: rand.Intn(y)}
		point2 := point{X: rand.Intn(first) + end, Y: rand.Intn(y)}
		if i%2 == 0 {
			point1.Y = rand.Intn(y) + y*2
			point2.Y = rand.Intn(y)
		} else {
			point1.Y = rand.Intn(y) + y*(i%2)
			point2.Y = rand.Intn(y) + y*2
		}
		captcha.drawBeeline(point1, point2, randDeepColor())
	}
	return captcha
}
func (captcha *CaptchaImageChar) drawBeeline(point1 point, point2 point, lineColor color.RGBA) {
	dx := math.Abs(float64(point1.X - point2.X))

	dy := math.Abs(float64(point2.Y - point1.Y))
	sx, sy := 1, 1
	if point1.X >= point2.X {
		sx = -1
	}
	if point1.Y >= point2.Y {
		sy = -1
	}
	err := dx - dy
	for {
		captcha.nrgba.Set(point1.X, point1.Y, lineColor)
		captcha.nrgba.Set(point1.X+1, point1.Y, lineColor)
		captcha.nrgba.Set(point1.X-1, point1.Y, lineColor)
		captcha.nrgba.Set(point1.X+2, point1.Y, lineColor)
		captcha.nrgba.Set(point1.X-2, point1.Y, lineColor)
		if point1.X == point2.X && point1.Y == point2.Y {
			return
		}
		e2 := err * 2
		if e2 > -dy {
			err -= dy
			point1.X += sx
		}
		if e2 < dx {
			err += dx
			point1.Y += sy
		}
	}
}

//drawNoise draw noise dots.
//画干扰点.
func (captcha *CaptchaImageChar) drawNoise(complex int) *CaptchaImageChar {
	density := 18
	if complex == CaptchaComplexLower {
		density = 28
	} else if complex == CaptchaComplexMedium {
		density = 18
	} else if complex == CaptchaComplexHigh {
		density = 8
	}
	maxSize := (captcha.ImageHeight * captcha.ImageWidth) / density
	for i := 0; i < maxSize; i++ {
		rw := rand.Intn(captcha.ImageWidth)
		rh := rand.Intn(captcha.ImageHeight)
		captcha.nrgba.Set(rw, rh, randColor())
		size := rand.Intn(maxSize)
		if size%3 == 0 {
			captcha.nrgba.Set(rw+1, rh+1, randColor())
		}
	}
	return captcha
}

//drawTextNoise draw noises which are single character.
//画文字噪点.
func (captcha *CaptchaImageChar) drawTextNoise(complex int, isSimpleFont bool) error {
	density := 1500
	if complex == CaptchaComplexLower {
		density = 2000
	} else if complex == CaptchaComplexMedium {
		density = 1500
	} else if complex == CaptchaComplexHigh {
		density = 1000
	}
	maxSize := (captcha.ImageHeight * captcha.ImageWidth) / density
	c := freetype.NewContext()
	c.SetDPI(imageStringDpi)

	c.SetClip(captcha.nrgba.Bounds())
	c.SetDst(captcha.nrgba)
	c.SetHinting(font.HintingFull)
	rawFontSize := float64(captcha.ImageHeight) / (1 + float64(rand.Intn(7))/float64(10))

	for i := 0; i < maxSize; i++ {

		rw := rand.Intn(captcha.ImageWidth)
		rh := rand.Intn(captcha.ImageHeight)

		text := xcm.GetRandomString(1, xcm.KEYSTR)
		fontSize := rawFontSize/2 + float64(rand.Intn(5))

		c.SetSrc(image.NewUniform(randLightColor()))
		c.SetFontSize(fontSize)

		if isSimpleFont {
			c.SetFont(trueTypeFontFamilys[0])
		} else {
			f := randFontFamily()
			c.SetFont(f)
		}

		pt := freetype.Pt(rw, rh)

		if _, err := c.DrawString(text, pt); err != nil {
			log.Println(err)
		}
	}
	return nil
}

//drawText draw captcha string to image.把文字写入图像验证码
func (captcha *CaptchaImageChar) drawText(text string, isSimpleFont bool) error {
	c := freetype.NewContext()
	c.SetDPI(imageStringDpi)

	c.SetClip(captcha.nrgba.Bounds())
	c.SetDst(captcha.nrgba)
	c.SetHinting(font.HintingFull)

	fontWidth := captcha.ImageWidth / len(text)

	for i, s := range text {
		fontSize := float64(captcha.ImageHeight) / (1 + float64(rand.Intn(3))/float64(9))
		c.SetSrc(image.NewUniform(randDeepColor()))
		c.SetFontSize(fontSize)
		if isSimpleFont {
			c.SetFont(trueTypeFontFamilys[0])
		} else {
			f := randFontFamily()
			c.SetFont(f)
		}
		x := int(fontWidth)*i + int(fontWidth)/int(fontSize)
		y := 5 + rand.Intn(captcha.ImageHeight/2) + int(fontSize/2)
		pt := freetype.Pt(x, y)
		if _, err := c.DrawString(string(s), pt); err != nil {
			log.Println(err)
		}
	}
	return nil

}

//EngineCharCreate create captcha with config struct.
func EngineCharCreate(id string, config ConfigCharacter) *CaptchaImageChar {
	var bgc color.RGBA
	if config.BgColor != nil {
		bgc = *config.BgColor
	} else {
		bgc = randLightColor()
	}
	captchaImage, err := newCaptchaImage(config.Width, config.Height, bgc)
	//背景有像素点干扰
	if config.IsShowNoiseDot {
		captchaImage.drawNoise(config.ComplexOfNoiseDot)
	}
	//波浪线       比较丑
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
	captchaImage.VerifyValue = id
	//写入string
	captchaImage.drawText(id, config.IsUseSimpleFont)
	captchaImage.Content = id

	if err != nil {
		log.Println(err)
	}

	return captchaImage
}

//BinaryEncoding save captcha image to binary.
//保存图片到io.
func (captcha *CaptchaImageChar) BinaryEncoding() (bstrs []byte, err error) {
	var buf bytes.Buffer
	if err = png.Encode(&buf, captcha.nrgba); err != nil {
		return
	}
	bstrs, err = buf.Bytes(), nil
	return
}

// WriteTo writes captcha image in PNG format into the given writer.
func (captcha *CaptchaImageChar) WriteTo(w io.Writer) (m int64, err error) {
	b, err := captcha.BinaryEncoding()
	if err != nil {
		return
	}
	n, err := w.Write(b)
	m = int64(n)
	return
}

//Random get random in min between max. 生成指定大小的随机数.
func random(min int64, max int64) float64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if max <= min {
		panic(fmt.Sprintf("invalid range %d >= %d", max, min))
	}
	decimal := r.Float64()

	if max <= 0 {
		return (float64(r.Int63n((min*-1)-(max*-1))+(max*-1)) + decimal) * -1
	}
	if min < 0 && max > 0 {
		if r.Int()%2 == 0 {
			return float64(r.Int63n(max)) + decimal
		}
		return (float64(r.Int63n(min*-1)) + decimal) * -1
	}
	return float64(r.Int63n(max-min)+min) + decimal
}

//randDeepColor get random deep color. 随机生成深色系.
func randDeepColor() color.RGBA {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randColor := randColor()
	increase := float64(30 + r.Intn(255))
	red := math.Abs(math.Min(float64(randColor.R)-increase, 255))
	green := math.Abs(math.Min(float64(randColor.G)-increase, 255))
	blue := math.Abs(math.Min(float64(randColor.B)-increase, 255))
	return color.RGBA{R: uint8(red), G: uint8(green), B: uint8(blue), A: uint8(255)}
}

//randLightColor get random ligth color. 随机生成浅色.
func randLightColor() color.RGBA {
	red, _ := xcm.GetRandomInt(0, 56)
	green, _ := xcm.GetRandomInt(0, 56)
	blue, _ := xcm.GetRandomInt(0, 56)
	return color.RGBA{R: uint8(red + 200), G: uint8(green + 200), B: uint8(blue + 200), A: uint8(255)}
}

//randColor get random color. 生成随机颜色.
func randColor() color.RGBA {
	red, _ := xcm.GetRandomInt(0, 256)
	green, _ := xcm.GetRandomInt(0, 256)
	var blue int
	if (red + green) > 400 {
		blue = 0
	} else {
		blue = 400 - green - red
	}
	if blue > 255 {
		blue = 255
	}
	return color.RGBA{R: uint8(red), G: uint8(green), B: uint8(blue), A: uint8(255)}
}

//readFontsToSliceOfTrueTypeFonts import fonts from dir.
func readFontsToSliceOfTrueTypeFonts() []*truetype.Font {
	fonts := make([]*truetype.Font, 0)
	assetFontNames := []string{"fonts/actionj.ttf", "fonts/RitaSmith.ttf", "fonts/chromohv.ttf", "fonts/Flim-Flam.ttf", "fonts/ApothecaryFont.ttf", "fonts/3Dumb.ttf"}
	for _, assetName := range assetFontNames {
		fonts = appendAssetFontToTrueTypeFonts(assetName, fonts)
	}
	return fonts
}
func appendAssetFontToTrueTypeFonts(assetName string, fonts []*truetype.Font) []*truetype.Font {
	fontBytes, _ := Asset(assetName)
	trueTypeFont, _ := freetype.ParseFont(fontBytes)
	fonts = append(fonts, trueTypeFont)
	return fonts
}

//randFontFamily choose random font family.选择随机的字体
func randFontFamily() *truetype.Font {
	fontCount := len(trueTypeFontFamilys)
	index := rand.Intn(fontCount)
	return trueTypeFontFamilys[index]
}
