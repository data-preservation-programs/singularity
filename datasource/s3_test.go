package datasource

import (
	"bytes"
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"testing"
	"time"
)

type mockS3API struct {
	mock.Mock
}

func (m *mockS3API) GetObject(ctx context.Context, params *s3.GetObjectInput, optFns ...func(*s3.Options)) (*s3.GetObjectOutput, error) {
	args := m.Called(ctx, params, optFns)
	return args.Get(0).(*s3.GetObjectOutput), args.Error(1)
}

func (m *mockS3API) ListObjectsV2(ctx context.Context, params *s3.ListObjectsV2Input, optFns ...func(*s3.Options)) (*s3.ListObjectsV2Output, error) {
	args := m.Called(ctx, params, optFns)
	return args.Get(0).(*s3.ListObjectsV2Output), args.Error(1)
}

func (m *mockS3API) HeadObject(ctx context.Context, params *s3.HeadObjectInput, optFns ...func(*s3.Options)) (*s3.HeadObjectOutput, error) {
	args := m.Called(ctx, params, optFns)
	return args.Get(0).(*s3.HeadObjectOutput), args.Error(1)
}

func TestS3_Open(t *testing.T) {
	mockData := []byte("This is a test object.")
	mockClient := new(mockS3API)
	mockClient.On("GetObject", mock.Anything, mock.Anything, mock.Anything).Return(&s3.GetObjectOutput{
		Body: ioutil.NopCloser(bytes.NewReader([]byte("test "))),
	}, nil)

	s3 := S3{
		s3Client: mockClient,
	}

	ctx := context.Background()
	path := "s3://test-bucket/test-object"
	offset := uint64(10)
	length := uint64(5)

	reader, err := s3.Read(ctx, path, offset, length)
	assert.NoError(t, err)

	data, err := ioutil.ReadAll(reader)
	assert.NoError(t, err)
	assert.Equal(t, mockData[offset:offset+length], data)

	// Test invalid path
	_, err = s3.Read(ctx, "invalid-path", 0, 10)
	assert.Error(t, err)
}

func TestS3_Scan(t *testing.T) {
	mockS3 := &mockS3API{}
	s3Client := S3{s3Client: mockS3}

	// Set up the mock response
	mockS3.On("ListObjectsV2", mock.Anything, mock.Anything, mock.Anything).Return(
		&s3.ListObjectsV2Output{
			Contents: []types.Object{
				{
					Key:          aws.String("file1.txt"),
					LastModified: aws.Time(time.Now()),
					Size:         100,
				},
				{
					Key:          aws.String("file2.txt"),
					LastModified: aws.Time(time.Now()),
					Size:         200,
				},
			},
		}, nil,
	)

	ctx := context.Background()
	path := "s3://my-bucket/"
	last := "s3://my-bucket/file0.txt"

	entriesCh := s3Client.Scan(ctx, path, last)

	entry1 := <-entriesCh
	assert.Equal(t, "s3://my-bucket/file1.txt", entry1.Path)
	assert.EqualValues(t, int64(100), entry1.Size)

	entry2 := <-entriesCh
	assert.Equal(t, "s3://my-bucket/file2.txt", entry2.Path)
	assert.EqualValues(t, int64(200), entry2.Size)

	_, ok := <-entriesCh
	assert.False(t, ok, "The channel should be closed")

	mockS3.AssertExpectations(t)
}
