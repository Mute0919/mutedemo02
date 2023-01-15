package serivice

import (
	"WestOnline2/model"
	//"WestOnline2/pkg/e"
	//"WestOnline2/pkg/util"
	"WestOnline2/serializer"
	"time"
)

type ShowTaskService struct {
}

type DeleteTaskService struct {
}

type UpdateTaskService struct {
	ID      uint   `form:"id" json:"id"`
	Title   string `form:"title" json:"title" binding:"required,min=2,max=100"`
	Content string `form:"content" json:"content" binding:"max=1000"`
	Status  int    `form:"status" json:"status"` //0 待办   1已完成
}

type CreateTaskService struct {
	Title   string `form:"title" json:"title" binding:"required,min=2,max=100"`
	Content string `form:"content" json:"content" binding:"max=1000"`
	Status  int    `form:"status" json:"status"` //0 待办   1已完成
}

type SearchTaskService struct {
	Info string `form:"info" json:"info"`
}

type ListTasksService struct {
	Limit int `form:"limit" json:"limit"`
	Start int `form:"start" json:"start"`
}

func (service *CreateTaskService) Create(id uint) serializer.Response {
	var user model.User
	model.DB.First(&user, id)
	task := model.Task{
		User:      user,
		Uid:       user.ID,
		Title:     service.Title,
		Content:   service.Content,
		Status:    0,
		StartTime: time.Now().Unix(),
		EndTime:   0,
	}
	err := model.DB.Create(&task).Error
	if err != nil {
		code := 500
		return serializer.Response{
			Status: code,
			Msg:    "创建目录失败",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "成功",
	}
}
func (service *ShowTaskService) Show(id uint) serializer.Response {
	var task model.Task
	code := 200
	err := model.DB.First(&task, id).Error
	if err != nil {
		code = 500
		return serializer.Response{
			Status: code,
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Data:   serializer.BuildTask(task),
	}
}

func (service *ListTasksService) List(id uint) serializer.Response {
	var tasks []model.Task
	var total int64
	if service.Limit == 0 {
		service.Limit = 15
	}
	model.DB.Model(model.Task{}).Preload("User").Where("uid = ?", id).Count(&total).
		Limit(service.Limit).Offset((service.Start - 1) * service.Limit).
		Find(&tasks)
	return serializer.BuildListResponse(serializer.BuildTasks(tasks), uint(total))
}

func (service *DeleteTaskService) Delete(id string) serializer.Response {
	var task model.Task
	code := 200
	err := model.DB.First(&task, id).Error
	if err != nil {
		code = 500
		return serializer.Response{
			Status: code,
			Error:  err.Error(),
		}
	}
	err = model.DB.Delete(&task).Error
	if err != nil {
		code = 500
		return serializer.Response{
			Status: code,
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "成功",
	}
}

func (service *UpdateTaskService) Update(id string) serializer.Response {
	var task model.Task
	model.DB.Model(model.Task{}).Where("id = ?", id).First(&task)
	task.Content = service.Content
	task.Status = service.Status
	task.Title = service.Title
	code := 200
	err := model.DB.Save(&task).Error
	if err != nil {
		code = 500
		return serializer.Response{
			Status: code,
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Data:   "修改成功",
	}
}

func (service *SearchTaskService) Search(uId uint) serializer.Response {
	var tasks []model.Task
	code := 200
	model.DB.Where("uid=?", uId).Preload("User").First(&tasks)
	err := model.DB.Model(&model.Task{}).Where("title LIKE ? OR content LIKE ?",
		"%"+service.Info+"%", "%"+service.Info+"%").Find(&tasks).Error
	if err != nil {

		code = 500
		return serializer.Response{
			Status: code,
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildTasks(tasks),
	}
}
