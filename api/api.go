/*
 * @Author: gongluck
 * @Date: 2020-03-16 18:41:17
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-03-17 17:37:39
 */

package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error"`
}

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Status: 0,
		Msg:    "Pong",
	})
}
