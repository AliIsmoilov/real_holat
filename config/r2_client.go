package config

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func CreateR2Client(cfg Config) *s3.Client {
	endpoint := fmt.Sprintf("https://%s.r2.cloudflarestorage.com", cfg.CloudflareR2.R2_ACCOUNT_ID)

	awsCfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("auto"),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(
				cfg.CloudflareR2.R2_ACCESS_KEY_ID,
				cfg.CloudflareR2.R2_SECRET_ACCESS_KEY,
				"",
			),
		),
	)
	if err != nil {
		panic(fmt.Sprintf("failed to init Cloudflare R2: %v", err))
	}

	return s3.NewFromConfig(awsCfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(endpoint)
		o.UsePathStyle = true
	})
}
