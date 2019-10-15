package imgyzm

import (
	"bytes"
	"cmx/xcm"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io"
	"log"
	"math"
	"math/rand"

	"github.com/golang/freetype"
	"golang.org/x/image/font"
)

//newCaptchaImage new blank captchaImage context.
func newCaptchaImage(width int, height int, bgColor color.RGBA) (cImage *captchaImageChar) {
	m := image.NewNRGBA(image.Rect(-8, -5, width, height))
	draw.Draw(m, m.Bounds(), &image.Uniform{bgColor}, image.ZP, draw.Src)
	cImage = &captchaImageChar{}
	cImage.nrgba = m
	cImage.ImageHeight = height
	cImage.ImageWidth = width
	return
}

var trueTypeFontFamilys = readFontsToSliceOfTrueTypeFonts()

//CaptchaImageChar captcha-engine-char return type.
type captchaImageChar struct {
	ImageWidth  int
	ImageHeight int
	nrgba       *image.NRGBA
	Complex     int
}

//添加一个较粗的空白直线
func (captcha *captchaImageChar) drawHollowLine() *captchaImageChar {
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

//画一条正弦曲线.
func (captcha *captchaImageChar) drawSineLine() *captchaImageChar {
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

//画n条随机颜色的细线
func (captcha *captchaImageChar) drawSlimLine(num int) *captchaImageChar {
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
func (captcha *captchaImageChar) drawBeeline(point1 point, point2 point, lineColor color.RGBA) {
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

//画干扰点.
func (captcha *captchaImageChar) drawNoise(complex int) *captchaImageChar {
	density := 18
	if complex == captchaComplexLower {
		density = 28
	} else if complex == captchaComplexMedium {
		density = 18
	} else if complex == captchaComplexHigh {
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

//画文字噪点.
func (captcha *captchaImageChar) drawTextNoise(complex int, isSimpleFont bool) error {
	density := 1500
	if complex == captchaComplexLower {
		density = 2000
	} else if complex == captchaComplexMedium {
		density = 1500
	} else if complex == captchaComplexHigh {
		density = 1000
	}
	maxSize := (captcha.ImageHeight * captcha.ImageWidth) / density
	c := freetype.NewContext()
	c.SetDPI(72.0)
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
func (captcha *captchaImageChar) drawText(text string, isSimpleFont bool) error {
	c := freetype.NewContext()
	c.SetDPI(72.0)
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

//BinaryEncoding save captcha image to binary.
//保存图片到io.
func (captcha *captchaImageChar) BinaryEncoding() (bstrs []byte, err error) {
	var buf bytes.Buffer
	if err = png.Encode(&buf, captcha.nrgba); err != nil {
		return
	}
	bstrs, err = buf.Bytes(), nil
	return
}

// WriteTo writes captcha image in PNG format into the given writer.
func (captcha *captchaImageChar) WriteTo(w io.Writer) (m int64, err error) {
	b, err := captcha.BinaryEncoding()
	if err != nil {
		return
	}
	n, err := w.Write(b)
	m = int64(n)
	return
}
