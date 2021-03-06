// Code generated by Kitex v0.0.8. DO NOT EDIT.

package indexservice

import (
	"alert/kitex_gen/api"
	"alert/kitex_gen/rpc_dto"
	"context"
	"github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return indexServiceServiceInfo
}

var indexServiceServiceInfo = newServiceInfo()

func newServiceInfo() *kitex.ServiceInfo {
	serviceName := "IndexService"
	handlerType := (*api.IndexService)(nil)
	methods := map[string]kitex.MethodInfo{
		"SelectIndex":     kitex.NewMethodInfo(selectIndexHandler, newIndexServiceSelectIndexArgs, newIndexServiceSelectIndexResult, false),
		"SelectAllIndex":  kitex.NewMethodInfo(selectAllIndexHandler, newIndexServiceSelectAllIndexArgs, newIndexServiceSelectAllIndexResult, false),
		"AddIndex":        kitex.NewMethodInfo(addIndexHandler, newIndexServiceAddIndexArgs, newIndexServiceAddIndexResult, false),
		"DeleteIndex":     kitex.NewMethodInfo(deleteIndexHandler, newIndexServiceDeleteIndexArgs, newIndexServiceDeleteIndexResult, false),
		"UpdateIndex":     kitex.NewMethodInfo(updateIndexHandler, newIndexServiceUpdateIndexArgs, newIndexServiceUpdateIndexResult, false),
		"SelectRoomIndex": kitex.NewMethodInfo(selectRoomIndexHandler, newIndexServiceSelectRoomIndexArgs, newIndexServiceSelectRoomIndexResult, false),
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

func selectIndexHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*api.IndexServiceSelectIndexArgs)
	realResult := result.(*api.IndexServiceSelectIndexResult)
	success, err := handler.(api.IndexService).SelectIndex(ctx, realArg.Code)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newIndexServiceSelectIndexArgs() interface{} {
	return api.NewIndexServiceSelectIndexArgs()
}

func newIndexServiceSelectIndexResult() interface{} {
	return api.NewIndexServiceSelectIndexResult()
}

func selectAllIndexHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {

	realResult := result.(*api.IndexServiceSelectAllIndexResult)
	success, err := handler.(api.IndexService).SelectAllIndex(ctx)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newIndexServiceSelectAllIndexArgs() interface{} {
	return api.NewIndexServiceSelectAllIndexArgs()
}

func newIndexServiceSelectAllIndexResult() interface{} {
	return api.NewIndexServiceSelectAllIndexResult()
}

func addIndexHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*api.IndexServiceAddIndexArgs)
	realResult := result.(*api.IndexServiceAddIndexResult)
	success, err := handler.(api.IndexService).AddIndex(ctx, realArg.Name, realArg.Code, realArg.Content, realArg.TimeRange)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newIndexServiceAddIndexArgs() interface{} {
	return api.NewIndexServiceAddIndexArgs()
}

func newIndexServiceAddIndexResult() interface{} {
	return api.NewIndexServiceAddIndexResult()
}

func deleteIndexHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*api.IndexServiceDeleteIndexArgs)

	err := handler.(api.IndexService).DeleteIndex(ctx, realArg.Code)
	if err != nil {
		return err
	}

	return nil
}
func newIndexServiceDeleteIndexArgs() interface{} {
	return api.NewIndexServiceDeleteIndexArgs()
}

func newIndexServiceDeleteIndexResult() interface{} {
	return api.NewIndexServiceDeleteIndexResult()
}

func updateIndexHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*api.IndexServiceUpdateIndexArgs)

	err := handler.(api.IndexService).UpdateIndex(ctx, realArg.Index)
	if err != nil {
		return err
	}

	return nil
}
func newIndexServiceUpdateIndexArgs() interface{} {
	return api.NewIndexServiceUpdateIndexArgs()
}

func newIndexServiceUpdateIndexResult() interface{} {
	return api.NewIndexServiceUpdateIndexResult()
}

func selectRoomIndexHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*api.IndexServiceSelectRoomIndexArgs)
	realResult := result.(*api.IndexServiceSelectRoomIndexResult)
	success, err := handler.(api.IndexService).SelectRoomIndex(ctx, realArg.Code, realArg.RoomId)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newIndexServiceSelectRoomIndexArgs() interface{} {
	return api.NewIndexServiceSelectRoomIndexArgs()
}

func newIndexServiceSelectRoomIndexResult() interface{} {
	return api.NewIndexServiceSelectRoomIndexResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) SelectIndex(ctx context.Context, code string) (r *rpc_dto.IndexResponse, err error) {
	var _args api.IndexServiceSelectIndexArgs
	_args.Code = code
	var _result api.IndexServiceSelectIndexResult
	if err = p.c.Call(ctx, "SelectIndex", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) SelectAllIndex(ctx context.Context) (r *rpc_dto.IndexsResponse, err error) {
	var _args api.IndexServiceSelectAllIndexArgs
	var _result api.IndexServiceSelectAllIndexResult
	if err = p.c.Call(ctx, "SelectAllIndex", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AddIndex(ctx context.Context, name string, code string, content string, timeRange int64) (r *rpc_dto.IndexResponse, err error) {
	var _args api.IndexServiceAddIndexArgs
	_args.Name = name
	_args.Code = code
	_args.Content = content
	_args.TimeRange = timeRange
	var _result api.IndexServiceAddIndexResult
	if err = p.c.Call(ctx, "AddIndex", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeleteIndex(ctx context.Context, code string) (err error) {
	var _args api.IndexServiceDeleteIndexArgs
	_args.Code = code
	var _result api.IndexServiceDeleteIndexResult
	if err = p.c.Call(ctx, "DeleteIndex", &_args, &_result); err != nil {
		return
	}
	return nil
}

func (p *kClient) UpdateIndex(ctx context.Context, index *rpc_dto.Index) (err error) {
	var _args api.IndexServiceUpdateIndexArgs
	_args.Index = index
	var _result api.IndexServiceUpdateIndexResult
	if err = p.c.Call(ctx, "UpdateIndex", &_args, &_result); err != nil {
		return
	}
	return nil
}

func (p *kClient) SelectRoomIndex(ctx context.Context, code []string, roomId int64) (r *rpc_dto.MapIndexResponse, err error) {
	var _args api.IndexServiceSelectRoomIndexArgs
	_args.Code = code
	_args.RoomId = roomId
	var _result api.IndexServiceSelectRoomIndexResult
	if err = p.c.Call(ctx, "SelectRoomIndex", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
