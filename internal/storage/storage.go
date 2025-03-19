package storage

import (
    "context"
  

    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/service/s3"
)

// S3Client wraps an S3 client with helpful methods
type S3Client struct {
    client *s3.Client
    bucket string
}

// NewS3Client creates a new S3 client
func NewS3Client(bucket string) (*S3Client, error) {
    // Load the AWS SDK configuration
    cfg, err := config.LoadDefaultConfig(context.TODO())
    if err != nil {
        return nil, err
    }

    // Create the S3 client
    client := s3.NewFromConfig(cfg)

    return &S3Client{
        client: client,
        bucket: bucket,
    }, nil
}

// GetClient returns the underlying S3 client
func (s *S3Client) GetClient() *s3.Client {
    return s.client
}

// GetBucket returns the configured bucket name
func (s *S3Client) GetBucket() string {
    return s.bucket
}