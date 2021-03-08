package minio

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"time"

	minio "github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// NewMinioClient - returns new minio client
func NewMinioClient(endpoint string, accessKeyID string, secretAccessKey string, useSSL bool) (*minio.Client, error) {

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})

	return minioClient, err
}

// UploadFileToMinio - uploads file to minio
func UploadFileToMinio(client *minio.Client, bucketName string, objectName string, reader io.Reader) (minio.UploadInfo, error) {
	conttype := "application/octet-stream"

	f, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Println(err)
	}

	conttype = http.DetectContentType(f[:511])
	log.Println(conttype)
	reader = bytes.NewReader(f)

	uploadInfo, err := client.PutObject(context.Background(), bucketName, objectName, reader, -1, minio.PutObjectOptions{ContentType: conttype})
	if err != nil {
		log.Fatalln(err)
		return uploadInfo, err
	}
	return uploadInfo, nil
}

// GetObjectFromMinio - get file from minio
func GetObjectFromMinio(client *minio.Client, bucketName string, objectName string) (*minio.Object, error) {
	object, err := client.GetObject(context.Background(), bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		fmt.Println(err)
		return object, nil
	}
	return object, nil
}

// DeleteObjectInMinio - delete object in minio
func DeleteObjectInMinio(client *minio.Client, bucketName string, objectName string) error {
	err := client.RemoveObject(context.Background(), bucketName, objectName, minio.RemoveObjectOptions{})
	return err
}

// GetPresignedURLForObject - Get signed url for object
func GetPresignedURLForObject(client *minio.Client, bucketName string, objectName string, expiresIn time.Duration) (*url.URL, error) {
	// Set request parameters for content-disposition.
	reqParams := make(url.Values)
	// reqParams.Set("response-content-disposition", "attachment; filename=\"your-filename.txt\"")
	// Generates a presigned url which expires in a day.
	presignedURL, err := client.PresignedGetObject(context.Background(), bucketName, objectName, expiresIn, reqParams)
	if err != nil {
		fmt.Println(err)
		return presignedURL, err
	}
	return presignedURL, err
}

// UploadAndReturnURL to upload a file with exp date and return its URL
func UploadAndReturnURL(client *minio.Client, bucketName string, fileFullPath string, expiresIn time.Duration) (*url.URL, error) {

	contentType, err := getContentType(fileFullPath)
	if err != nil {
		log.Fatalln(err)
		return &url.URL{}, err
	}
	objectName := filepath.Base(fileFullPath)
	_, err = client.FPutObject(context.Background(), bucketName, objectName, fileFullPath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return &url.URL{}, err
	}

	return GetPresignedURLForObject(client, bucketName, objectName, expiresIn)
}

func getContentType(fileFullPath string) (string, error) {

	out, err := os.Open(fileFullPath)
	if err != nil {
		return "", err
	}

	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)

	_, err = out.Read(buffer)
	if err != nil {
		return "", err
	}

	// Use the net/http package's handy DectectContentType function. Always returns a valid
	// content-type by returning "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(buffer)

	return contentType, nil
}

// DownloadObject to download theoject
func DownloadObject(cleanPresignedURL string, outputFileLocation string) error {

	// Get the data
	resp, err := http.Get(cleanPresignedURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(outputFileLocation)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err

}

func CheckIfBucketExists(client *minio.Client, bucketName string) (bool, error) {

	// Check to see if we already own this bucket
	exists, errBucketExists := client.BucketExists(context.Background(), bucketName)

	if errBucketExists != nil {
		log.Fatalln(errBucketExists)
	}

	return exists, errBucketExists
}

//Create new bucket
func CreateNewBucket(client *minio.Client, bucketName string) error {
	opt := minio.MakeBucketOptions{Region: "us-east-1"}
	err := client.MakeBucket(context.Background(), bucketName, opt)

	if err != nil {
		log.Fatalln(err)
	}

	return err
}
