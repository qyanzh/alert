// Code generated by Kitex v0.0.8. DO NOT EDIT.

package alertservice

import (
	"alert/kitex_gen/rpc_dto"
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	AddAlert(ctx context.Context, ruleId int64, time string, indexNum map[int64]float64, callOptions ...callopt.Option) (err error)
	SelectAlert(ctx context.Context, ruleId int64, startTime string, endTime string, callOptions ...callopt.Option) (r *rpc_dto.AlertsResponse, err error)
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
	return &kAlertServiceClient{
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

type kAlertServiceClient struct {
	*kClient
}

func (p *kAlertServiceClient) AddAlert(ctx context.Context, ruleId int64, time string, indexNum map[int64]float64, callOptions ...callopt.Option) (err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AddAlert(ctx, ruleId, time, indexNum)
}

func (p *kAlertServiceClient) SelectAlert(ctx context.Context, ruleId int64, startTime string, endTime string, callOptions ...callopt.Option) (r *rpc_dto.AlertsResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SelectAlert(ctx, ruleId, startTime, endTime)
}
