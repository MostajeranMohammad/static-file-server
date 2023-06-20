package usecase

import (
	"context"
	"io"

	"github.com/MostajeranMohammad/static-file-server/internal/contracts/consts"
	"github.com/minio/minio-go/v7"
)

type objectStorageManagerUseCase struct {
	minioClient *minio.Client
}

func NewObjectStorageManagerUseCase(minioClient *minio.Client) ObjectStorageManager {
	return &objectStorageManagerUseCase{
		minioClient: minioClient,
	}
}

func (u *objectStorageManagerUseCase) InitBuckets() error {
	existingBuckets, err := u.minioClient.ListBuckets(context.Background())
	if err != nil {
		return err
	}

	mb := make(map[string]bool)
	for _, x := range existingBuckets {
		mb[x.Name] = true
	}
	var bucketsToCreate []string
	for x := range consts.Buckets {
		if !mb[x] {
			bucketsToCreate = append(bucketsToCreate, x)
		}
	}
	for _, bucketName := range bucketsToCreate {
		err = u.minioClient.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{})
		if err != nil {
			return err
		}
	}
	return nil
}

func (u *objectStorageManagerUseCase) GetObject(ctx context.Context, bucketName string, objectName string) (io.Reader, func() error, error) {
	object, err := u.minioClient.GetObject(ctx, bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		return nil, object.Close, err
	}
	return object, object.Close, nil
}

func (u *objectStorageManagerUseCase) SaveObject(ctx context.Context, bucketName string, fileName string, file io.Reader, fileSize int64) (minio.UploadInfo, error) {
	uploadInfo, err := u.minioClient.PutObject(ctx, bucketName, fileName, file, fileSize, minio.PutObjectOptions{})
	if err != nil {
		return minio.UploadInfo{}, err
	}
	return uploadInfo, nil
}

func (u *objectStorageManagerUseCase) DeleteObject(ctx context.Context, bucketName string, fileName string) error {
	err := u.minioClient.RemoveObject(ctx, bucketName, fileName, minio.RemoveObjectOptions{})
	if err != nil {
		return err
	}
	return nil
}
