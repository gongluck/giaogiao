/*
 * @Author: gongluck
 * @Date: 2020-03-16 18:34:06
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-03-16 19:39:11
 */

package server

import (
	"giaogiao/api"
	"giaogiao/util"
	"os"

	"github.com/gin-gonic/gin"
)

func CreateServer() *gin.Engine {
	r := gin.Default()

	_, err := os.Stat("./SaveVideos")
	if err != nil {
		if os.IsNotExist(err) {
			err := os.Mkdir("./SaveVideos", os.ModePerm)
			util.CheckErr(err)
		}
	}
	r.Static("/videos", "./SaveVideos")

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("ping", api.Ping)
		apiv1.PUT("contribute", api.Contribute)
	}

	return r
}
