package qxUps

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/Technology-99/qx-sdk-go-v6/qx/qxCtx"
	"github.com/Technology-99/qx-sdk-go-v6/qx/qxTypes"
	"github.com/Technology-99/qxLib/qxCodes"
	"github.com/Technology-99/qxLib/qxRes"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
)

type (
	UpsBaseService interface {
		Bootstrap(ctx context.Context, params *qxTypes.UpsBaseBootstrapReq) (result *qxTypes.UpsBaseBootstrapResp, err error)
	}
	defaultUpsBaseService struct {
		Svc   string
		qxCtx *qxCtx.QxCtx
	}
)

func NewUpsBaseService(qxCtx *qxCtx.QxCtx) UpsBaseService {
	return &defaultUpsBaseService{
		Svc:   "ups",
		qxCtx: qxCtx,
	}
}

func (m *defaultUpsBaseService) Bootstrap(ctx context.Context, params *qxTypes.UpsBaseBootstrapReq) (result *qxTypes.UpsBaseBootstrapResp, err error) {
	tmp := &qxRes.BaseResponse[qxTypes.UpsBaseBootstrapResp]{}
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, m.Svc, "/ups/base/bootstrap", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("qx sdk: request error: %v", err)
		return nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != qxCodes.QxEngineStatusOK {
		logx.Errorf("qx sdk: request fail: %s", res)
		return &tmp.Data, errors.New(tmp.Msg)
	}
	return &tmp.Data, nil
}
