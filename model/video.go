/*
 * @Author: gongluck 
 * @Date: 2020-03-15 19:41:43 
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-03-15 19:44:38
 */

 package model

import (
	"github.com/jinzhu/gorm"
)

type Video struct{
	gorm.Model
	Title string
	Info string
	Url string
}