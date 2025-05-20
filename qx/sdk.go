package qx

import (
	"github.com/Technology-99/qx-sdk-go-v6/qx/qxBase"
	"github.com/Technology-99/qx-sdk-go-v6/qx/qxClient"
	"github.com/Technology-99/qx-sdk-go-v6/qx/qxConfig"
	"github.com/Technology-99/qx-sdk-go-v6/qx/qxCtx"
	"github.com/Technology-99/qx-sdk-go-v6/qx/qxKms"
	"github.com/Technology-99/qx-sdk-go-v6/qx/qxMas"
	"github.com/Technology-99/qx-sdk-go-v6/qx/qxSas"
)

type QxSdk struct {
	config        *qxConfig.Config
	client        *qxClient.QxClient
	qxCtx         *qxCtx.QxCtx
	QxBaseService qxBase.QxBaseService
	KmsService    qxKms.KmsService
	MasService    qxMas.MasService
	SasService    qxSas.SasService
}

func NewQxSdk(c *qxConfig.Config) (*QxSdk, error) {
	// note: sdk初始化检查
	if err := c.Check(); err != nil {
		return nil, err
	}
	// note: sdk初始化
	client := qxClient.NewQxClient(c)
	qxC := qxCtx.NewQxCtx(client)

	sdk := &QxSdk{
		config:        c,
		client:        client,
		qxCtx:         qxC,
		QxBaseService: qxBase.NewQxBaseService(qxC),
		KmsService:    qxKms.NewKmsService(qxC),
		MasService:    qxMas.NewMasService(qxC),
		SasService:    qxSas.NewSasService(qxC),
	}
	return sdk, nil
}
