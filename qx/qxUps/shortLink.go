package qxUps

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Technology-99/qx-sdk-go-v6/qx/qxCtx"
	"github.com/Technology-99/qx-sdk-go-v6/qx/qxTypes"
	"github.com/Technology-99/qxLib/qxCodes"
	"github.com/Technology-99/qxLib/qxRes"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
)

type (
	UpsShortLinkService interface {
		GetRedirectResult(ctx context.Context, params *qxTypes.GetRedirectResultReq) (result *qxTypes.GetRedirectResultResp, err error)
		Create(ctx context.Context, params *qxTypes.AllowCreateModelShortLink) (result *qxTypes.ShortLinkApiCreateResp, err error)
		Delete(ctx context.Context, params *qxTypes.ShortLinkApiFormIdReq) (result *qxTypes.ShortLinkApiOKResp, err error)
		DeleteMany(ctx context.Context, params *qxTypes.ShortLinkApiJsonIdsReq) (result *qxTypes.ShortLinkApiOKResp, err error)
		Update(ctx context.Context, params *qxTypes.AllowUpdateModelShortLink) (result *qxTypes.ShortLinkApiOKResp, err error)
		UpdateStatus(ctx context.Context, params *qxTypes.AllowUpdateStatusModelShortLink) (result *qxTypes.ShortLinkApiOKResp, err error)
		Query(ctx context.Context, params *qxTypes.ShortLinkApiFormIdReq) (result *qxTypes.ModelShortLink, err error)
		QueryWhereKey(ctx context.Context, params *qxTypes.ShortLinkApiFormKeyReq) (result *qxTypes.ModelShortLink, err error)
		QueryListWhereIds(ctx context.Context, params *qxTypes.ShortLinkApiJsonIdsReq) (result *qxTypes.ShortLinkCommonQueryListResp, err error)
		QueryList(ctx context.Context, params *qxTypes.ShortLinkCommonSearchParams) (result *qxTypes.ShortLinkCommonQueryListResp, err error)
	}
	defaultUpsShortLinkService struct {
		Svc   string
		qxCtx *qxCtx.QxCtx
	}
)

func NewUpsShortLinkService(qxCtx *qxCtx.QxCtx) UpsShortLinkService {
	return &defaultUpsShortLinkService{
		Svc:   "ups",
		qxCtx: qxCtx,
	}
}

func (m *defaultUpsShortLinkService) GetRedirectResult(ctx context.Context, params *qxTypes.GetRedirectResultReq) (result *qxTypes.GetRedirectResultResp, err error) {
	tmp := &qxRes.BaseResponse[qxTypes.GetRedirectResultResp]{}
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, m.Svc, "/ups/shortLink/getRedirectResult", http.MethodPost, &params)

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

func (m *defaultUpsShortLinkService) Create(ctx context.Context, params *qxTypes.AllowCreateModelShortLink) (result *qxTypes.ShortLinkApiCreateResp, err error) {
	tmp := &qxRes.BaseResponse[qxTypes.ShortLinkApiCreateResp]{}
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, m.Svc, "/ups/shortLink/create", http.MethodPost, &params)

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

func (m *defaultUpsShortLinkService) Delete(ctx context.Context, params *qxTypes.ShortLinkApiFormIdReq) (result *qxTypes.ShortLinkApiOKResp, err error) {
	tmp := &qxRes.BaseResponse[qxTypes.ShortLinkApiOKResp]{}
	relativePath := "/ups/shortLink/query"
	if params.Id != 0 {
		relativePath = fmt.Sprintf("/ups/shortLink/delete?id=%d", params.Id)
	}
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, m.Svc, relativePath, http.MethodPost, &params)

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

func (m *defaultUpsShortLinkService) DeleteMany(ctx context.Context, params *qxTypes.ShortLinkApiJsonIdsReq) (result *qxTypes.ShortLinkApiOKResp, err error) {
	tmp := &qxRes.BaseResponse[qxTypes.ShortLinkApiOKResp]{}
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, m.Svc, "/ups/shortLink/deleteMany", http.MethodPost, &params)

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

func (m *defaultUpsShortLinkService) Update(ctx context.Context, params *qxTypes.AllowUpdateModelShortLink) (result *qxTypes.ShortLinkApiOKResp, err error) {
	tmp := &qxRes.BaseResponse[qxTypes.ShortLinkApiOKResp]{}
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, m.Svc, "/ups/shortLink/update", http.MethodPost, &params)

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

func (m *defaultUpsShortLinkService) UpdateStatus(ctx context.Context, params *qxTypes.AllowUpdateStatusModelShortLink) (result *qxTypes.ShortLinkApiOKResp, err error) {
	tmp := &qxRes.BaseResponse[qxTypes.ShortLinkApiOKResp]{}
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, m.Svc, "/ups/shortLink/updateStatus", http.MethodPost, &params)

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

func (m *defaultUpsShortLinkService) Query(ctx context.Context, params *qxTypes.ShortLinkApiFormIdReq) (result *qxTypes.ModelShortLink, err error) {
	tmp := &qxRes.BaseResponse[qxTypes.ModelShortLink]{}
	relativePath := "/ups/shortLink/query"
	if params.Id != 0 {
		relativePath = fmt.Sprintf("/ups/shortLink/query?id=%d", params.Id)
	}
	logx.Infof("qx sdk: request path: %s", relativePath)
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, m.Svc, relativePath, http.MethodGet, &params)
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

func (m *defaultUpsShortLinkService) QueryWhereKey(ctx context.Context, params *qxTypes.ShortLinkApiFormKeyReq) (result *qxTypes.ModelShortLink, err error) {
	tmp := &qxRes.BaseResponse[qxTypes.ModelShortLink]{}
	relativePath := "/ups/shortLink/queryWhereKey"
	if params.Key != "" {
		relativePath = fmt.Sprintf("/ups/shortLink/queryWhereKey?key=%d", params.Key)
	}
	logx.Infof("qx sdk: request path: %s", relativePath)
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, m.Svc, relativePath, http.MethodGet, &params)
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

func (m *defaultUpsShortLinkService) QueryListWhereIds(ctx context.Context, params *qxTypes.ShortLinkApiJsonIdsReq) (result *qxTypes.ShortLinkCommonQueryListResp, err error) {
	tmp := &qxRes.BaseResponse[qxTypes.ShortLinkCommonQueryListResp]{}
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, m.Svc, "/ups/shortLink/queryListWhereIds", http.MethodPost, &params)

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

func (m *defaultUpsShortLinkService) QueryList(ctx context.Context, params *qxTypes.ShortLinkCommonSearchParams) (result *qxTypes.ShortLinkCommonQueryListResp, err error) {
	tmp := &qxRes.BaseResponse[qxTypes.ShortLinkCommonQueryListResp]{}
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, m.Svc, "/ups/shortLink/queryList", http.MethodPost, &params)

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
