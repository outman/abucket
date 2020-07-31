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
	"time"

	"github.com/gin-contrib/cors"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/outman/abucket/internal/api"
	"github.com/outman/abucket/internal/pkg"
	"github.com/spf13/viper"
)

type route struct{}

// NewRoute export route
func NewRoute() *route {
	return &route{}
}

// Register export gin.Engine
func (r *route) Register() *gin.Engine {

	gin.DisableConsoleColor()
	gin.SetMode(viper.GetString("GIN_MODE"))
	router := gin.New()

	// zap logger init
	logger := pkg.NewZapLogger()
	router.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	router.Use(ginzap.RecoveryWithZap(logger, true))

	// cors
	router.Use(cors.New(cors.Config{
		AllowOrigins:     viper.GetStringSlice("CORS_ALLOW_ORIGINS"),
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// route
	e := api.NewActionExperiment()
	l := api.NewActionLayer()

	admin := router.Group("/api/v1/admin")
	{
		admin.GET("/experiment/index", e.Index)
		admin.POST("/experiment/create", e.Create)
		admin.POST("/experiment/update", e.Update)
		admin.POST("/experiment/delete", e.Delete)
		admin.POST("/experiment/group/create", e.CreateGroup)
		admin.GET("/layer/index", l.Index)
		admin.POST("/layer/create", l.Create)
	}

	g := api.NewActionGroup()
	group := router.Group("/api/v1/group")
	{
		group.GET("/get", g.Group)
	}
	return router
}
