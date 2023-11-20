package clarity

import (
	fileModel "c-vod/models/fileModel"
	"c-vod/utils/globals"
	"c-vod/utils/helper"
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

/*
	Stigma will store the finalized files to object storage
*/

type Stigma struct{}

func (st *Stigma) storeFileInStorage(file *fileModel.File) error {

	var err error

	// + stigma will store the files
	err = st.StoreFinalDir(file)

	if err != nil {
		//TODO: rollback
		fmt.Println("error on moving files to s3 storage :", err.Error())
		return err
	}

	// + stigma will delete stage3 files
	err = Trauma.deleteObjectDirectory(file)
	if err != nil {
		//TODO: rollback
		fmt.Println("error on removing file from stage2 :", err.Error())
		return err
	}

	// + stigma will call edborn to update state to completed
	err = Edborn.updateFileStatus(file, fileModel.COMPLETED)
	if err != nil {
		//TODO: rollback
		fmt.Println("error on updating file status: ", err.Error())
		return err
	}

	return nil
}

func (st *Stigma) mockStoreFileInStorage(file *fileModel.File) {

	file_id := fmt.Sprintf("%d", file.Id)

	source_path := filepath.Join(helper.GetCurrentDir(), "./storage/stage3/"+file_id)

	destination_path := filepath.Join(helper.GetCurrentDir(), "./storage/stage4/"+file_id)

	err := helper.MoveDirectory(source_path, destination_path)

	if err != nil {
		//TODO: rollback
		fmt.Println("mockStoreFileInStorage failed :", err.Error())
		return
	}
}

func (st *Stigma) uploadFileToStorage(file *fileModel.File) {

	file_id := fmt.Sprintf("%d", file.Id)
	ctx := context.Background()
	endpoint := "play.min.io"
	accessKeyID := "Q3AM3UQ867SPQQA43P2F"
	secretAccessKey := "zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG"
	useSSL := true

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	// Make a new bucket called testbucket.
	bucketName := "vod-storage-test"
	//location := "us-east-1"

	err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Successfully created %s\n", bucketName)
	}

	// Upload the test file
	// Change the value of filePath if the file is in another location
	objectName := file_id + "." + file.Ext
	filePath := filepath.Join(helper.GetCurrentDir(), "packager")
	contentType := "application/octet-stream"

	// Upload the test file with FPutObject
	info, err := minioClient.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Successfully uploaded %s of size %d\n", objectName, info.Size)
}

func (st *Stigma) Test_upload_file_to_s3() {

	file_id := "1"

	ctx := context.Background()

	endpoint := "s3.ir-thr-at1.arvanstorage.ir"
	accessKeyID := "dd9f0541-134a-472d-ac68-648805841048"
	secretAccessKey := "f59ca61b2a584b6d5738ecc3d120d3e7da35a510"
	useSSL := true

	// Initialize minio client object
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})

	if err != nil {
		log.Fatalln(err)
	}

	// Make a new bucket called testbucket.
	bucketName := "vod-storage-test"
	//location := "us-east-1"

	/*
		err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
		if err != nil {
			// Check to see if we already own this bucket (which happens if you run this twice)
			exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
			if errBucketExists == nil && exists {
				log.Printf("We already own %s\n", bucketName)
			} else {
				log.Fatalln(err)
			}
		} else {
			log.Printf("Successfully created %s\n", bucketName)
		}
	*/

	// Upload the test file
	// Change the value of filePath if the file is in another location
	objectName := "edu-arch/1/2.mp4"

	filePath := filepath.Join(helper.GetCurrentDir(), "storage/stage3/1", file_id+".mp4")

	contentType := "application/octet-stream"

	// Upload the test file with FPutObject
	info, err := minioClient.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Successfully uploaded %s of size %d\n", objectName, info.Size)
}

func (st *Stigma) StoreFinalDir(file *fileModel.File) error {

	file_id := fmt.Sprintf("%d", file.Id)

	source_path := filepath.Join(helper.GetCurrentDir(), "./storage/stage3/"+file_id)

	files, err := os.ReadDir(source_path)

	if err != nil {
		return err
	}

	for _, f := range files {

		objectName := "edu-arch/" + file_id + "/" + f.Name()
		filePath := source_path + "/" + f.Name()
		info, err := globals.App.Storage.StoreObject(objectName, filePath)

		if err != nil {
			//TODO: return and rollback
			return err
		}

		if globals.App.Config.Log_enabled == "true" {
			fmt.Println(info)
		}
	}

	return nil
}
