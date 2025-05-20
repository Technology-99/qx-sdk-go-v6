package skc

import (
	"context"
	"encoding/json"
	"github.com/Technology-99/qxLib/qxCodes"
	"github.com/Technology-99/qxLib/qxRes"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"qx-sdk-go-v6/qx/qxCtx"
	"qx-sdk-go-v6/qx/qxTypes"
)

type (
	KmsSkcService interface {
		KmsSkcCreateKeychain(ctx context.Context, params *qxTypes.KmsSkcCreateKeychainReq) (result *qxTypes.KmsSkcCreateKeychainResp, err error)
		KmsSkcEncrypt(ctx context.Context, params *qxTypes.KmsSkcEncryptReq) (result *qxTypes.KmsSkcEncryptResp, err error)
		KmsSkcDecrypt(ctx context.Context, params *qxTypes.KmsSkcDecryptReq) (result *qxTypes.KmsSkcDecryptResp, err error)
		KmsSkcBatchEncrypt(ctx context.Context, params *qxTypes.KmsSkcBatchEncryptReq) (result *qxTypes.KmsSkcBatchEncryptResp, err error)
		KmsSkcBatchDecrypt(ctx context.Context, params *qxTypes.KmsSkcBatchDecryptReq) (result *qxTypes.KmsSkcBatchDecryptResp, err error)
		KmsSkcCompare(ctx context.Context, params *qxTypes.KmsSkcCompareReq) (result *qxTypes.KmsSkcCompareResp, err error)
	}

	defaultKmsSkcService struct {
		Svc   string
		qxCtx *qxCtx.QxCtx
	}
)

func NewKmsSkcService(qxCtx *qxCtx.QxCtx) KmsSkcService {
	// note: 初始化Kms系统
	return &defaultKmsSkcService{
		Svc:   "kms",
		qxCtx: qxCtx,
	}
}

func (m *defaultKmsSkcService) KmsSkcCreateKeychain(ctx context.Context, params *qxTypes.KmsSkcCreateKeychainReq) (result *qxTypes.KmsSkcCreateKeychainResp, err error) {
	tmp := &qxRes.BaseResponse[qxTypes.KmsSkcCreateKeychainResp]{}
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, m.Svc, "/kms/skc/createKeychain", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("qx sdk: request error: %v", err)
		return nil, nil
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != qxCodes.QxEngineStatusOK {
		logx.Errorf("qx sdk: KmsSkcCreateKeychain fail: %v", tmp)
		return &tmp.Data, nil
	}
	return &tmp.Data, nil
}

func (m *defaultKmsSkcService) KmsSkcEncrypt(ctx context.Context, params *qxTypes.KmsSkcEncryptReq) (result *qxTypes.KmsSkcEncryptResp, err error) {
	tmp := &qxRes.BaseResponse[qxTypes.KmsSkcEncryptResp]{}
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, m.Svc, "/kms/skc/encrypt", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("qx sdk: request error: %v", err)
		return nil, nil
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != qxCodes.QxEngineStatusOK {
		logx.Errorf("qx sdk: KmsSkcCreateKeychain fail: %v", tmp)
		return &tmp.Data, nil
	}
	return &tmp.Data, nil
}

func (m *defaultKmsSkcService) KmsSkcDecrypt(ctx context.Context, params *qxTypes.KmsSkcDecryptReq) (result *qxTypes.KmsSkcDecryptResp, err error) {
	tmp := &qxRes.BaseResponse[qxTypes.KmsSkcDecryptResp]{}
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, m.Svc, "/kms/skc/decrypt", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("qx sdk: request error: %v", err)
		return nil, nil
	}
	_ = json.Unmarshal(res, &result)
	if tmp.Code != qxCodes.QxEngineStatusOK {
		logx.Errorf("qx sdk: KmsSkcDecrypt fail: %v", tmp)
		return &tmp.Data, nil
	}
	return &tmp.Data, nil
}

func (m *defaultKmsSkcService) KmsSkcBatchEncrypt(ctx context.Context, params *qxTypes.KmsSkcBatchEncryptReq) (result *qxTypes.KmsSkcBatchEncryptResp, err error) {
	tmp := &qxRes.BaseResponse[qxTypes.KmsSkcBatchEncryptResp]{}
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, m.Svc, "/kms/skc/batchEncrypt", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("qx sdk: request error: %v", err)
		return nil, nil
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != qxCodes.QxEngineStatusOK {
		logx.Errorf("qx sdk: KmsSkcCreateKeychain fail: %v", tmp)
		return &tmp.Data, nil
	}
	return &tmp.Data, nil
}

func (m *defaultKmsSkcService) KmsSkcBatchDecrypt(ctx context.Context, params *qxTypes.KmsSkcBatchDecryptReq) (result *qxTypes.KmsSkcBatchDecryptResp, err error) {
	tmp := &qxRes.BaseResponse[qxTypes.KmsSkcBatchDecryptResp]{}
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, m.Svc, "/kms/skc/batchDecrypt", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("qx sdk: request error: %v", err)
		return nil, nil
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != qxCodes.QxEngineStatusOK {
		logx.Errorf("qx sdk: KmsSkcCreateKeychain fail: %v", tmp)
		return &tmp.Data, nil
	}
	return &tmp.Data, nil
}

func (m *defaultKmsSkcService) KmsSkcCompare(ctx context.Context, params *qxTypes.KmsSkcCompareReq) (result *qxTypes.KmsSkcCompareResp, err error) {
	tmp := &qxRes.BaseResponse[qxTypes.KmsSkcCompareResp]{}
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, m.Svc, "/kms/skc/compare", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("qx sdk: request error: %v", err)
		return nil, nil
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != qxCodes.QxEngineStatusOK {
		logx.Errorf("qx sdk: KmsSkcCreateKeychain fail: %v", tmp)
		return &tmp.Data, nil
	}
	return &tmp.Data, nil
}
