package utils

import (
	"encoding/base64"
	"os"
)

const FileLength = 100000

func ImgToBase64(imgPath string) string {
	file, _ := os.Open(imgPath)

	bufByte := make([]byte, FileLength)
	n, _ := file.Read(bufByte)
	imgBase64Str := base64.StdEncoding.EncodeToString(bufByte[:n])
	bufByte = []byte{}
	return imgBase64Str
}
