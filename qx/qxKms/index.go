package qxKms

import (
	"qx-sdk-go-v6/qx/qxCtx"
	"qx-sdk-go-v6/qx/qxKms/akc"
	"qx-sdk-go-v6/qx/qxKms/skc"
)

type (
	KmsService struct {
		Akc akc.KmsAkcService
		Skc skc.KmsSkcService
	}
)

func NewKmsService(qxCtx *qxCtx.QxCtx) KmsService {
	return KmsService{
		Akc: akc.NewKmsAkcService(qxCtx),
		Skc: skc.NewKmsSkcService(qxCtx),
	}
}
