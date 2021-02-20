package main

import (
	"github.com/gin-gonic/gin"

	"dapan/dbx"
	"dapan/router"
)



func main() {
	r := gin.Default()

	dbx.SetMysqlDB()
	dbx.InitData()

	router.SetView(r)
	r.Run(":8888")
}