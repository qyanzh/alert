/**
 * @author  qyanzh
 * @create  2022/03/07 19:02
 */

package service

import (
	"alert/internal/dao"
	"alert/internal/model"
	"time"
)

type TaskService struct {
	taskDao *dao.TaskDao
}

func NewTaskService() *TaskService {
	return &TaskService{taskDao: dao.NewTaskDao()}
}

func (ts *TaskService) AddTask(code, name string, ruleID, frequency uint) (*model.Task, error) {
	newTask := model.Task{
		Code:      code,
		Name:      name,
		RuleID:    ruleID,
		Frequency: frequency,
		Enable:    false,
		Status:    model.TSReady,
		NextTime:  time.Now(),
	}
	_, err := ts.taskDao.AddTask(&newTask)
	if err != nil {
		return nil, err
	}
	return &newTask, nil
}

func (ts *TaskService) SelectTaskByID(id uint) (*model.Task, error) {
	return ts.taskDao.SelectTaskByID(id)
}

func (ts *TaskService) SelectTaskByCode(code string) (*model.Task, error) {
	return ts.taskDao.SelectTaskByCode(code)
}

func (ts *TaskService) SelectRunnableTasks() (*[]model.Task, error) {
	return ts.taskDao.SelectRunnableTasks()
}

func (ts *TaskService) UpdateTask(task *model.Task) error {
	_, err := ts.taskDao.UpdateTask(task)
	return err
}

func (ts *TaskService) UpdateTaskEnableByCode(code string, enable bool) error {
	_, err := ts.taskDao.UpdateTaskEnableByCode(code, enable)
	return err
}

func (ts *TaskService) UpdateTaskSuccessByCode(code string, nextTime *time.Time) error {
	_, err := ts.taskDao.UpdateTaskStatusAndNextTimeByCode(code, model.TSReady, nextTime)
	if err != nil {
		return err
	}
	return nil
}

func (ts *TaskService) UpdateTaskStatusRunningByCodes(codes *[]string) error {
	_, err := ts.taskDao.UpdateTaskStatusByCodes(codes, model.TSRunning)
	if err != nil {
		return err
	}
	return nil
}

func (ts *TaskService) UpdateTaskFailByCode(code string) error {
	_, err := ts.taskDao.UpdateTaskStatusAndNextTimeByCode(code, model.TSFail, nil)
	if err != nil {
		return err
	}
	return nil
}

func (ts *TaskService) DeleteTaskByCode(code string) error {
	_, err := ts.taskDao.DeleteTaskByCode(code, false)
	return err
}
