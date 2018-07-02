package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-sample-webapp/router"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("view/*")

	r.GET("/", router.GetIndex)
	r.GET("/tasks", router.GetAllTask)
	r.GET("/get", router.GetTask)
	r.POST("/post", router.PostTask)

	/*
		本来、更新(UPDATE)と削除(DELETE)を行う場は、PUT・DELETEメソッドを使うがHTML5で
		PUT・DELETEのメソッドがサポートされていないので、このテンプレートではPOSTを使う。
	*/
	r.POST("/put", router.PutTask)
	r.POST("/delete", router.DeleteTask)

	r.Run(":8080")
}
