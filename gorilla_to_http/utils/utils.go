package utils

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

// GenSHA is used to generate SHA Code
func GenSHA(src multipart.File) string {
	h := sha1.New()
	io.Copy(h, src)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// DecodeSplit decodes base64 string and splits to return slice of strings
func DecodeSplit(data string) []string {
	ckVal, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		fmt.Println(err)
	}
	return strings.Split(string(ckVal), "|")
}

// UploadImage writes to file and returns the file name
func UploadImage(src multipart.File) string {
	fname := GenSHA(src) + ".jpg"
	dst, err := os.Create(filepath.Join("./", "assets", "imgs", fname))
	if err != nil {
		fmt.Println("err", err)
	}
	defer dst.Close()
	src.Seek(0, 0)
	// Copy the image
	io.Copy(dst, src)
	return fname

}
