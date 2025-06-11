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
	UpsIndustryService interface {
		Create(ctx context.Context, params *qxTypes.AllowCreateModelIndustry) (result *qxTypes.IndustryApiCreateResp, err error)
		Delete(ctx context.Context, params *qxTypes.IndustryApiFormIdReq) (result *qxTypes.IndustryApiOKResp, err error)
		DeleteMany(ctx context.Context, params *qxTypes.IndustryApiJsonIdsReq) (result *qxTypes.IndustryApiOKResp, err error)
		Update(ctx context.Context, params *qxTypes.AllowUpdateModelIndustry) (result *qxTypes.IndustryApiOKResp, err error)
		UpdateStatus(ctx context.Context, params *qxTypes.AllowUpdateStatusModelIndustry) (result *qxTypes.IndustryApiOKResp, err error)
		Query(ctx context.Context, params *qxTypes.IndustryApiFormIdReq) (result *qxTypes.ModelIndustry, err error)
		QueryListWhereIds(ctx context.Context, params *qxTypes.IndustryApiJsonIdsReq) (result *qxTypes.IndustryCommonQueryListResp, err error)
		QueryList(ctx context.Context, params *qxTypes.IndustryCommonSearchParams) (result *qxTypes.IndustryCommonQueryListResp, err error)
	}
	defaultUpsIndustryService struct {
		Svc   string
		qxCtx *qxCtx.QxCtx
	}
)

func NewUpsIndustryService(qxCtx *qxCtx.QxCtx) UpsIndustryService {
	return &defaultUpsIndustryService{
		Svc:   "ups",
		qxCtx: qxCtx,
	}
}

func (m *defaultUpsIndustryService) Create(ctx context.Context, params *qxTypes.AllowCreateModelIndustry) (result *qxTypes.IndustryApiCreateResp, err error) {
	tmp := &qxRes.BaseResponse[qxTypes.IndustryApiCreateResp]{}
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, m.Svc, "/ups/industry/create", http.MethodPost, &params)

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

func (m *defaultUpsIndustryService) Delete(ctx context.Context, params *qxTypes.IndustryApiFormIdReq) (result *qxTypes.IndustryApiOKResp, err error) {
	tmp := &qxRes.BaseResponse[qxTypes.IndustryApiOKResp]{}
	relativePath := "/ups/industry/delete"
	if params.Id != 0 {
		relativePath = fmt.Sprintf("/ups/industry/query?id=%d", params.Id)
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

func (m *defaultUpsIndustryService) DeleteMany(ctx context.Context, params *qxTypes.IndustryApiJsonIdsReq) (result *qxTypes.IndustryApiOKResp, err error) {
	tmp := &qxRes.BaseResponse[qxTypes.IndustryApiOKResp]{}
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, m.Svc, "/ups/industry/deleteMany", http.MethodPost, &params)

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

func (m *defaultUpsIndustryService) Update(ctx context.Context, params *qxTypes.AllowUpdateModelIndustry) (result *qxTypes.IndustryApiOKResp, err error) {
	tmp := &qxRes.BaseResponse[qxTypes.IndustryApiOKResp]{}
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, m.Svc, "/ups/industry/update", http.MethodPost, &params)

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

func (m *defaultUpsIndustryService) UpdateStatus(ctx context.Context, params *qxTypes.AllowUpdateStatusModelIndustry) (result *qxTypes.IndustryApiOKResp, err error) {
	tmp := &qxRes.BaseResponse[qxTypes.IndustryApiOKResp]{}
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, m.Svc, "/ups/industry/updateStatus", http.MethodPost, &params)

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

func (m *defaultUpsIndustryService) Query(ctx context.Context, params *qxTypes.IndustryApiFormIdReq) (result *qxTypes.ModelIndustry, err error) {
	tmp := &qxRes.BaseResponse[qxTypes.ModelIndustry]{}
	relativePath := "/ups/industry/query"
	if params.Id != 0 {
		relativePath = fmt.Sprintf("/ups/industry/query?id=%d", params.Id)
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

func (m *defaultUpsIndustryService) QueryListWhereIds(ctx context.Context, params *qxTypes.IndustryApiJsonIdsReq) (result *qxTypes.IndustryCommonQueryListResp, err error) {
	tmp := &qxRes.BaseResponse[qxTypes.IndustryCommonQueryListResp]{}
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, m.Svc, "/ups/industry/queryListWhereIds", http.MethodPost, &params)

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

func (m *defaultUpsIndustryService) QueryList(ctx context.Context, params *qxTypes.IndustryCommonSearchParams) (result *qxTypes.IndustryCommonQueryListResp, err error) {
	tmp := &qxRes.BaseResponse[qxTypes.IndustryCommonQueryListResp]{}
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, m.Svc, "/ups/industry/queryList", http.MethodPost, &params)

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
