/*
 * @Author: gongluck
 * @Date: 2020-03-16 18:57:38
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-03-17 17:37:36
 */

package api

import (
	"giaogiao/model"
	"net/http"
	"path"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Contribute(c *gin.Context) {
	var v model.Video
	v.Title = c.Request.FormValue("title")
	v.Info = c.Request.FormValue("info")

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Status: http.StatusBadRequest,
			Msg:    "Can not find upload file!",
			Error:  err.Error(),
		})
		return
	}

	model.DB.Save(&v)

	err = c.SaveUploadedFile(file, "SaveVideos/"+strconv.Itoa(int(v.ID))+path.Ext(file.Filename))
	if err != nil {
		model.DB.Delete(&v)
		c.JSON(http.StatusInternalServerError, Response{
			Status: http.StatusInternalServerError,
			Msg:    "Can not save upload file!",
			Error:  err.Error(),
		})
		return
	}

	v.Url = "/videos/" + strconv.Itoa(int(v.ID)) + path.Ext(file.Filename)
	model.DB.Update(&v)

	c.JSON(http.StatusOK, Response{
		Status: 0,
		Msg:    "Upload succeed.",
		Data:   "ID:" + strconv.Itoa(int(v.ID)) + ",Url:" + v.Url,
	})
}
