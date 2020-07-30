package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/outman/abucket/internal/form"
	"github.com/outman/abucket/internal/service"
)

type actionGroup struct {
}

func NewActionGroup() *actionGroup {
	return &actionGroup{}
}

func (a *actionGroup) Group(c *gin.Context) {
	var query form.RequestFormGroup
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": actionParameterError,
			"text": err.Error(),
		})
		return
	}
	s := service.NewExperimentService()
	code, data, bucket := s.HitGroup(&query)
	response := form.ResponseFormGroup{
		Key:      query.UniqKey,
		HitGroup: data,
		Bucket:   bucket,
	}

	if code != service.ServiceOptionSuccess {
		c.JSON(http.StatusOK, gin.H{
			"code": actionServerError,
			"text": service.CodeMessage(code),
			"data": response,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": actionSuccess,
		"data": response,
	})
}
