package api

import (
	"WestOnline2/pkg/utils"
	"WestOnline2/serivice"

	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	createService := serivice.CreateTaskService{}
	chaim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&createService); err == nil {
		res := createService.Create(chaim.Id)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}

func ShowTask(c *gin.Context) {
	var showTask serivice.ShowTaskService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&showTask); err == nil {
		res := showTask.Show(claim.Id)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}

func ListTask(c *gin.Context) {
	var listTask serivice.ListTasksService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&listTask); err == nil {
		res := listTask.List(claim.Id)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}
func UpdateTask(c *gin.Context) {
	updateTaskService := serivice.UpdateTaskService{}
	if err := c.ShouldBind(&updateTaskService); err == nil {
		res := updateTaskService.Update(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}

func SearchTasks(c *gin.Context) {
	searchTaskService := serivice.SearchTaskService{}
	chaim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&searchTaskService); err == nil {
		res := searchTaskService.Search(chaim.Id)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}

func DeleteTask(c *gin.Context) {
	deleteTaskService := serivice.DeleteTaskService{}
	res := deleteTaskService.Delete(c.Param("id"))
	c.JSON(200, res)
}
