package main

import (
	"context"
	"github.com/Technology-99/qx-sdk-go-v6/qx"
	"github.com/Technology-99/qx-sdk-go-v6/qx/qxConfig"
	"github.com/Technology-99/qx-sdk-go-v6/qx/qxTypes"
	"github.com/zeromicro/go-zero/core/logx"
	"os"
)

func main() {
	qxConfig := qxConfig.DefaultConfig(os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	logx.Infof("打印一下请求的accessKey :%s", qxConfig.AccessKeyID)

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
