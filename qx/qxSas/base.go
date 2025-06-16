package qxSas

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
	SasBaseService interface {
		QueryBucket(ctx context.Context, params *qxTypes.SasQueryBucketReq) (result *qxTypes.SasQueryBucketResp, err error)
		PresignerGet(ctx context.Context, params *qxTypes.SasPresignerGetObjectReq) (result *qxTypes.SasPresignerGetObjectResp, err error)
		PresignerUpload(ctx context.Context, params *qxTypes.SasPresignerUploadReq) (result *qxTypes.SasPresignerUploadResp, err error)
		PresignerHeadObject(ctx context.Context, params *qxTypes.SasPresignerHeadObjectReq) (result *qxTypes.SasPresignerHeadObjectResp, err error)
		CreateBucketAndConfig(ctx context.Context, params *qxTypes.CreateExistBucketAndConfigReq) (result *qxTypes.CreateExistBucketAndConfigResp, err error)
		CreateBucketNoConfig(ctx context.Context, params *qxTypes.CreateExistBucketNoConfigReq) (result *qxTypes.CreateExistBucketNoConfigResp, err error)
	}
	defaultSasBaseService struct {
		Svc   string
		qxCtx *qxCtx.QxCtx
	}
)

func NewSasBaseService(qxCtx *qxCtx.QxCtx) SasBaseService {
	return &defaultSasBaseService{
		Svc:   "sas",
		qxCtx: qxCtx,
	}
}

// @doc "预签名访问"
func (m *defaultSasBaseService) PresignerGet(ctx context.Context, params *qxTypes.SasPresignerGetObjectReq) (result *qxTypes.SasPresignerGetObjectResp, err error) {
	tmp := &qxRes.BaseResponse[qxTypes.SasPresignerGetObjectResp]{}
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, m.Svc, "/sas/presignerGetObject", http.MethodPost, &params)

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

func (m *defaultSasBaseService) QueryBucket(ctx context.Context, params *qxTypes.SasQueryBucketReq) (result *qxTypes.SasQueryBucketResp, err error) {
	tmp := &qxRes.BaseResponse[qxTypes.SasQueryBucketResp]{}
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, m.Svc, "/sas/queryBucket", http.MethodPost, &params)

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

func (m *defaultSasBaseService) PresignerUpload(ctx context.Context, params *qxTypes.SasPresignerUploadReq) (result *qxTypes.SasPresignerUploadResp, err error) {
	tmp := &qxRes.BaseResponse[qxTypes.SasPresignerUploadResp]{}
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, m.Svc, "/sas/presignerUpload", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("qx sdk: request error: %v", err)
		return nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != qxCodes.QxEngineStatusOK {
		logx.Errorf("qx sdk: request fail: %v", tmp)
		return &tmp.Data, errors.New(tmp.Msg)
	}
	return &tmp.Data, nil
}

func (m *defaultSasBaseService) PresignerHeadObject(ctx context.Context, params *qxTypes.SasPresignerHeadObjectReq) (result *qxTypes.SasPresignerHeadObjectResp, err error) {
	tmp := &qxRes.BaseResponse[qxTypes.SasPresignerHeadObjectResp]{}
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, m.Svc, "/sas/presignerHeadObject", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("qx sdk: request error: %v", err)
		return nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != qxCodes.QxEngineStatusOK {
		logx.Errorf("qx sdk: request fail: %v", tmp)
		return &tmp.Data, errors.New(tmp.Msg)
	}
	return &tmp.Data, nil
}

func (m *defaultSasBaseService) CreateBucketAndConfig(ctx context.Context, params *qxTypes.CreateExistBucketAndConfigReq) (result *qxTypes.CreateExistBucketAndConfigResp, err error) {
	tmp := &qxRes.BaseResponse[qxTypes.CreateExistBucketAndConfigResp]{}
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, m.Svc, "/sas/createExistBucketAndConfig", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("qx sdk: request error: %v", err)
		return nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != qxCodes.QxEngineStatusOK {
		logx.Errorf("qx sdk: request fail: %v", tmp)
		return &tmp.Data, errors.New(tmp.Msg)
	}
	return &tmp.Data, nil
}

func (m *defaultSasBaseService) CreateBucketNoConfig(ctx context.Context, params *qxTypes.CreateExistBucketNoConfigReq) (result *qxTypes.CreateExistBucketNoConfigResp, err error) {
	tmp := &qxRes.BaseResponse[qxTypes.CreateExistBucketNoConfigResp]{}
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, m.Svc, "/sas/createExistBucketNoConfig", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("qx sdk: request error: %v", err)
		return nil, err
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != qxCodes.QxEngineStatusOK {
		logx.Errorf("qx sdk: request fail: %v", tmp)
		return &tmp.Data, errors.New(tmp.Msg)
	}
	return &tmp.Data, nil
}
