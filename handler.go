package main

import (
	"alert/kitex_gen/rpc_dto"
	"context"
)

// CombineServiceImpl implements the last service interface defined in the IDL.
type CombineServiceImpl struct{}

// AddAlert implements the AlertServiceImpl interface.
func (s *CombineServiceImpl) AddAlert(ctx context.Context, ruleCode string, time string, roomId int64) (resp *rpc_dto.ErrResponse, err error) {
	// TODO: Your code here...
	return
}

// SelectAlert implements the AlertServiceImpl interface.
func (s *CombineServiceImpl) SelectAlert(ctx context.Context, roomId int64, ruleCode string, startTime string, endTime string) (resp *rpc_dto.AlertsResponse, err error) {
	// TODO: Your code here...
	alert := &rpc_dto.Alert{}
	resp.Alerts = []*rpc_dto.Alert{alert}
	return
}

// SelectIndex implements the IndexServiceImpl interface.
func (s *CombineServiceImpl) SelectIndex(ctx context.Context, code string) (resp *rpc_dto.IndexResponse, err error) {
	// TODO: Your code here...
	return
}

// SelectAllIndex implements the IndexServiceImpl interface.
func (s *CombineServiceImpl) SelectAllIndex(ctx context.Context) (resp *rpc_dto.IndexsResponse, err error) {
	// TODO: Your code here...
	return
}

// AddIndex implements the IndexServiceImpl interface.
func (s *CombineServiceImpl) AddIndex(ctx context.Context, name string, code string, indexType bool, content string) (resp *rpc_dto.IndexResponse, err error) {
	// TODO: Your code here...
	return
}

// DeleteIndex implements the IndexServiceImpl interface.
func (s *CombineServiceImpl) DeleteIndex(ctx context.Context, code string) (resp *rpc_dto.ErrResponse, err error) {
	// TODO: Your code here...
	return
}

// UpdateIndex implements the IndexServiceImpl interface.
func (s *CombineServiceImpl) UpdateIndex(ctx context.Context, index *rpc_dto.Index) (resp *rpc_dto.ErrResponse, err error) {
	// TODO: Your code here...
	return
}

// SelectRoomIndex implements the IndexServiceImpl interface.
func (s *CombineServiceImpl) SelectRoomIndex(ctx context.Context, code []string, roomId int64) (resp *rpc_dto.MapIndexResponse, err error) {
	// TODO: Your code here...
	return
}

// SelectRule implements the RuleServiceImpl interface.
func (s *CombineServiceImpl) SelectRule(ctx context.Context, code string) (resp *rpc_dto.RuleResponse, err error) {
	// TODO: Your code here...
	return
}

// SelectAllRule implements the RuleServiceImpl interface.
func (s *CombineServiceImpl) SelectAllRule(ctx context.Context) (resp *rpc_dto.RulesResponse, err error) {
	// TODO: Your code here...
	return
}

// AddRule implements the RuleServiceImpl interface.
func (s *CombineServiceImpl) AddRule(ctx context.Context, roomId int64, name string, code string, ruleType bool, content string) (resp *rpc_dto.RuleResponse, err error) {
	// TODO: Your code here...
	return
}

// DeleteRule implements the RuleServiceImpl interface.
func (s *CombineServiceImpl) DeleteRule(ctx context.Context, code string) (resp *rpc_dto.ErrResponse, err error) {
	// TODO: Your code here...
	return
}

// UpdateRule implements the RuleServiceImpl interface.
func (s *CombineServiceImpl) UpdateRule(ctx context.Context, rule *rpc_dto.Rule) (resp *rpc_dto.ErrResponse, err error) {
	// TODO: Your code here...
	return
}

// SelectTask implements the TaskServiceImpl interface.
func (s *CombineServiceImpl) SelectTask(ctx context.Context, code string) (resp *rpc_dto.TaskResponse, err error) {
	// TODO: Your code here...
	return
}

// AddTask implements the TaskServiceImpl interface.
func (s *CombineServiceImpl) AddTask(ctx context.Context, name string, code string, ruleCode string, frequency int64) (resp *rpc_dto.TaskResponse, err error) {
	// TODO: Your code here...
	return
}

// SelectRoomTask implements the TaskServiceImpl interface.
func (s *CombineServiceImpl) SelectRoomTask(ctx context.Context, roomId int64) (resp *rpc_dto.TasksResponse, err error) {
	// TODO: Your code here...
	return
}

// DeleteTask implements the TaskServiceImpl interface.
func (s *CombineServiceImpl) DeleteTask(ctx context.Context, code string) (resp *rpc_dto.ErrResponse, err error) {
	// TODO: Your code here...
	return
}

// UpdateTask implements the TaskServiceImpl interface.
func (s *CombineServiceImpl) UpdateTask(ctx context.Context, task *rpc_dto.Task) (resp *rpc_dto.ErrResponse, err error) {
	// TODO: Your code here...
	return
}
