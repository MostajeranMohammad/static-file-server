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
	buckets, err := u.minioClient.ListBuckets(context.Background())
	if err != nil {
		return err
	}

	mb := make(map[string]bool, len(consts.Buckets))
	for x := range consts.Buckets {
		mb[x] = true
	}
	var bucketsToCreate []string
	for _, x := range buckets {
		if !mb[x.Name] {
			bucketsToCreate = append(bucketsToCreate, x.Name)
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

func (u *objectStorageManagerUseCase) GetObject(ctx context.Context, bucketName string, objectName string) (io.Reader, error) {
	object, err := u.minioClient.GetObject(ctx, bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	defer object.Close()

	return object, nil
}

func (u *objectStorageManagerUseCase) SaveObject(ctx context.Context, bucketName string, fileName string, file io.Reader, fileSize int64) (minio.UploadInfo, error) {
	uploadInfo, err := u.minioClient.PutObject(ctx, bucketName, fileName, file, int64(fileSize), minio.PutObjectOptions{})
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
