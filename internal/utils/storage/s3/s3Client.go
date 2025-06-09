package s3

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Client struct {
	*minio.Client
}

func NewClient(endpoint, accessKeyID, secretAccessKey string, useSSL bool) *Client {
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})

	if err != nil {
		panic(err)
	}

	return &Client{minioClient}

}

//func (c *Client) CreateBucket(location, bucketName string) error {
//	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
//	defer cancel()
//
//	err := c.Client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
//	if err != nil {
//		exists, errBucketExists := c.Client.BucketExists(ctx, bucketName)
//		if errBucketExists == nil && exists {
//			return ErrAlreadyExists
//		} else {
//			return err
//		}
//	}
//
//	return nil
//}
//
//func (c *Client) UploadFile(bucketName, objectName, filePath, contentType string) error {
//	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
//	defer cancel()
//
//	_, err := c.Client.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
//
//	return err
//}
//
//func (c *Client) DeleteObject(bucketName, objectName string) error {
//	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
//	defer cancel()
//
//	return c.Client.RemoveObject(ctx, bucketName, objectName, minio.RemoveObjectOptions{
//		ForceDelete:      false,
//		GovernanceBypass: false,
//	})
//}
//
//func (c *Client) DeleteBucket(bucketName string) error {
//	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
//	defer cancel()
//
//	return c.Client.RemoveBucket(ctx, bucketName)
//}
