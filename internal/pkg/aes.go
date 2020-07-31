package pkg

import (
	"encoding/base64"
	"time"

	"github.com/forgoer/openssl"
	"github.com/spf13/viper"
)

/*
Copyright © 2020 pochonlee@gmail.com

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

// EncryptByAes plaintext
func EncryptByAes(plaintext string) string {
	key := []byte(viper.GetString("CRYPT_KEY"))
	src := []byte(plaintext)
	dst, err := openssl.AesECBEncrypt(src, key, openssl.PKCS7_PADDING)
	if err != nil {
		return ""
	}

	return base64.StdEncoding.EncodeToString(dst)
}

// DecryptByAes dst
func DecryptByAes(dst string) string {
	ept, err := base64.StdEncoding.DecodeString(dst)
	if err != nil {
		return ""
	}
	key := []byte(viper.GetString("CRYPT_KEY"))
	src, err := openssl.AesECBDecrypt(ept, key, openssl.PKCS7_PADDING)
	if err != nil {
		return ""
	}
	return string(src)
}

// GetAccessToken name
func GetAccessToken(name string) string {
	return EncryptByAes(name + "|" + time.Now().String())
}
