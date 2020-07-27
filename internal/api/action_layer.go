package api

import "github.com/gin-gonic/gin"

type actionLayer struct {
}

func NewActionLayer() *actionLayer {
	return &actionLayer{}
}

func (a *actionLayer) Index(c *gin.Context) {

}

func (a *actionLayer) Create(c *gin.Context) {

}

func (a *actionLayer) Update(c *gin.Context) {

}
