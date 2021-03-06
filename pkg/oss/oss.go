package oss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"go.uber.org/zap"
	"io"
	"pinkacg/settings"
)

func OssUpload(key string, file io.Reader) (string, error) {
	client, err := oss.New(settings.Conf.Endpoint, settings.Conf.AccessKeyId, settings.Conf.AccessKeySecret)
	if err != nil {
		zap.L().Error("OssUpload failed", zap.Error(err))
		return "", err
	}
	// 获取存储空间。
	bucket, err := client.Bucket(settings.Conf.BucketName)
	if err != nil {
		zap.L().Error("client.Bucket failed", zap.Error(err))
		return "", err
	}
	// 上传文件。
	err = bucket.PutObject(key, file)
	if err != nil {
		zap.L().Error("bucket.PutObject failed", zap.Error(err))
		return "", err
	}
	return key, nil
}
