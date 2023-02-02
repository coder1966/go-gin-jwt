package service

import (
	"context"
	"github.com/qiniu/api.v7/v7/storage"
	"go-gin-jwt/util"
	"go-gin-jwt/util/errmsg"
	"mime/multipart"

	"github.com/qiniu/api.v7/v7/auth/qbox"
)

var ImgUrl = util.ImgUrl
var AccessKey = util.AccessKey
var SecretKey = util.SecretKey
var Bucket = util.Bucket

func UploadFile(file multipart.File, fileSize int64) (string, int) {
	putPolicy := storage.PutPolicy{
		Scope: Bucket,
	}
	mac := qbox.NewMac(AccessKey, SecretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{
		Zone:          &storage.ZoneHuadong,
		UseCdnDomains: false,
		UseHTTPS:      false,
	}

	putExtra := storage.PutExtra{}

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	err := formUploader.PutWithoutKey(context.Background(), &ret, upToken, file, fileSize, &putExtra)
	if err != nil {
		return "", errmsg.ERROR
	}
	url := ImgUrl + ret.Key
	return url, errmsg.SUCCESS
}
