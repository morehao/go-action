package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("insertData", InsertData)
	r.GET("clearData", ClearData)
	r.POST("searchData", SearchData)

	r.Run(":8888")
}
