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

func (dao *TaskDao) AddTask(task *model.Task) (int64, error) {
	result := dao.db.Create(task)
	if result.Error != nil {
		return 0, fmt.Errorf("adding task %v: %v", *task, result.Error)
	}
	return result.RowsAffected, nil
}

func (dao *TaskDao) DeleteTaskByID(ID uint) (int64, error) {
	result := dao.db.Delete(&model.Task{}, ID)
	if result.Error != nil {
		return 0, fmt.Errorf("deleting task by id=%d: %v", ID, result.Error)
	}
	return result.RowsAffected, nil
}

func (dao *TaskDao) DeleteTaskByCode(code string, permanent bool) (int64, error) {
	where := dao.db.Where("code = ?", code)
	if permanent {
		where = where.Unscoped()
	}
	result := where.Delete(&model.Task{})
	if result.Error != nil {
		return 0, fmt.Errorf("deleting task by code=%s: %v", code, result.Error)
	}
	return result.RowsAffected, nil
}

func (dao *TaskDao) UpdateTask(task *model.Task) (int64, error) {
	result := dao.db.Save(task)
	if result.Error != nil {
		return 0, fmt.Errorf("updating task %v: %v", *task, result.Error)
	}
	return result.RowsAffected, nil
}

func (dao *TaskDao) UpdateTaskEnableByCode(code string, enable bool) (int64, error) {
	task := model.Task{Enable: enable}
	result := dao.db.Model(&task).Select("enable").Where("code = ?", code).Updates(task)
	if result.Error != nil {
		return 0, fmt.Errorf("updating task enable: %v", result.Error)
	}
	return result.RowsAffected, result.Error
}

func (dao *TaskDao) UpdateTaskStatusByCode(code string, status model.TaskStatus, nextTime *time.Time, msg string) (int64, error) {
	return dao.UpdateTaskStatusByCodes(&[]string{code}, status, nextTime, msg)
}

func (dao *TaskDao) UpdateTaskStatusByCodes(codes *[]string, status model.TaskStatus, nextTime *time.Time, msg string) (int64, error) {
	task := model.Task{Status: status}
	if nextTime != nil {
		task.NextTime = *nextTime
	}
	if msg != "" {
		task.Msg = msg
	}
	result := dao.db.Model(&task).Where("code IN ?", *codes).Updates(&task)
	if result.Error != nil {
		return 0, fmt.Errorf("updating task status: %v", result.Error)
	}
	return result.RowsAffected, result.Error
}

func (dao *TaskDao) SelectTaskByID(ID uint) (*model.Task, error) {
	task := model.Task{}
	result := dao.db.First(&task, ID)
	if result.Error != nil {
		return nil, fmt.Errorf("selecting task by id=%d: %v", ID, result.Error)
	}
	return &task, nil
}

func (dao *TaskDao) SelectTaskByCode(code string) (*model.Task, error) {
	task := model.Task{Code: code}
	result := dao.db.Where(&task).First(&task)
	if result.Error != nil {
		return nil, fmt.Errorf("selecting task by code=%s: %v", code, result.Error)
	}
	return &task, nil
}

func (dao *TaskDao) SelectRunnableTasks() (*[]model.Task, error) {
	var tasks []model.Task
	result := dao.db.
		Where("enable = ? AND status = ? AND next_time <= ?",
			true, model.TSReady, time.Now()).
		Find(&tasks)
	if result.Error != nil {
		return nil, fmt.Errorf("selecting runnable tasks: %v", result.Error)
	}
	return &tasks, nil
}
