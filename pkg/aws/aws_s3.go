package aws

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	awsSession "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/globalsign/mgo/bson"
)

type AWSS3 struct {
	AWSConfig aws.Config
	Region    string
	Bucket    string
	Directory string
	AccessKey string
	SecretKey string
}

func (awss3 AWSS3) UploadManager(fileToBeUploaded *multipart.FileHeader) (s3path string, fileName string, err error) {
	session, err := awsSession.NewSession(&awss3.AWSConfig)
	if err != nil {
		return s3path, fileName, err
	}

	file, err := fileToBeUploaded.Open()
	if err != nil {
		return s3path, fileName, err
	}
	size := fileToBeUploaded.Size
	buffer := make([]byte, size)
	file.Read(buffer)
	fileName = bson.NewObjectId().Hex() + "_" + fileToBeUploaded.Filename
	//fileName = fileToBeUploaded.Filename + bson.NewObjectId().Hex() + filepath.Ext(fileToBeUploaded.Filename)

	tempFile := awss3.Directory + "/" + fileName
	_, err = s3.New(session).PutObject(&s3.PutObjectInput{
		ACL:                aws.String("public-read"),
		Body:               bytes.NewReader(buffer),
		Bucket:             aws.String(awss3.Bucket),
		ContentDisposition: aws.String("attachment"),
		ContentLength:      aws.Int64(int64(size)),
		ContentType:        aws.String(http.DetectContentType(buffer)),
		Key:                aws.String(tempFile),
	})
	if err != nil {
		return s3path, fileName, err
	}

	s3path = awss3.Directory + "/" + fileName

	return s3path, fileName, err
}

func (cred *AWSS3) GetURL(key string) (res string, err error) {
	sess, err := awsSession.NewSession(&aws.Config{
		Endpoint:    cred.AWSConfig.Endpoint,
		Region:      cred.AWSConfig.Region,
		Credentials: credentials.NewStaticCredentials(cred.AccessKey, cred.SecretKey, ""),
	})
	if err != nil {
		return res, err
	}

	svc := s3.New(sess)
	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(cred.Bucket),
		Key:    aws.String(key),
	})

	res, err = req.Presign(15 * time.Minute)

	return res, err
}

func (awss3 AWSS3) DeleteManager(fileToBeRemoved string) (s3path string, fileName string, err error) {
	session, err := awsSession.NewSession(&awss3.AWSConfig)
	if err != nil {
		return s3path, fileName, err
	}

	fmt.Println("Data gambar = "+fileToBeRemoved, awss3)
	fmt.Println(awss3.Bucket)

	_, err = s3.New(session).DeleteObject(&s3.DeleteObjectInput{
		Key:    aws.String(fileToBeRemoved),
		Bucket: &awss3.Bucket,
	})
	if err != nil {
		return s3path, fileName, err
	}

	s3path = awss3.Directory + "/" + fileName

	return s3path, fileName, err
}
