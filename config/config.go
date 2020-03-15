/*
 * @Author: gongluck
 * @Date: 2020-03-15 17:16:35
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-03-15 17:17:36
 */

package config

import (
	"log"
	"os"

	"gopkg.in/ini.v1"
)

// httpserver address
var AddressConfig string

func checkErr(err error) {
	if err != nil {
		log.Fatalln("Fail happen,", err)
		os.Exit(1)
	}
}

func init() {
	cfg, err := ini.Load("./config.ini")
	checkErr(err)
	address := cfg.Section("address")
	ip, err := address.GetKey("ip")
	checkErr(err)
	port, err := address.GetKey("port")
	checkErr(err)

	AddressConfig = ip.String() + ":" + port.String()
}
