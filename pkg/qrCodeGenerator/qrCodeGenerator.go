package qrCodeGenerator

import (
	"fmt"
	"image"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

type QRCodeGenerator struct{}

func NewQRCodeGenerator() *QRCodeGenerator {
	return &QRCodeGenerator{}
}

func (b *QRCodeGenerator) GenerateCode(data string, width int, height int) (image.Image, error) {
	code, err := qr.Encode(data, qr.M, qr.Auto)
	if err != nil {
		return nil, fmt.Errorf("failed to generate QR code: %w", err)
	}

	code, err = barcode.Scale(code, width, height)
	if err != nil {
		return nil, fmt.Errorf("failed to scale QR code: %w", err)
	}

	return code, nil
}
