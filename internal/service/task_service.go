package service

import "alert/internal/dao"

type TaskService struct {
	alertDao dao.TaskDao
}

func NewTaskService() *TaskService {
	return &TaskService{alertDao: *dao.NewTaskDao()}
}
