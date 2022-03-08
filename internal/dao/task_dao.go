/**
 * @author  qyanzh
 * @create  2022/03/07 18:23
 */

package dao

import (
	"alert/internal/db"
	"alert/internal/model"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type TaskDao struct {
	db *gorm.DB
}

func NewTaskDao() *TaskDao {
	return &TaskDao{db: db.DbClient}
}

func (td *TaskDao) AddTask(task *model.Task) (int64, error) {
	result := td.db.Create(task)
	if result.Error != nil {
		return 0, fmt.Errorf("adding task %v: %v", *task, result.Error)
	}
	return result.RowsAffected, nil
}

func (td *TaskDao) DeleteTaskByID(ID uint) (int64, error) {
	result := td.db.Delete(&model.Task{}, ID)
	if result.Error != nil {
		return 0, fmt.Errorf("deleting task by id=%d: %v", ID, result.Error)
	}
	return result.RowsAffected, nil
}

func (td *TaskDao) DeleteTaskByCode(code string, permanent bool) (int64, error) {
	where := td.db.Where("code = ?", code)
	if permanent {
		where = where.Unscoped()
	}
	result := where.Delete(&model.Task{})
	if result.Error != nil {
		return 0, fmt.Errorf("deleting task by code=%s: %v", code, result.Error)
	}
	return result.RowsAffected, nil
}

func (td *TaskDao) UpdateTask(task *model.Task) (int64, error) {
	result := td.db.Save(task)
	if result.Error != nil {
		return 0, fmt.Errorf("updating task %v: %v", *task, result.Error)
	}
	return result.RowsAffected, nil
}

func (td *TaskDao) UpdateTaskEnableByCode(code string, enable bool) (int64, error) {
	task := model.Task{Enable: enable}
	result := td.db.Model(&task).Select("enable").Where("code = ?", code).Updates(task)
	if result.Error != nil {
		return 0, fmt.Errorf("updating task enable: %v", result.Error)
	}
	return result.RowsAffected, result.Error
}

func (td *TaskDao) UpdateTaskStatusByCode(code string, status model.TaskStatus, nextTime *time.Time, msg string) (int64, error) {
	return td.UpdateTaskStatusByCodes(&[]string{code}, status, nextTime, msg)
}

func (td *TaskDao) UpdateTaskStatusByCodes(codes *[]string, status model.TaskStatus, nextTime *time.Time, msg string) (int64, error) {
	task := model.Task{Status: status}
	if nextTime != nil {
		task.NextTime = *nextTime
	}
	if msg != "" {
		task.Msg = msg
	}
	result := td.db.Model(&task).Where("code IN ?", *codes).Updates(&task)
	if result.Error != nil {
		return 0, fmt.Errorf("updating task status: %v", result.Error)
	}
	return result.RowsAffected, result.Error
}

func (td *TaskDao) SelectTaskByID(ID uint) (*model.Task, error) {
	task := model.Task{}
	result := td.db.First(&task, ID)
	if result.Error != nil {
		return nil, fmt.Errorf("selecting task by id=%d: %v", ID, result.Error)
	}
	return &task, nil
}

func (td *TaskDao) SelectTaskByCode(code string) (*model.Task, error) {
	task := model.Task{Code: code}
	result := td.db.Where(&task).First(&task)
	if result.Error != nil {
		return nil, fmt.Errorf("selecting task by code=%s: %v", code, result.Error)
	}
	return &task, nil
}

func (td *TaskDao) SelectRunnableTasks() (*[]model.Task, error) {
	var tasks []model.Task
	result := td.db.
		Where("enable = ? AND status = ? AND next_time <= ?",
			true, model.TSReady, time.Now()).
		Find(&tasks)
	if result.Error != nil {
		return nil, fmt.Errorf("selecting runnable tasks: %v", result.Error)
	}
	return &tasks, nil
}
