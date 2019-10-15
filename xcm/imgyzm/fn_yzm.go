package imgyzm

import (
	"cmx/xcm"
	"fmt"
	"image/color"
	"math"
	"math/rand"
	"time"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
)

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
