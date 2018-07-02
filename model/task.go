package model

// 投稿情報の構造体
type Task struct {
	ID   int    `json:"id" `
	Text string `json:"text" `
}

type TaskRepository struct{}

func NewTaskRepository() TaskRepository {
	return TaskRepository{}
}

var task Task
var tasks []Task

// 指定したIDのタスク情報
func (m TaskRepository) GetByTaskID(id int) *Task {
	db := GormConnect()
	task.ID = id
	db.Find(&task)
	return &task
}

// 全タスク情報
func (m TaskRepository) GetAllTask() []Task {
	db := GormConnect()
	db.Find(&tasks)
	return tasks
}

// タスク作成
func (m TaskRepository) CreateTask(text string) bool {
	db := GormConnect()
	db.Create(&Task{Text: text})
	return true
}

// タスク更新
func (m TaskRepository) UpdateTask(id int, text string) bool {
	db := GormConnect()
	task.ID = id
	db.First(&task)
	db.Model(&task).Update("text", text)
	return true
}

// タスク削除
//本来、DELETE文使わずに、フラグの値を変えるなどして対応するがこのコードはDELETE文で行ごと消している。
func (m TaskRepository) DeleteTask(id int) bool {
	db := GormConnect()
	task.ID = id
	db.First(&task)
	db.Delete(&task)
	return true
}
