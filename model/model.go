/*
 * @Author: gongluck
 * @Date: 2020-03-15 19:24:56
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-03-17 17:37:48
 */

package model

import (
	"giaogiao/config"
	"giaogiao/util"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func init() {
	var err error
	DB, err = gorm.Open("mysql", config.MysqlConfig)
	util.CheckErr(err)

	DB.AutoMigrate(&Video{})
}

func Uninit() {
	err := DB.Close()
	util.CheckErr(err)
}
