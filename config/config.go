/*
 * @Author: gongluck
 * @Date: 2020-03-15 17:16:35
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-03-16 19:41:08
 */

package config

import (
	"fmt"
	"giaogiao/util"

	"gopkg.in/ini.v1"
)

// httpserver address
var AddressConfig string

// mysql config
var MysqlConfig string

func init() {
	cfg, err := ini.Load("./config.ini")
	util.CheckErr(err)
	address := cfg.Section("address")
	ip, err := address.GetKey("ip")
	util.CheckErr(err)
	port, err := address.GetKey("port")
	util.CheckErr(err)
	AddressConfig = ip.String() + ":" + port.String()

	mysql := cfg.Section("mysql")
	ip, err = mysql.GetKey("ip")
	util.CheckErr(err)
	port, err = mysql.GetKey("port")
	util.CheckErr(err)
	user, err := mysql.GetKey("user")
	util.CheckErr(err)
	password, err := mysql.GetKey("password")
	util.CheckErr(err)
	MysqlConfig = fmt.Sprintf("%s:%s@tcp(%s:%s)/giaogiao?charset=utf8", user, password, ip, port)
}
