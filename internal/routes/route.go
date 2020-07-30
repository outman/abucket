package routes

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
	"github.com/gin-gonic/gin"
	"github.com/outman/abucket/internal/api"
	"github.com/spf13/viper"
)

type route struct{}

// NewRoute export route
func NewRoute() *route {
	return &route{}
}

// Register export gin.Engine
func (r *route) Register() *gin.Engine {

	e := api.NewActionExperiment()
	l := api.NewActionLayer()

	gin.SetMode(viper.GetString("GIN_MODE"))
	router := gin.Default()
	v1 := router.Group("/api/v1/admin")
	{
		v1.GET("/experiment/index", e.Index)
		v1.POST("/experiment/create", e.Create)
		v1.POST("/experiment/update", e.Update)
		v1.POST("/experiment/delete", e.Delete)
		v1.POST("/experiment/group/create", e.CreateGroup)
		v1.GET("/layer/index", l.Index)
		v1.POST("/layer/create", l.Create)
	}
	return router
}
