package main

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"qx-sdk-go-v6/qx"
	"qx-sdk-go-v6/qx/qxConfig"
	"qx-sdk-go-v6/qx/qxTypes"
)

func main() {
	qxConfig := qxConfig.DefaultConfig("qx96355103fbb4115d", "46f99d73940ffec05fdb2ccf45b93730")
	//qxConfig.Endpoint = "localhost:8888"
	//qxConfig.Protocol = "http"
	sdk, err := qx.NewQxSdk(qxConfig)
	if err != nil {
		panic(err)
	}
	allCodes, err := sdk.QxBaseService.Codes(context.Background(), &qxTypes.CodesReq{})
	if err != nil {
		panic(err)
	}
	logx.Infof("%+v", allCodes)

	queryBucketResult, err := sdk.SasService.QueryBucket(context.Background(), &qxTypes.SasQueryBucketReq{
		BucketKey: "default",
	})
	if err != nil {
		panic(err)
	}
	logx.Infof("%+v", queryBucketResult)
}
