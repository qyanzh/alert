/**
 * @author  qyanzh
 * @create  2022/03/07 18:41
 */

package dao

import (
	"alert/internal/model"
	"testing"
	"time"
)

var taskDao *TaskDao

func init() {
	taskDao = NewTaskDao()
}

func TestAddTask(t *testing.T) {
	code := "test_task5"
	name := "测试任务5"
	_, _ = taskDao.DeleteTaskByCode(code, true)
	task := model.Task{}
	task.Code = code
	task.Name = name
	task.RuleID = 0
	task.Frequency = 15
	task.Enable = true
	task.Status = model.TSRunning
	task.NextTime = time.Now()
	_, err := taskDao.AddTask(&task)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", task)
}

func TestTaskDao_SelectRunnableTask(t *testing.T) {
	tasks, err := taskDao.SelectRunnableTasks()
	if err != nil {
		t.Error(err)
	}
	for _, task := range *tasks {
		t.Logf("%+v\n", task)
	}
}

func TestTaskDao_UpdateTaskStatusByCodes(t *testing.T) {
	codes := []string{"test_task",
		"test_task2",
		"test_task3",
	}
	_, err := taskDao.UpdateTaskStatusByCodes(&codes, model.TSReady, nil, "")
	if err != nil {
		t.Error(err)
	}
}

func TestTaskDao_UpdateTaskEnablesByCode(t *testing.T) {
	code := "test_task"
	_, err := taskDao.UpdateTaskEnableByCode(code, false)
	if err != nil {
		t.Error(err)
	}
}

func TestTaskDao_UpdateTaskStatusAndNextTimeByCode(t *testing.T) {
	code := "test_task"
	nextTime := time.Now().Add(time.Hour)
	_, err := taskDao.UpdateTaskStatusByCodes(&[]string{code}, model.TSReady, &nextTime, "test")
	if err != nil {
		t.Error(err)
	}
}

func TestTaskDao_UpdateTask(t *testing.T) {
	task, err := taskDao.SelectTaskByCode("test_task")
	if err != nil {
		t.Error(err)
	}
	task.Status = 1
	task.Enable = true
	_, err = taskDao.UpdateTask(task)
	if err != nil {
		t.Error(err)
	}
}
