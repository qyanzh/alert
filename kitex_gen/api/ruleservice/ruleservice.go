// Code generated by Kitex v0.0.8. DO NOT EDIT.

package ruleservice

import (
	"alert/kitex_gen/api"
	"alert/kitex_gen/rpc_dto"
	"context"
	"github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return ruleServiceServiceInfo
}

var ruleServiceServiceInfo = newServiceInfo()

func newServiceInfo() *kitex.ServiceInfo {
	serviceName := "RuleService"
	handlerType := (*api.RuleService)(nil)
	methods := map[string]kitex.MethodInfo{
		"SelectRule":    kitex.NewMethodInfo(selectRuleHandler, newRuleServiceSelectRuleArgs, newRuleServiceSelectRuleResult, false),
		"SelectAllRule": kitex.NewMethodInfo(selectAllRuleHandler, newRuleServiceSelectAllRuleArgs, newRuleServiceSelectAllRuleResult, false),
		"AddRule":       kitex.NewMethodInfo(addRuleHandler, newRuleServiceAddRuleArgs, newRuleServiceAddRuleResult, false),
		"DeleteRule":    kitex.NewMethodInfo(deleteRuleHandler, newRuleServiceDeleteRuleArgs, newRuleServiceDeleteRuleResult, false),
		"UpdateRule":    kitex.NewMethodInfo(updateRuleHandler, newRuleServiceUpdateRuleArgs, newRuleServiceUpdateRuleResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "api",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.0.8",
		Extra:           extra,
	}
	return svcInfo
}

func selectRuleHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*api.RuleServiceSelectRuleArgs)
	realResult := result.(*api.RuleServiceSelectRuleResult)
	success, err := handler.(api.RuleService).SelectRule(ctx, realArg.Code)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newRuleServiceSelectRuleArgs() interface{} {
	return api.NewRuleServiceSelectRuleArgs()
}

func newRuleServiceSelectRuleResult() interface{} {
	return api.NewRuleServiceSelectRuleResult()
}

func selectAllRuleHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {

	realResult := result.(*api.RuleServiceSelectAllRuleResult)
	success, err := handler.(api.RuleService).SelectAllRule(ctx)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newRuleServiceSelectAllRuleArgs() interface{} {
	return api.NewRuleServiceSelectAllRuleArgs()
}

func newRuleServiceSelectAllRuleResult() interface{} {
	return api.NewRuleServiceSelectAllRuleResult()
}

func addRuleHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*api.RuleServiceAddRuleArgs)
	realResult := result.(*api.RuleServiceAddRuleResult)
	success, err := handler.(api.RuleService).AddRule(ctx, realArg.RoomId, realArg.Name, realArg.Code, realArg.RuleType, realArg.Content)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newRuleServiceAddRuleArgs() interface{} {
	return api.NewRuleServiceAddRuleArgs()
}

func newRuleServiceAddRuleResult() interface{} {
	return api.NewRuleServiceAddRuleResult()
}

func deleteRuleHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*api.RuleServiceDeleteRuleArgs)
	realResult := result.(*api.RuleServiceDeleteRuleResult)
	success, err := handler.(api.RuleService).DeleteRule(ctx, realArg.Code)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newRuleServiceDeleteRuleArgs() interface{} {
	return api.NewRuleServiceDeleteRuleArgs()
}

func newRuleServiceDeleteRuleResult() interface{} {
	return api.NewRuleServiceDeleteRuleResult()
}

func updateRuleHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*api.RuleServiceUpdateRuleArgs)
	realResult := result.(*api.RuleServiceUpdateRuleResult)
	success, err := handler.(api.RuleService).UpdateRule(ctx, realArg.Rule)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newRuleServiceUpdateRuleArgs() interface{} {
	return api.NewRuleServiceUpdateRuleArgs()
}

func newRuleServiceUpdateRuleResult() interface{} {
	return api.NewRuleServiceUpdateRuleResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) SelectRule(ctx context.Context, code string) (r *rpc_dto.RuleResponse, err error) {
	var _args api.RuleServiceSelectRuleArgs
	_args.Code = code
	var _result api.RuleServiceSelectRuleResult
	if err = p.c.Call(ctx, "SelectRule", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) SelectAllRule(ctx context.Context) (r *rpc_dto.RulesResponse, err error) {
	var _args api.RuleServiceSelectAllRuleArgs
	var _result api.RuleServiceSelectAllRuleResult
	if err = p.c.Call(ctx, "SelectAllRule", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AddRule(ctx context.Context, roomId int64, name string, code string, ruleType bool, content string) (r *rpc_dto.RuleResponse, err error) {
	var _args api.RuleServiceAddRuleArgs
	_args.RoomId = roomId
	_args.Name = name
	_args.Code = code
	_args.RuleType = ruleType
	_args.Content = content
	var _result api.RuleServiceAddRuleResult
	if err = p.c.Call(ctx, "AddRule", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeleteRule(ctx context.Context, code string) (r *rpc_dto.ErrResponse, err error) {
	var _args api.RuleServiceDeleteRuleArgs
	_args.Code = code
	var _result api.RuleServiceDeleteRuleResult
	if err = p.c.Call(ctx, "DeleteRule", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateRule(ctx context.Context, rule *rpc_dto.Rule) (r *rpc_dto.ErrResponse, err error) {
	var _args api.RuleServiceUpdateRuleArgs
	_args.Rule = rule
	var _result api.RuleServiceUpdateRuleResult
	if err = p.c.Call(ctx, "UpdateRule", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
