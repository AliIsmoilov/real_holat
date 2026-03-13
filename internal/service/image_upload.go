package service

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type R2ServiceI interface {
	UploadImage(ctx context.Context, bucket, key string, body io.Reader, contentType string) (string, error)
}

type R2Service struct {
	client *s3.Client
}

func NewR2Service(client *s3.Client) *R2Service {
	return &R2Service{client: client}
}

// UploadImage uploads an object to the provided bucket/key and returns the public URL
func (r *R2Service) UploadImage(ctx context.Context, bucket, key string, body io.Reader, contentType string) (string, error) {
	_, err := r.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(key),
		Body:        body,
		ContentType: aws.String(contentType),
		ACL:         types.ObjectCannedACLPublicRead,
	})
	if err != nil {
		return "", err
	}

	publicURL := fmt.Sprintf("https://cdn.safar-uz.com/%s", key)
	_ = time.Now()

	return publicURL, nil
}
