package qxKms

import (
	"github.com/Technology-99/qx-sdk-go-v6/qx/qxCtx"
	"github.com/Technology-99/qx-sdk-go-v6/qx/qxKms/akc"
	"github.com/Technology-99/qx-sdk-go-v6/qx/qxKms/skc"
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
