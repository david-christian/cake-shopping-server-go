package services

import (
	"context"
	"io"
	"log"
	"mime/multipart"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

func (s *Service) UploadImage(files []*multipart.FileHeader) (bool, error) {
	credentialFilePath := "../keyfile/shopping-go-397813-3dd2185818ff.json"
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(credentialFilePath))
	if err != nil {
		log.Fatal(err)
		return false, err
	}

	openFile, err := files[0].Open()
	if err != nil {
		return false, err
	}
	defer openFile.Close()

	bucketName := "shopping-go"

	// 创建存储桶对象
	bucket := client.Bucket(bucketName)
	obj := bucket.Object(files[0].Filename)

	// 创建对象写入器
	writer := obj.NewWriter(ctx)
	defer writer.Close()

	if _, err := io.Copy(writer, openFile); err != nil {
		log.Fatalf("Failed to copy file to object: %v", err)
		return false, err
	}
	return true, nil
}
