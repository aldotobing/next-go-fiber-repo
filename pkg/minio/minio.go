package minio

import (
	"context"
	"mime/multipart"
	"net/url"
	"path/filepath"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/minio/minio-go/v7"
)

// IMinio ...
type IMinio interface {
	Upload(bucketName, path string, file *multipart.FileHeader) (string, error)
	UploadOs(bucketName, path, fileName, filePath, contentType string) (string, error)
	GetFile(bucketName, objectName string) (string, error)
	Delete(bucketName, objectName string) error
}

type minioModel struct {
	Client *minio.Client
}

const defaultDuration = 15

// NewMinioModel ...
func NewMinioModel(client *minio.Client) IMinio {
	return &minioModel{Client: client}
}

// Upload ...
func (model minioModel) Upload(bucketName, path string, fileHeader *multipart.FileHeader) (res string, err error) {
	src, err := fileHeader.Open()
	if err != nil {
		return res, err
	}
	defer src.Close()

	fileName := bson.NewObjectId().Hex() + filepath.Ext(fileHeader.Filename)
	contentType := fileHeader.Header.Get("Content-Type")
	path += `/` + fileName

	_, err = model.Client.PutObject(context.Background(), bucketName, path, src, fileHeader.Size, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return res, err
	}
	res = path

	return res, nil
}

// UploadOs ...
func (model minioModel) UploadOs(bucketName, path, fileName, filePath, contentType string) (res string, err error) {
	fileName = bson.NewObjectId().Hex() + fileName
	res = path + `/` + fileName

	_, err = model.Client.FPutObject(context.Background(), bucketName, res, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return res, err
	}

	return res, nil
}

// GetFile ...
func (model minioModel) GetFile(bucketName, objectName string) (res string, err error) {
	reqParams := make(url.Values)

	duration := time.Minute * defaultDuration
	uri, err := model.Client.PresignedGetObject(context.Background(), bucketName, objectName, duration, reqParams)
	if err != nil {
		return res, err
	}
	res = uri.String()

	return res, err
}

// Delete ...
func (model minioModel) Delete(bucketName, objectName string) (err error) {
	options := minio.RemoveObjectOptions{}
	err = model.Client.RemoveObject(context.Background(), bucketName, objectName, options)

	return err
}
