package xcm

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"image/png"

	ebar "github.com/boombuler/barcode"
	eqr "github.com/boombuler/barcode/qr"
)

//CreateQrPng 生成二维码图片
func CreateQrPng(width, height int, url, key string) (base64QrPng string, err error) {
	b64code := url + key
	code, err := eqr.Encode(b64code, eqr.L, eqr.Unicode)
	if err != nil {
		return
	}
	if b64code != code.Content() {
		err = errors.New("编码不正确")
		return
	}
	code, err = ebar.Scale(code, width, height)
	if err != nil {
		return
	}
	var buf bytes.Buffer
	if err = png.Encode(&buf, code); err != nil {
		return
	}

	return fmt.Sprintf("data:%s;base64,%s", "image/png", base64.StdEncoding.EncodeToString(buf.Bytes())), nil
}
