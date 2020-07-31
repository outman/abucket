package middleware

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

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/outman/abucket/internal/pkg"
	"github.com/spf13/viper"
)

// Auth check
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("AccessToken")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"code": http.StatusForbidden,
				"text": "Un auth",
			})
			return
		}
		srcValue := decryptAccessToken(token)
		admin := viper.GetStringMap("ADMIN")
		if srcValue == "" || admin[srcValue] == "" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"code": http.StatusForbidden,
				"text": "Un auth",
			})
			return
		}
		c.Next()
	}
}

func decryptAccessToken(token string) string {
	v := pkg.DecryptByAes(token)
	if v != "" {
		n := strings.Split(v, "|")
		return n[0]
	}
	return ""
}
