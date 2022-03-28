package main

import (
	"alert/internal/model"
	"alert/internal/service"
	"alert/kitex_gen/rpc_dto"
	"context"
	"encoding/json"
	"errors"
	T "time"
)

var alertService service.AlertService
var indexService service.IndexService
var ruleService service.RuleService
var taskService service.TaskService

func init() {
	alertService = *service.NewAlertService()
	indexService = *service.NewIndexService()
	ruleService = *service.NewRuleService()
	taskService = *service.NewTaskService()
}

// CombineServiceImpl implements the last service interface defined in the IDL.
type CombineServiceImpl struct{}

func convertIndexToDto(index model.Index) *rpc_dto.Index {
	dtoIndex := rpc_dto.Index{
		Id:        int64(index.ID),
		Code:      index.Code,
		Name:      index.Name,
		Expr:      index.Expr,
		TimeRange: int64(index.TimeRange),
	}
	if index.Type == model.ITNormal {
		dtoIndex.Type = true
	} else {
		dtoIndex.Type = false
	}
	return &dtoIndex
}
func convertRuleToDto(rule model.Rule) *rpc_dto.Rule {
	dtoRule := rpc_dto.Rule{
		Id:     int64(rule.Id),
		Code:   rule.Code,
		RoomId: int64(rule.RoomId),
		Name:   rule.Name,
		Expr:   rule.Expr,
	}
	if rule.Type == model.NORMALRULE {
		dtoRule.Type = true
	} else {
		dtoRule.Type = false
	}
	return &dtoRule
}
func convertTaskToDto(task model.Task) *rpc_dto.Task {

	dtoTask := rpc_dto.Task{
		Id:         int64(task.ID),
		Code:       task.Code,
		Name:       task.Name,
		RuleId:     int64(task.RuleID),
		Frequency:  int64(task.Frequency),
		Enable:     task.Enable,
		NextTime:   task.NextTime.Format(T.ANSIC),
		LastStatus: int16(task.Status),
	}
	return &dtoTask
}

// AddAlert implements the AlertServiceImpl interface.
func (s *CombineServiceImpl) AddAlert(ctx context.Context, ruleId int64, time string, indexNum map[int64]float64) (err error) {
	// TODO: Your code here...
	if err != nil {
		return
	}
	t, err := T.Parse(T.ANSIC, time)
	if err != nil {
		return
	}
	var indexNumConvert map[uint]float64
	for v1, v2 := range indexNum {
		indexNumConvert[uint(v1)] = v2
	}
	err = alertService.AddAlert(uint(ruleId), t, indexNumConvert)
	return
}

// SelectAlert implements the AlertServiceImpl interface.
func (s *CombineServiceImpl) SelectAlert(ctx context.Context, ruleId int64, startTime string, endTime string) (resp *rpc_dto.AlertsResponse, err error) {
	// TODO: Your code here...
	if err != nil {
		return
	}
	st, err := T.Parse(T.ANSIC, startTime)
	if err != nil {
		return
	}
	ed, err := T.Parse(T.ANSIC, endTime)
	if err != nil {
		return
	}
	tmpAlert, err := alertService.SelectAlert(uint(ruleId), st, ed)
	if err != nil {
		return
	}
	for _, v := range *tmpAlert {
		var indexNum map[uint]float64
		err = json.Unmarshal(v.IndexNum, &indexNum)
		if err != nil {
			return
		}
		var indexNumConvert map[int64]float64
		for v1, v2 := range indexNum {
			indexNumConvert[int64(v1)] = v2
		}
		a := rpc_dto.Alert{Id: int64(v.Id), Time: v.Time.Format(T.ANSIC), IndexNum: indexNumConvert}
		resp.Alerts = append(resp.Alerts, &a)
	}
	return
}

// SelectIndex implements the IndexServiceImpl interface.
func (s *CombineServiceImpl) SelectIndex(ctx context.Context, code string) (resp *rpc_dto.IndexResponse, err error) {
	// TODO: Your code here...
	index, err := indexService.SelectIndexByCode(code)
	if err != nil {
		return
	}
	resp.Index = convertIndexToDto(*index)
	return
}

// SelectAllIndex implements the IndexServiceImpl interface.
func (s *CombineServiceImpl) SelectAllIndex(ctx context.Context) (resp *rpc_dto.IndexsResponse, err error) {
	// TODO: Your code here...
	indexs, err := indexService.SelectAllIndices()
	if err != nil {
		return
	}
	for _, v := range *indexs {
		resp.Indexs = append(resp.Indexs, convertIndexToDto(v))
	}
	return
}

// AddIndex implements the IndexServiceImpl interface.
func (s *CombineServiceImpl) AddIndex(ctx context.Context, name string, code string, content string, timeRange int64) (resp *rpc_dto.IndexResponse, err error) {
	// TODO: Your code here...
	index, err := indexService.AddIndex(name, code, content, uint(timeRange))
	if err != nil {
		return
	}
	resp.Index = convertIndexToDto(*index)
	return
}

// DeleteIndex implements the IndexServiceImpl interface.
func (s *CombineServiceImpl) DeleteIndex(ctx context.Context, code string) (err error) {
	// TODO: Your code here...
	err = indexService.DeleteIndex(code)
	return
}

// UpdateIndex implements the IndexServiceImpl interface.
func (s *CombineServiceImpl) UpdateIndex(ctx context.Context, index *rpc_dto.Index) (err error) {
	// TODO: Your code here...
	modelIndex, err := indexService.SelectIndexByID(uint(index.Id))
	if err != nil {
		return err
	}
	if modelIndex.Code != index.Code {
		return errors.New("id与code不匹配")
	}
	modelIndex.Name = index.Name
	if modelIndex.Expr != index.Expr {
		modelIndex.Expr = index.Expr
		modelIndex.Type, modelIndex.Serialized, err = indexService.EvaluatorIndex(modelIndex)
		if err != nil {
			return
		}
	}
	modelIndex.TimeRange = uint(index.TimeRange)
	err = indexService.UpdateIndex(modelIndex)
	return
}

// SelectRoomIndex implements the IndexServiceImpl interface.
func (s *CombineServiceImpl) SelectRoomIndex(ctx context.Context, code []string, roomId int64) (resp *rpc_dto.MapIndexResponse, err error) {
	// TODO: Your code here...
	roomIndex, err := indexService.SelectIndexValuesByCodesAndRoomID(code, uint(roomId))
	resp.Indexs = roomIndex
	return
}

// SelectRule implements the RuleServiceImpl interface.
func (s *CombineServiceImpl) SelectRule(ctx context.Context, code string) (resp *rpc_dto.RuleResponse, err error) {
	// TODO: Your code here...
	rule, err := ruleService.SelectRule(code)
	if err != nil {
		return
	}
	resp.Rule = convertRuleToDto(*rule)
	return
}

// SelectAllRule implements the RuleServiceImpl interface.
func (s *CombineServiceImpl) SelectAllRule(ctx context.Context) (resp *rpc_dto.RulesResponse, err error) {
	// TODO: Your code here...
	rules, err := ruleService.SelectAllRules()
	if err != nil {
		return
	}
	for _, v := range *rules {
		resp.Rules = append(resp.Rules, convertRuleToDto(v))
	}
	return
}

// AddRule implements the RuleServiceImpl interface.
func (s *CombineServiceImpl) AddRule(ctx context.Context, roomId int64, name string, code string, ruleType bool, content string) (resp *rpc_dto.RuleResponse, err error) {
	// TODO: Your code here...
	rule, err := ruleService.AddRule(uint(roomId), name, code, ruleType, content)
	if err != nil {
		return
	}
	resp.Rule = convertRuleToDto(*rule)
	return
}

// CheckRule implements the RuleServiceImpl interface.
func (s *CombineServiceImpl) CheckRule(ctx context.Context, rule string) (resp *rpc_dto.CheckResponse, err error) {
	// TODO: Your code here...
	r, _, err := ruleService.CheckRule(rule)
	resp.Result_ = r
	return
}

// DeleteRule implements the RuleServiceImpl interface.
func (s *CombineServiceImpl) DeleteRule(ctx context.Context, code string) (err error) {
	// TODO: Your code here...
	err = ruleService.DeleteRule(code)
	return
}

// UpdateRule implements the RuleServiceImpl interface.
func (s *CombineServiceImpl) UpdateRule(ctx context.Context, rule *rpc_dto.Rule) (err error) {
	// TODO: Your code here...
	modelRule, err := ruleService.SelectRuleById(uint(rule.Id))
	if err != nil {
		return err
	}
	if modelRule.Code != rule.Code {
		return errors.New("id与code不匹配")
	}
	modelRule.Name = rule.Name
	modelRule.RoomId = uint(rule.RoomId)
	if modelRule.Expr != rule.Expr {
		if rule.Type {
			modelRule.Type = model.NORMALRULE
		} else {
			modelRule.Type = model.COMPLEXRULE
		}
		modelRule.Expr = rule.Expr
		modelRule.Serialized, err = ruleService.EvaluatorRule(modelRule)
		if err != nil {
			return
		}
	}
	err = ruleService.UpdateRule(modelRule)
	return
}

// SelectTask implements the TaskServiceImpl interface.
func (s *CombineServiceImpl) SelectTask(ctx context.Context, code string) (resp *rpc_dto.TaskResponse, err error) {
	// TODO: Your code here...
	task, err := taskService.SelectTaskByCode(code)
	if err != nil {
		return
	}
	resp.Task = convertTaskToDto(*task)
	return
}

// AddTask implements the TaskServiceImpl interface.
func (s *CombineServiceImpl) AddTask(ctx context.Context, name string, code string, ruleId int64, frequency int64) (resp *rpc_dto.TaskResponse, err error) {
	// TODO: Your code here...
	task, err := taskService.AddTask(code, name, uint(ruleId), uint(frequency))
	if err != nil {
		return
	}
	resp.Task = convertTaskToDto(*task)
	return
}

// DeleteTask implements the TaskServiceImpl interface.
func (s *CombineServiceImpl) DeleteTask(ctx context.Context, code string) (err error) {
	// TODO: Your code here...
	err = taskService.DeleteTaskByCode(code)
	return
}

func (s *CombineServiceImpl) UpdateTaskEnable(ctx context.Context, code string, enable bool) (err error) {
	err = taskService.UpdateTaskEnableByCode(code, enable)
	return
}

func (s *CombineServiceImpl) SelectRunnableTask(ctx context.Context) (resp *rpc_dto.TasksResponse, err error) {
	tasks, err := taskService.SelectRunnableTasks()
	if err != nil {
		return
	}
	for _, v := range *tasks {
		resp.Tasks = append(resp.Tasks, convertTaskToDto(v))
	}
	return
}

// UpdateTask implements the TaskServiceImpl interface.
func (s *CombineServiceImpl) UpdateTask(ctx context.Context, task *rpc_dto.Task) (err error) {
	// TODO: Your code here...
	modelTask, err := taskService.SelectTaskByID(uint(task.Id))
	if err != nil {
		return
	}
	if modelTask.Code != task.Code {
		return errors.New("id与code不匹配")
	}
	modelTask.Name = task.Name
	modelTask.Enable = task.Enable
	modelTask.Frequency = uint(task.Frequency)
	modelTask.RuleID = uint(task.RuleId)
	err = taskService.UpdateTask(modelTask)
	return
}
