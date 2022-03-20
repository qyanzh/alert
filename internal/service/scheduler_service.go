package service

import (
	"alert/internal/model"
	"alert/internal/scheduler/gpool"
	"fmt"
	"log"
	"time"
)

type SchedulerService struct {
	taskService    *TaskService
	ruleService    *RuleService
	alertService   *AlertService
	runnableTasks  chan model.Task
	ruleCheckGPool *gpool.GPool
	done           chan struct{}
}

func NewSchedulerService() *SchedulerService {
	return &SchedulerService{
		taskService:    NewTaskService(),
		ruleService:    NewRuleService(),
		alertService:   NewAlertService(),
		runnableTasks:  make(chan model.Task),
		ruleCheckGPool: gpool.New(),
	}
}

func (ss *SchedulerService) Start() error {
	if ss.done != nil {
		return fmt.Errorf("scheduler already started")
	}
	ss.done = make(chan struct{})
	go ss.queryRunnableTasks(5 * time.Second)
	go ss.checkTaskRules()
	return nil
}

func (ss *SchedulerService) Stop() error {
	close(ss.done)
	return ss.ruleCheckGPool.Close()
}

func (ss *SchedulerService) queryRunnableTasks(delay time.Duration) {
	for {
		select {
		case <-ss.done:
			return
		default:
			tasks, err := ss.taskService.SelectRunnableTasks()
			if err != nil {
				log.Println(err)
			}
			for _, task := range *tasks {
				select {
				case <-ss.done:
					return
				case ss.runnableTasks <- task:
				}
			}
			time.Sleep(delay)
		}
	}
}

func (ss *SchedulerService) checkTaskRules() {
	for {
		select {
		case <-ss.done:
			return
		case task := <-ss.runnableTasks:
			err := ss.ruleCheckGPool.Enqueue(func() {
				err := ss.doCheckTaskRule(task)
				if err != nil {
					log.Println(err)
				}
			})
			if err != nil {
				log.Println(err)
			}
		}
	}
}

func (ss *SchedulerService) doCheckTaskRule(task model.Task) error {
	if err := ss.taskService.UpdateTaskStatusRunningByCode(task.Code); err != nil {
		return err
	}
	isSatisfy, indexMap, err := ss.ruleService.CheckRuleWithId(task.RuleID)
	if err != nil {
		if err := ss.taskService.UpdateTaskFailByCode(task.Code, err.Error()); err != nil {
			return err
		}
		return err
	}
	if !isSatisfy {
		if err := ss.alertService.AddAlert(task.RuleID, time.Now(), indexMap); err != nil {
			return err
		}
	}

	nextTime := time.Now().Add(time.Duration(task.Frequency) * time.Second)
	if err := ss.taskService.UpdateTaskSuccessByCode(task.Code, &nextTime); err != nil {
		return err
	}
	return nil
}
