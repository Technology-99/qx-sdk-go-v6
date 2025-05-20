package qxSas

import (
	"context"
	"encoding/json"
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
		PresignerUpload(ctx context.Context, params *qxTypes.SasPresignerUploadReq) (result *qxTypes.SasPresignerUploadResp, err error)
		PresignerHeadObject(ctx context.Context, params *qxTypes.SasPresignerHeadObjectReq) (result *qxTypes.SasPresignerHeadObjectResp, err error)
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

func (m *defaultSasBaseService) QueryBucket(ctx context.Context, params *qxTypes.SasQueryBucketReq) (result *qxTypes.SasQueryBucketResp, err error) {
	tmp := &qxRes.BaseResponse[qxTypes.SasQueryBucketResp]{}
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, m.Svc, "/sas/queryBucket", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("qx sdk: request error: %v", err)
		return nil, nil
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != qxCodes.QxEngineStatusOK {
		logx.Errorf("qx sdk: request fail: %s", res)
		return &tmp.Data, nil
	}
	return &tmp.Data, nil
}

func (m *defaultSasBaseService) PresignerUpload(ctx context.Context, params *qxTypes.SasPresignerUploadReq) (result *qxTypes.SasPresignerUploadResp, err error) {
	tmp := &qxRes.BaseResponse[qxTypes.SasPresignerUploadResp]{}
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, m.Svc, "/sas/presignerUpload", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("qx sdk: request error: %v", err)
		return nil, nil
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != qxCodes.QxEngineStatusOK {
		logx.Errorf("qx sdk: request fail: %v", tmp)
		return &tmp.Data, nil
	}
	return &tmp.Data, nil
}

func (m *defaultSasBaseService) PresignerHeadObject(ctx context.Context, params *qxTypes.SasPresignerHeadObjectReq) (result *qxTypes.SasPresignerHeadObjectResp, err error) {
	tmp := &qxRes.BaseResponse[qxTypes.SasPresignerHeadObjectResp]{}
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, m.Svc, "/sas/presignerHeadObject", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("qx sdk: request error: %v", err)
		return nil, nil
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != qxCodes.QxEngineStatusOK {
		logx.Errorf("qx sdk: request fail: %v", tmp)
		return &tmp.Data, nil
	}
	return &tmp.Data, nil
}
