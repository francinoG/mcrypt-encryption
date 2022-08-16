package mcrypt

import (
	"encoding/base64"

	"github.com/forgoer/openssl"
)

func Encrypt(key, text []byte) (s string) {
	dst, _ := openssl.AesECBEncrypt(text, key, openssl.PKCS7_PADDING)
	// fmt.Println(base64.StdEncoding.EncodeToString(dst)) // yXVUkR45PFz0UfpbDB8/ew==

	return base64.StdEncoding.EncodeToString(dst)
}

func Decrypt(key []byte, text string) (s string) {
	decodeStr, _ := base64.StdEncoding.DecodeString(text)

	dst, _ := openssl.AesECBDecrypt(decodeStr, key, openssl.PKCS7_PADDING)
	// fmt.Println(string(dst)) // 123456

	return string(dst)
}
