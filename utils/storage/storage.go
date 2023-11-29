package storage

import (
	"context"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Storage struct {
	BucketName string
	Client     *minio.Client
}

func InitMinio() (*Storage, error) {

	endpoint := os.Getenv("STORAGE_ENDPOINT")
	accessKeyID := os.Getenv("STORAGE_ACCESS_KEY")
	secretAccessKey := os.Getenv("STORAGE_SECRET_KEY")
	bucketName := os.Getenv("STORAGE_BUCKET_NAME")
	useSSL := true

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})

	if err != nil {
		return nil, err
	}

	storage := &Storage{
		BucketName: bucketName,
		Client:     minioClient,
	}

	return storage, nil
}

func (st *Storage) StoreObject(objectName string, filePath string) (minio.UploadInfo, error) {

	bucketName := st.BucketName

	contentType := "application/octet-stream"

	// Upload the test file with FPutObject
	info, err := st.Client.FPutObject(
		context.Background(),
		bucketName,
		objectName,
		filePath,
		minio.PutObjectOptions{
			ContentType: contentType,
			UserMetadata: map[string]string{
				"x-amz-acl": "public-read",
			},
		},
	)

	if err != nil {
		return minio.UploadInfo{}, err
	}

	return info, nil
}

func (st *Storage) DeleteObject(objectName string) error {

	err := st.Client.RemoveObject(
		context.Background(),
		st.BucketName,
		objectName,
		minio.RemoveObjectOptions{},
	)

	if err != nil {
		return err
	}

	return nil
}

func (st *Storage) DeleteDirectory(objectName string) error {

	// List all objects within the directory.
	objectsCh := st.Client.ListObjects(
		context.Background(),
		st.BucketName,
		minio.ListObjectsOptions{
			Prefix:    objectName,
			Recursive: true,
		},
	)

	// Delete each object within the directory.
	for object := range objectsCh {

		if object.Err != nil {
			return object.Err
		}

		err := st.Client.RemoveObject(
			context.Background(),
			st.BucketName,
			object.Key,
			minio.RemoveObjectOptions{},
		)

		if err != nil {
			return err
		}
	}

	return nil
}
