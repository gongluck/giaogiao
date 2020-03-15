/*
 * @Author: gongluck 
 * @Date: 2020-03-15 19:28:23 
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-03-15 19:30:43
 */

 package util

 import(
	 "log"
	 "os"
 )

 func CheckErr(err error) {
	if err != nil {
		log.Fatalln("Fail happen,", err)
		os.Exit(1)
	}
}