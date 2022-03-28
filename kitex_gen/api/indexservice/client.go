// Code generated by Kitex v0.0.8. DO NOT EDIT.

package indexservice

import (
	"alert/kitex_gen/rpc_dto"
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	SelectIndex(ctx context.Context, code string, callOptions ...callopt.Option) (r *rpc_dto.IndexResponse, err error)
	SelectAllIndex(ctx context.Context, callOptions ...callopt.Option) (r *rpc_dto.IndexsResponse, err error)
	AddIndex(ctx context.Context, name string, code string, content string, timeRange int64, callOptions ...callopt.Option) (r *rpc_dto.IndexResponse, err error)
	DeleteIndex(ctx context.Context, code string, callOptions ...callopt.Option) (err error)
	UpdateIndex(ctx context.Context, index *rpc_dto.Index, callOptions ...callopt.Option) (err error)
	SelectRoomIndex(ctx context.Context, code []string, roomId int64, callOptions ...callopt.Option) (r *rpc_dto.MapIndexResponse, err error)
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
	return &kIndexServiceClient{
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

type kIndexServiceClient struct {
	*kClient
}

func (p *kIndexServiceClient) SelectIndex(ctx context.Context, code string, callOptions ...callopt.Option) (r *rpc_dto.IndexResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SelectIndex(ctx, code)
}

func (p *kIndexServiceClient) SelectAllIndex(ctx context.Context, callOptions ...callopt.Option) (r *rpc_dto.IndexsResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SelectAllIndex(ctx)
}

func (p *kIndexServiceClient) AddIndex(ctx context.Context, name string, code string, content string, timeRange int64, callOptions ...callopt.Option) (r *rpc_dto.IndexResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AddIndex(ctx, name, code, content, timeRange)
}

func (p *kIndexServiceClient) DeleteIndex(ctx context.Context, code string, callOptions ...callopt.Option) (err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteIndex(ctx, code)
}

func (p *kIndexServiceClient) UpdateIndex(ctx context.Context, index *rpc_dto.Index, callOptions ...callopt.Option) (err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateIndex(ctx, index)
}

func (p *kIndexServiceClient) SelectRoomIndex(ctx context.Context, code []string, roomId int64, callOptions ...callopt.Option) (r *rpc_dto.MapIndexResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SelectRoomIndex(ctx, code, roomId)
}