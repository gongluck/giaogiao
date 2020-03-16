/*
 * @Author: gongluck
 * @Date: 2020-03-16 18:41:17
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-03-16 18:46:59
 */

package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
