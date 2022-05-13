package main

import (
	"context"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func main() {
	endpoint := "play.min.io"
	accessKeyID := "Q3AM3UQ867SPQQA43P2F"
	secretAccessKey := "zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG"

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: true,
	})
	if err != nil {
		log.Panicln(err)
	}
	log.Printf("%#v\n", minioClient)

	ctx := context.Background()
	bucketName := "00test"
	objectName := "Golock Holmes.png"
	filePath := "/Users/kimdh/Dropbox/Images/Golock Holmes.png"
	contentType := "image/png"
	info, err := minioClient.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		log.Panicln(err)
	}
	log.Printf("Successfully uploaded %s of size %d\n", objectName, info.Size)
}
