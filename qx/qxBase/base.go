package qxBase

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Technology-99/qxLib/qxCodes"
	"github.com/Technology-99/qxLib/qxRes"
	"github.com/google/go-querystring/query"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"qx-sdk-go-v6/qx/qxCtx"
	"qx-sdk-go-v6/qx/qxTypes"
)

type (
	QxBaseService interface {
		Codes(ctx context.Context, params *qxTypes.CodesReq) (result *qxTypes.CodesResp, err error)
		Zones(ctx context.Context, params *qxTypes.ZonesReq) (result *qxTypes.ZonesResp, err error)
	}
	defaultQxBaseService struct {
		Svc   string
		qxCtx *qxCtx.QxCtx
	}
)

func NewQxBaseService(qxCtx *qxCtx.QxCtx) QxBaseService {
	return &defaultQxBaseService{
		Svc:   "base",
		qxCtx: qxCtx,
	}
}

func (m *defaultQxBaseService) Codes(ctx context.Context, params *qxTypes.CodesReq) (result *qxTypes.CodesResp, err error) {
	tmp := &qxRes.BaseResponse[qxTypes.CodesResp]{}
	relativePath := ""
	if params.Lang != "" || params.Svc != "" {
		v, _ := query.Values(params)
		relativePath = fmt.Sprintf("/base/codes?%s", v.Encode())
	} else {
		relativePath = "/base/codes"
	}
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, m.Svc, relativePath, http.MethodGet, nil)
	if err != nil {
		logx.Errorf("qx sdk: request error: %v", err)
		return nil, nil
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != qxCodes.QxEngineStatusOK {
		logx.Errorf("qx sdk: codes fail: %v", tmp)
		return &tmp.Data, nil
	}
	return &tmp.Data, nil
}

func (m *defaultQxBaseService) Zones(ctx context.Context, params *qxTypes.ZonesReq) (result *qxTypes.ZonesResp, err error) {
	tmp := &qxRes.BaseResponse[qxTypes.ZonesResp]{}
	relativePath := ""
	if params.Lang != "" {
		v, _ := query.Values(params)
		relativePath = fmt.Sprintf("/base/zones?%s", v.Encode())
	} else {
		relativePath = "/base/zones"
	}
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, m.Svc, relativePath, http.MethodGet, nil)
	if err != nil {
		logx.Errorf("qx sdk: request error: %v", err)
		return nil, nil
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != qxCodes.QxEngineStatusOK {
		logx.Errorf("qx sdk: zones fail: %v", tmp)
		return &tmp.Data, nil
	}
	return &tmp.Data, nil
}
