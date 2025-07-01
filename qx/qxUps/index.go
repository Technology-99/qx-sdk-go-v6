package qxUps

import (
	"github.com/Technology-99/qx-sdk-go-v6/qx/qxCtx"
)

type (
	UpsService struct {
		UpsBaseService
		UpsTagService       UpsTagService
		UpsIndustryService  UpsIndustryService
		UpsShortLinkService UpsShortLinkService
	}
)

func NewUpsService(qxCtx *qxCtx.QxCtx) UpsService {
	return UpsService{
		UpsBaseService:      NewUpsBaseService(qxCtx),
		UpsTagService:       NewUpsTagService(qxCtx),
		UpsIndustryService:  NewUpsIndustryService(qxCtx),
		UpsShortLinkService: NewUpsShortLinkService(qxCtx),
	}
}
