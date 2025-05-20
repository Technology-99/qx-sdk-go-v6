package akc

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
	KmsAkcService interface {
		KmsAkcCreateKeychain(ctx context.Context, params *qxTypes.KmsAkcCreateKeychainReq) (result *qxTypes.KmsAkcCreateKeychainResp, err error)
		KmsAkcSign(ctx context.Context, params *qxTypes.KmsAkcSignReq) (result *qxTypes.KmsAkcSignResp, err error)
		KmsAkcVerify(ctx context.Context, params *qxTypes.KmsAkcVerifyReq) (result *qxTypes.KmsAkcVerifyResp, err error)
	}

	defaultKmsAkcService struct {
		Svc   string
		qxCtx *qxCtx.QxCtx
	}
)

func NewKmsAkcService(qxCtx *qxCtx.QxCtx) KmsAkcService {
	// note: 初始化Kms系统
	return &defaultKmsAkcService{
		Svc:   "kms",
		qxCtx: qxCtx,
	}
}

func (m *defaultKmsAkcService) KmsAkcCreateKeychain(ctx context.Context, params *qxTypes.KmsAkcCreateKeychainReq) (result *qxTypes.KmsAkcCreateKeychainResp, err error) {
	tmp := &qxRes.BaseResponse[qxTypes.KmsAkcCreateKeychainResp]{}
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, m.Svc, "/kms/akc/createKeychain", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("qx sdk: request error: %v", err)
		return nil, nil
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != qxCodes.QxEngineStatusOK {
		logx.Errorf("qx sdk: KmsAkcCreateKeychain fail: %v", tmp)
		return &tmp.Data, nil
	}
	return &tmp.Data, nil
}

func (m *defaultKmsAkcService) KmsAkcSign(ctx context.Context, params *qxTypes.KmsAkcSignReq) (result *qxTypes.KmsAkcSignResp, err error) {
	tmp := &qxRes.BaseResponse[qxTypes.KmsAkcSignResp]{}
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, m.Svc, "/kms/akc/sign", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("qx sdk: request error: %v", err)
		return nil, nil
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != qxCodes.QxEngineStatusOK {
		logx.Errorf("qx sdk: KmsAkcSign fail: %v", tmp)
		return &tmp.Data, nil
	}
	return &tmp.Data, nil
}

func (m *defaultKmsAkcService) KmsAkcVerify(ctx context.Context, params *qxTypes.KmsAkcVerifyReq) (result *qxTypes.KmsAkcVerifyResp, err error) {
	tmp := &qxRes.BaseResponse[qxTypes.KmsAkcVerifyResp]{}
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, m.Svc, "/kms/akc/verify", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("qx sdk: request error: %v", err)
		return nil, nil
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != qxCodes.QxEngineStatusOK {
		logx.Errorf("qx sdk: KmsAkcVerify fail: %v", tmp)
		return &tmp.Data, nil
	}
	return &tmp.Data, nil
}
