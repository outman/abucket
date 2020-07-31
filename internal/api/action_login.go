package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/outman/abucket/internal/form"
	"github.com/outman/abucket/internal/pkg"
	"github.com/spf13/viper"
)

// errUserPassword error
var errUserPassword = errors.New("用户名或者密码错误")

type actionLogin struct{}

// NewActionLogin *actionLogin
func NewActionLogin() *actionLogin {
	return &actionLogin{}
}

func (a *actionLogin) Login(c *gin.Context) {
	var params form.FormLogin
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": actionParameterError,
			"text": err.Error(),
		})
		return
	}
	admin := viper.GetStringMap("ADMIN")

	if len(admin) == 0 || admin[params.Name] == "" || admin[params.Name] != params.Password {
		c.JSON(http.StatusOK, gin.H{
			"code": actionParameterError,
			"text": errUserPassword.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": actionSuccess,
		"data": pkg.GetAccessToken(params.Name),
	})
}
