package qxCtx

import (
	"context"
	"github.com/Technology-99/qx-sdk-go-v6/qx/qxClient"
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
