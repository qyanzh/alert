package main

import (
	"alert/kitex_gen/api/combineservice"
	"context"
	"github.com/cloudwego/kitex/client"
	"testing"
	"time"
)

var c combineservice.Client

func init() {
	c, _ = combineservice.NewClient("alert", client.WithHostPorts("0.0.0.0:9999"))
}

func TestAddAlert(t *testing.T) {
	err := c.AddAlert(context.Background(), 0, time.Now().Format(time.ANSIC), map[int64]float64{1: 100, 2: 200})
	if err != nil {
		t.Error(err)
	}
}

func TestAddIndex(t *testing.T) {
	now := time.Now()
	resp, err := c.AddIndex(context.Background(), "rpc测试指标"+now.String(), "rpc测试指标"+now.String(), "raw[count(*)]", 10)
	t.Log(resp.Index)
	if err != nil {
		t.Error(err)
	}
}
