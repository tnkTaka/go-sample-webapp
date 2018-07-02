package controller

import (
	"github.com/go-sample-webapp/model"
)

type Task struct {
}

func NewTask() Task {
	return Task{}
}

// 指定したIDのタスク情報
func (c Task) GetId(id int) interface{} {
	repo := model.NewTaskRepository()
	task := repo.GetByTaskID(id)
	return task
}

// 全タスク情報
func (c Task) GetAllTask() interface{} {
	repo := model.NewTaskRepository()
	tasks := repo.GetAllTask()
	return tasks
}

// タスク作成
func (c Task) CreateTask(text string) bool {
	repo := model.NewTaskRepository()
	task := repo.CreateTask(text)
	return task
}

// タスク更新
func (c Task) UpdateTask(id int, text string) bool {
	repo := model.NewTaskRepository()
	task := repo.UpdateTask(id, text)
	return task
}

// タスク削除
func (c Task) DeleteTask(id int) bool {
	repo := model.NewTaskRepository()
	task := repo.DeleteTask(id)
	return task
}
