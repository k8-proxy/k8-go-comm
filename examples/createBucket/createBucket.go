package main

import (
	"log"
	"os"

	minio "github.com/k8-proxy/k8-go-comm/pkg/minio"
	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

func main() {
	endpoint := os.Getenv("MINIO_ENDPOINT")
	accessKey := os.Getenv("ACCESS_KEY")
	secretKey := os.Getenv("SECRET_KEY")
	bucketName := os.Getenv("BUCKET_NAME")

	endpoint = "localhost:9000"
	accessKey = "minioadmin"
	secretKey = "minioadmin"
	bucketName = "mytest3t"

	client, err := minio.NewMinioClient(endpoint, accessKey, secretKey, false)
	if err != nil {
		log.Println(err)
		return
	}

	//Check if bucket already exists
	exists, errBucketExists := minio.CheckIfBucketExists(client, bucketName)
	if errBucketExists != nil {
		log.Println("error checkbucket ", err)
		return
	}
	if exists {
		log.Println(" bucket exists ")
		return
	}

	//If bucket doesn't exists create a bucket
	err = minio.CreateNewBucket(client, bucketName)
	if err != nil {
		log.Println("bucket not created ", err)
	} else {
		log.Println("bucket created succefully ")

	}
}
