package qrcode

import qrCode "github.com/skip2/go-qrcode"

func GenerateQrCode(text string) ([]byte, error) {
	png, err := qrCode.Encode(text, qrCode.Medium, 256)
	if err != nil {
		return nil, err
	}
	return png, nil
}
