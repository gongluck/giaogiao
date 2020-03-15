/*
 * @Author: gongluck
 * @Date: 2020-03-15 17:16:58
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-03-15 19:27:29
 */

package main

import (
	"log"

	"giaogiao/config"
	_ "giaogiao/model"
)

func main() {
	log.Println("begin to start ...")
	log.Println(config.AddressConfig)
	log.Println(config.MysqlConfig)
}
