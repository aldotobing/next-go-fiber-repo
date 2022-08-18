package minio

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// Connection ..
type Connection struct {
	AccessKey   string
	SecretKey   string
	UseSSL      bool
	EndSuperhub string
	Duration    int
}

// InitClient ...
func (conn Connection) InitClient() (client *minio.Client, err error) {
	client, err = minio.New(conn.EndSuperhub, &minio.Options{
		Creds:  credentials.NewStaticV4(conn.AccessKey, conn.SecretKey, ""),
		Secure: conn.UseSSL,
	})
	if err != nil {
		return client, err
	}

	return client, nil
}
