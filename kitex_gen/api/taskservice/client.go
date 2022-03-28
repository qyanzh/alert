// Code generated by Kitex v0.0.8. DO NOT EDIT.

package taskservice

import (
	"alert/kitex_gen/rpc_dto"
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	SelectTask(ctx context.Context, code string, callOptions ...callopt.Option) (r *rpc_dto.TaskResponse, err error)
	AddTask(ctx context.Context, name string, code string, ruleId int64, frequency int64, callOptions ...callopt.Option) (r *rpc_dto.TaskResponse, err error)
	UpdateTaskEnable(ctx context.Context, code string, enable bool, callOptions ...callopt.Option) (err error)
	SelectRunnableTask(ctx context.Context, callOptions ...callopt.Option) (r *rpc_dto.TasksResponse, err error)
	DeleteTask(ctx context.Context, code string, callOptions ...callopt.Option) (err error)
	UpdateTask(ctx context.Context, task *rpc_dto.Task, callOptions ...callopt.Option) (err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kTaskServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kTaskServiceClient struct {
	*kClient
}

func (p *kTaskServiceClient) SelectTask(ctx context.Context, code string, callOptions ...callopt.Option) (r *rpc_dto.TaskResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SelectTask(ctx, code)
}

func (p *kTaskServiceClient) AddTask(ctx context.Context, name string, code string, ruleId int64, frequency int64, callOptions ...callopt.Option) (r *rpc_dto.TaskResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AddTask(ctx, name, code, ruleId, frequency)
}

func (p *kTaskServiceClient) UpdateTaskEnable(ctx context.Context, code string, enable bool, callOptions ...callopt.Option) (err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateTaskEnable(ctx, code, enable)
}

func (p *kTaskServiceClient) SelectRunnableTask(ctx context.Context, callOptions ...callopt.Option) (r *rpc_dto.TasksResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SelectRunnableTask(ctx)
}

func (p *kTaskServiceClient) DeleteTask(ctx context.Context, code string, callOptions ...callopt.Option) (err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteTask(ctx, code)
}

func (p *kTaskServiceClient) UpdateTask(ctx context.Context, task *rpc_dto.Task, callOptions ...callopt.Option) (err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateTask(ctx, task)
}