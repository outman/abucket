package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/outman/abucket/internal/form"
	"github.com/outman/abucket/internal/service"
)

type actionLayer struct {
}

func NewActionLayer() *actionLayer {
	return &actionLayer{}
}

func (a *actionLayer) Index(c *gin.Context) {
	s := service.NewLayerService()
	data := s.Index()
	c.JSON(http.StatusOK, gin.H{
		"code": actionSuccess,
		"data": data,
	})
}

func (a *actionLayer) Create(c *gin.Context) {
	var formParams form.FormCreateLayer
	if err := c.ShouldBind(&formParams); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": actionParameterError,
			"text": err.Error(),
		})
		return
	}

	s := service.NewLayerService()
	code, data := s.Create(&formParams)
	if code != service.ServiceOptionSuccess {
		c.JSON(http.StatusOK, gin.H{
			"code": actionServerError,
			"text": service.CodeMessage(code),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": actionSuccess,
		"data": data,
	})
}
