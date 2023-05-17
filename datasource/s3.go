package datasource

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/data-preservation-programs/go-singularity/model"
	"github.com/pkg/errors"
	"io"
	"net/url"
	"strings"
	"time"
)

type S3API interface {
	GetObject(ctx context.Context, params *s3.GetObjectInput, optFns ...func(*s3.Options)) (*s3.GetObjectOutput, error)
	HeadObject(ctx context.Context, params *s3.HeadObjectInput, optFns ...func(*s3.Options)) (*s3.HeadObjectOutput, error)
	ListObjectsV2(context.Context, *s3.ListObjectsV2Input, ...func(*s3.Options)) (*s3.ListObjectsV2Output, error)
}

type S3 struct {
	s3Client S3API
}

func NewS3(ctx context.Context, region string, endpoint string, accessKeyID string, secretAccessKey string) (S3, error) {
	var configs []func(*config.LoadOptions) error
	if region != "" {
		configs = append(configs, config.WithRegion(region))
	}
	if endpoint != "" {
		configs = append(configs, config.WithEndpointResolverWithOptions(
			aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
				return aws.Endpoint{
					URL: endpoint,
				}, nil
			})))
	}
	if secretAccessKey != "" {
		configs = append(configs, config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyID, secretAccessKey, "")))
	} else {
		configs = append(configs, config.WithCredentialsProvider(aws.AnonymousCredentials{}))
	}

	awsConfig, err := config.LoadDefaultConfig(ctx, configs...)
	if err != nil {
		return S3{}, fmt.Errorf("failed to load AWS config: %w", err)
	}

	s3Client := s3.NewFromConfig(awsConfig)
	return S3{s3Client: s3Client}, nil
}

func (s S3) CheckItem(ctx context.Context, path string) (uint64, *time.Time, error) {
	parsedPath, err := url.Parse(path)
	if err != nil {
		return 0, nil, errors.Wrap(err, "failed to parse S3 path")
	}

	if parsedPath.Scheme != "s3" {
		return 0, nil, errors.New("invalid S3 path, missing s3:// scheme")
	}

	bucket := parsedPath.Host
	key := strings.TrimLeft(parsedPath.Path, "/")

	params := &s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}

	resp, err := s.s3Client.HeadObject(ctx, params)
	if err != nil {
		return 0, nil, errors.Wrap(err, "failed to get S3 object")
	}

	return uint64(resp.ContentLength), resp.LastModified, nil
}
func (s S3) Scan(ctx context.Context, path string, last string) <-chan Entry {
	entries := make(chan Entry)

	go func() {
		defer close(entries)

		parsedPath, err := url.Parse(path)
		if err != nil {
			select {
			case <-ctx.Done():
			case entries <- Entry{
				Error: errors.Wrap(err, "failed to parse S3 path"),
			}:
			}
			return
		}

		if parsedPath.Scheme != "s3" {
			select {
			case <-ctx.Done():
			case entries <- Entry{
				Error: errors.New("invalid S3 path, missing s3:// scheme"),
			}:
			}
			return
		}

		bucket := parsedPath.Host
		prefix := strings.TrimLeft(parsedPath.Path, "/")

		params := &s3.ListObjectsV2Input{
			Bucket: aws.String(bucket),
			Prefix: aws.String(prefix),
		}

		if last != "" {
			lastPath, err := url.Parse(last)
			if err != nil {
				select {
				case <-ctx.Done():
				case entries <- Entry{
					Error: errors.Wrap(err, "failed to parse S3 path"),
				}:
				}
				return
			}
			if lastPath.Scheme != "s3" {
				select {
				case <-ctx.Done():
				case entries <- Entry{
					Error: errors.New("invalid S3 path, missing s3:// scheme"),
				}:
				}
				return
			}
			lastBucket := lastPath.Host
			if lastBucket != bucket {
				select {
				case <-ctx.Done():
				case entries <- Entry{
					Error: errors.New("invalid S3 path, bucket does not match"),
				}:
				}

				return
			}

			lastKey := strings.TrimLeft(lastPath.Path, "/")
			params.StartAfter = aws.String(lastKey)
		}

		paginator := s3.NewListObjectsV2Paginator(s.s3Client, params)

		for paginator.HasMorePages() {
			output, err := paginator.NextPage(ctx)
			if err != nil {
				select {
				case <-ctx.Done():
				case entries <- Entry{
					Error: errors.Wrap(err, "failed to list S3 objects"),
				}:
				}

				return
			}

			for _, object := range output.Contents {
				select {
				case <-ctx.Done():
					return
				case entries <- Entry{
					Path:         fmt.Sprintf("s3://%s/%s", bucket, *object.Key),
					LastModified: object.LastModified,
					Size:         uint64(object.Size),
					Type:         model.S3Object,
					ScannedAt:    time.Now(),
				}:
				}
			}
		}
	}()

	return entries
}

func (s S3) Read(ctx context.Context, path string, offset uint64, length uint64) (io.ReadCloser, error) {
	if length == 0 {
		return &emptyReadCloser{}, nil
	}
	parsedPath, err := url.Parse(path)
	if err != nil {
		return nil, fmt.Errorf("failed to parse S3 path: %w", err)
	}

	if parsedPath.Scheme != "s3" {
		return nil, errors.New("invalid S3 path, missing s3:// scheme")
	}

	bucket := parsedPath.Host
	key := strings.TrimLeft(parsedPath.Path, "/")
	rangeHeaderValue := fmt.Sprintf("bytes=%d-%d", offset, offset+length-1)

	input := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Range:  aws.String(rangeHeaderValue),
	}

	result, err := s.s3Client.GetObject(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("failed to get object: %w", err)
	}

	return result.Body, nil
}

func (s S3) Open(ctx context.Context, path string) (ReadAtCloser, error) {
	parsedPath, err := url.Parse(path)
	if err != nil {
		return nil, fmt.Errorf("failed to parse S3 path: %w", err)
	}

	if parsedPath.Scheme != "s3" {
		return nil, errors.New("invalid S3 path, missing s3:// scheme")
	}

	bucket := parsedPath.Host
	key := strings.TrimLeft(parsedPath.Path, "/")

	return S3ReadAtCloser{
		s3Client: s.s3Client,
		bucket:   bucket,
		key:      key,
		ctx:      ctx,
	}, nil
}

type S3ReadAtCloser struct {
	s3Client S3API
	bucket   string
	key      string
	ctx      context.Context
}

func (s S3ReadAtCloser) ReadAt(p []byte, off int64) (n int, err error) {
	rangeHeaderValue := fmt.Sprintf("bytes=%d-%d", off, off+int64(len(p))-1)

	input := &s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(s.key),
		Range:  aws.String(rangeHeaderValue),
	}

	result, err := s.s3Client.GetObject(s.ctx, input)
	if err != nil {
		return 0, fmt.Errorf("failed to get object: %w", err)
	}

	defer result.Body.Close()
	return result.Body.Read(p)
}

func (s S3ReadAtCloser) Close() error {
	return nil
}
