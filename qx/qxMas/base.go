package qxMas

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
	MasBaseService interface {
		// note: 生成验证码
		CaptchaGenerate(ctx context.Context, params *qxTypes.ApiCaptchaGenerateReq) (result *qxTypes.ApiCaptchaGenerateResp, err error)
		SmsSend(ctx context.Context, params *qxTypes.ApiSmsSendReq) (result *qxTypes.ApiSmsSendResp, err error)

		BehavioralVerificationInit(ctx context.Context, params *qxTypes.BehavioralVerificationInitReq) (result *qxTypes.BehavioralVerificationInitResp, err error)
		BehavioralVerificationVerify(ctx context.Context, params *qxTypes.BehavioralVerificationVerifyReq) (result *qxTypes.BehavioralVerificationVerifyResp, err error)
		SmsVerificationInit(ctx context.Context, params *qxTypes.SmsInitReq) (result *qxTypes.SmsInitResp, err error)
		SmsVerificationVerify(ctx context.Context, params *qxTypes.SmsVerifyReq) (result *qxTypes.SmsVerifyResp, err error)
	}

	defaultMasBaseService struct {
		Svc   string
		qxCtx *qxCtx.QxCtx
	}
)

func NewMsgBaseService(qxCtx *qxCtx.QxCtx) MasBaseService {
	return &defaultMasBaseService{
		Svc:   "mas",
		qxCtx: qxCtx,
	}
}

func (m *defaultMasBaseService) CaptchaGenerate(ctx context.Context, params *qxTypes.ApiCaptchaGenerateReq) (result *qxTypes.ApiCaptchaGenerateResp, err error) {
	tmp := &qxRes.BaseResponse[qxTypes.ApiCaptchaGenerateResp]{}
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, m.Svc, "/mas/captcha/generate", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("qx sdk: request error: %v", err)
		return nil, nil
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != qxCodes.QxEngineStatusOK {
		logx.Errorf("qx sdk: captcha generate fail: %v", tmp)
		return &tmp.Data, nil
	}
	return &tmp.Data, nil
}

func (m *defaultMasBaseService) SmsSend(ctx context.Context, params *qxTypes.ApiSmsSendReq) (result *qxTypes.ApiSmsSendResp, err error) {
	tmp := &qxRes.BaseResponse[qxTypes.ApiSmsSendResp]{}
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, m.Svc, "/mas/sms/send", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("qx sdk: request error: %v", err)
		return nil, nil
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != qxCodes.QxEngineStatusOK {
		logx.Errorf("qx sdk: sms send fail: %v", tmp)
		return &tmp.Data, nil
	}
	return &tmp.Data, nil
}

func (m *defaultMasBaseService) BehavioralVerificationInit(ctx context.Context, params *qxTypes.BehavioralVerificationInitReq) (result *qxTypes.BehavioralVerificationInitResp, err error) {
	tmp := &qxRes.BaseResponse[qxTypes.BehavioralVerificationInitResp]{}
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, m.Svc, "/mas/bv/init", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("qx sdk: request error: %v", err)
		return nil, nil
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != qxCodes.QxEngineStatusOK {
		logx.Errorf("qx sdk:Behavioral Verification Init fail: %v", result)
		return &tmp.Data, nil
	}
	return &tmp.Data, nil
}

func (m *defaultMasBaseService) BehavioralVerificationVerify(ctx context.Context, params *qxTypes.BehavioralVerificationVerifyReq) (result *qxTypes.BehavioralVerificationVerifyResp, err error) {
	tmp := &qxRes.BaseResponse[qxTypes.BehavioralVerificationVerifyResp]{}
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, m.Svc, "/mas/bv/verify", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("qx sdk: request error: %v", err)
		return nil, nil
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != qxCodes.QxEngineStatusOK {
		logx.Errorf("qx sdk: Behavioral Verification Verify fail: %v", tmp)
		return &tmp.Data, nil
	}
	return &tmp.Data, nil
}

func (m *defaultMasBaseService) SmsVerificationInit(ctx context.Context, params *qxTypes.SmsInitReq) (result *qxTypes.SmsInitResp, err error) {
	tmp := &qxRes.BaseResponse[qxTypes.SmsInitResp]{}
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, m.Svc, "/mas/sms/init", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("qx sdk: request error: %v", err)
		return nil, nil
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != qxCodes.QxEngineStatusOK {
		logx.Errorf("qx sdk: sms init fail: %v", tmp)
		return &tmp.Data, nil
	}
	return &tmp.Data, nil
}

func (m *defaultMasBaseService) SmsVerificationVerify(ctx context.Context, params *qxTypes.SmsVerifyReq) (result *qxTypes.SmsVerifyResp, err error) {
	tmp := &qxRes.BaseResponse[qxTypes.SmsVerifyResp]{}
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, m.Svc, "/mas/sms/verify", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("qx sdk: request error: %v", err)
		return nil, nil
	}
	_ = json.Unmarshal(res, &tmp)
	if tmp.Code != qxCodes.QxEngineStatusOK {
		logx.Errorf("qx sdk: sms verify fail: %v", tmp)
		return &tmp.Data, nil
	}
	return &tmp.Data, nil
}
