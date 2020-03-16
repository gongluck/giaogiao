/*
 * @Author: gongluck
 * @Date: 2020-03-16 18:57:38
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-03-16 19:34:05
 */

package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Contribute(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Status: http.StatusBadRequest,
			Msg:    "Can not find upload file!",
			Error:  err.Error(),
		})
	}

	err = c.SaveUploadedFile(file, "SaveVideos/"+file.Filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Status: http.StatusInternalServerError,
			Msg:    "Can not save upload file!",
			Error:  err.Error(),
		})
	}

	c.JSON(http.StatusOK, Response{
		Status: 0,
		Msg:    "Upload succeed.",
	})
}
