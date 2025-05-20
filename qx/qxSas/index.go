package qxSas

import (
	"github.com/Technology-99/qx-sdk-go-v6/qx/qxCtx"
)

type (
	SasService struct {
		SasBaseService
	}
)

func NewSasService(qxCtx *qxCtx.QxCtx) SasService {
	return SasService{
		SasBaseService: NewSasBaseService(qxCtx),
	}
}
