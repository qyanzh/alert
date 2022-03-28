package main

import (
	"alert/kitex_gen/api"
	"context"
	"testing"
	"time"
)

func TestAddAlert(t *testing.T) {
	client := api.AlertServiceClient{}
	err := client.AddAlert(context.Background(), 0, time.Now().Format(time.ANSIC), map[int64]float64{1: 100, 2: 200})
	if err != nil {
		t.Error(err)
	}
}
