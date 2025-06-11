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
	UpsTagService interface {
		Create(ctx context.Context, params *qxTypes.AllowCreateModelTag) (result *qxTypes.TagApiCreateResp, err error)
		Delete(ctx context.Context, params *qxTypes.TagApiFormIdReq) (result *qxTypes.TagApiOKResp, err error)
		DeleteMany(ctx context.Context, params *qxTypes.TagApiJsonIdsReq) (result *qxTypes.TagApiOKResp, err error)
		Update(ctx context.Context, params *qxTypes.AllowUpdateModelTag) (result *qxTypes.TagApiOKResp, err error)
		UpdateStatus(ctx context.Context, params *qxTypes.AllowUpdateStatusModelTag) (result *qxTypes.TagApiOKResp, err error)
		Query(ctx context.Context, params *qxTypes.TagApiFormIdReq) (result *qxTypes.ModelTag, err error)
		QueryListWhereIds(ctx context.Context, params *qxTypes.TagApiJsonIdsReq) (result *qxTypes.TagCommonQueryListResp, err error)
		QueryList(ctx context.Context, params *qxTypes.TagCommonSearchParams) (result *qxTypes.TagCommonQueryListResp, err error)
	}
	defaultUpsTagService struct {
		Svc   string
		qxCtx *qxCtx.QxCtx
	}
)

func NewUpsTagService(qxCtx *qxCtx.QxCtx) UpsTagService {
	return &defaultUpsTagService{
		Svc:   "ups",
		qxCtx: qxCtx,
	}
}

func (m *defaultUpsTagService) Create(ctx context.Context, params *qxTypes.AllowCreateModelTag) (result *qxTypes.TagApiCreateResp, err error) {
	tmp := &qxRes.BaseResponse[qxTypes.TagApiCreateResp]{}
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, m.Svc, "/ups/tag/create", http.MethodPost, &params)

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

func (m *defaultUpsTagService) Delete(ctx context.Context, params *qxTypes.TagApiFormIdReq) (result *qxTypes.TagApiOKResp, err error) {
	tmp := &qxRes.BaseResponse[qxTypes.TagApiOKResp]{}
	relativePath := "/ups/tag/delete"
	if params.Id != 0 {
		relativePath = fmt.Sprintf("/ups/tag/query?id=%d", params.Id)
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

func (m *defaultUpsTagService) DeleteMany(ctx context.Context, params *qxTypes.TagApiJsonIdsReq) (result *qxTypes.TagApiOKResp, err error) {
	tmp := &qxRes.BaseResponse[qxTypes.TagApiOKResp]{}
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, m.Svc, "/ups/tag/deleteMany", http.MethodPost, &params)

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

func (m *defaultUpsTagService) Update(ctx context.Context, params *qxTypes.AllowUpdateModelTag) (result *qxTypes.TagApiOKResp, err error) {
	tmp := &qxRes.BaseResponse[qxTypes.TagApiOKResp]{}
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, m.Svc, "/ups/tag/update", http.MethodPost, &params)

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

func (m *defaultUpsTagService) UpdateStatus(ctx context.Context, params *qxTypes.AllowUpdateStatusModelTag) (result *qxTypes.TagApiOKResp, err error) {
	tmp := &qxRes.BaseResponse[qxTypes.TagApiOKResp]{}
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, m.Svc, "/ups/tag/updateStatus", http.MethodPost, &params)

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

func (m *defaultUpsTagService) Query(ctx context.Context, params *qxTypes.TagApiFormIdReq) (result *qxTypes.ModelTag, err error) {
	tmp := &qxRes.BaseResponse[qxTypes.ModelTag]{}
	relativePath := "/ups/tag/query"
	if params.Id != 0 {
		relativePath = fmt.Sprintf("/ups/tag/query?id=%d", params.Id)
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

func (m *defaultUpsTagService) QueryListWhereIds(ctx context.Context, params *qxTypes.TagApiJsonIdsReq) (result *qxTypes.TagCommonQueryListResp, err error) {
	tmp := &qxRes.BaseResponse[qxTypes.TagCommonQueryListResp]{}
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, m.Svc, "/ups/tag/queryListWhereIds", http.MethodPost, &params)

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

func (m *defaultUpsTagService) QueryList(ctx context.Context, params *qxTypes.TagCommonSearchParams) (result *qxTypes.TagCommonQueryListResp, err error) {
	tmp := &qxRes.BaseResponse[qxTypes.TagCommonQueryListResp]{}
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, m.Svc, "/ups/tag/queryList", http.MethodPost, &params)

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
