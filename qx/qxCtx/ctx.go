package qxCtx

import (
	"context"
	"qx-sdk-go-v6/qx/qxClient"
)

type QxCtx struct {
	Context context.Context
	Cancel  context.CancelFunc
	Cli     *qxClient.QxClient
}

func NewQxCtx(client *qxClient.QxClient) *QxCtx {
	ctx, cancel := context.WithCancel(context.Background())
	return &QxCtx{
		Context: ctx,
		Cancel:  cancel,
		Cli:     client,
	}
}
