package router

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-sample-webapp/controller"
)

func GetIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func GetTask(c *gin.Context) {
	n := c.Query("id")
	id, _ := strconv.Atoi(n)

	ctrl := controller.NewTask()
	result := ctrl.GetId(id)

	c.JSON(http.StatusOK, result)
}

func GetAllTask(c *gin.Context) {
	ctrl := controller.NewTask()
	result := ctrl.GetAllTask()

	c.JSON(http.StatusOK, result)
}

func PostTask(c *gin.Context) {
	text := c.PostForm("text")

	ctrl := controller.NewTask()
	post := ctrl.CreateTask(text)

	if post {
		fmt.Println("作成成功")
	}

	result := ctrl.GetAllTask()

	c.JSON(http.StatusOK, result)
}

func PutTask(c *gin.Context) {
	n := c.PostForm("id")
	id, _ := strconv.Atoi(n)
	text := c.PostForm("text")

	ctrl := controller.NewTask()
	put := ctrl.UpdateTask(id, text)

	if put {
		fmt.Println("更新成功")
	}

	result := ctrl.GetId(id)

	c.JSON(http.StatusOK, result)
}

func DeleteTask(c *gin.Context) {
	n := c.PostForm("id")
	id, _ := strconv.Atoi(n)

	ctrl := controller.NewTask()
	delete := ctrl.DeleteTask(id)

	if delete {
		fmt.Println("削除成功")
	}

	result := ctrl.GetId(id)

	c.JSON(http.StatusOK, result)
}
