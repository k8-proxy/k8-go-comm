<h1 align="center">k8-go-comm</h1>
<p align="center">
    <em> k8-go-comm library.</em>
</p>

<p align="center">
    <a href="https://goreportcard.com/report/github.com/k8-proxy/k8-go-comm">
        <img src="https://goreportcard.com/badge/k8-proxy/k8-go-comm" alt="Go Report Card">
    </a>
	<a href="https://github.com/k8-proxy/k8-go-comm/pulls">
        <img src="https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat" alt="Contributions welcome">
    </a>
    <a href="https://opensource.org/licenses/Apache-2.0">
        <img src="https://img.shields.io/badge/License-Apache%202.0-blue.svg" alt="Apache License, Version 2.0">
    </a>
    <a href="https://pkg.go.dev/github.com/k8-proxy/k8-go-comm">
        <img src="https://godoc.org/github.com/k8-proxy/k8-go-comm?status.svg" alt="k8-go-comm">
    </a>
</p>

# k8-go-comm

Library to communicate with RabbitMQ and MinIO

## Goals

Provide helper modules which will be used by all parties willing to implement communications with RabbitMQ or Minio. All common types and functions needs to be placed here.

## Usage

See the 'examples' subdirectory with example for using these modules. Each usecase is in a specific subdirectory inside examples.<br>
For quick start using docker to run containers for RabbitMQ and MinIO<br>
Run Standalone MinIO on Docker.

```
docker run -p 9000:9000 \
  -e "MINIO_ROOT_USER=<your_minio_root_user>" \
  -e "MINIO_ROOT_PASSWORD=<your_minio_root_password>" \
  minio/minio server /data

```

Run RabbitMQ on Docker.

```
docker run -d --hostname <your_host_name> --name <your_container_name> -p 15672:15672 -p 5672:5672 rabbitmq:3-management
```

## RabbitMQ

### Producer

To start a connection with RabbitMQ and publish a message, please check [producer folder](examples/producer)

### Consumer

To start a connection and consume the message that was published, please check [consumer directory](examples/consumer)

## MinIO

### Create new bucket

to create a new minio bucket , please check [createBucket directory](examples/createBucket)

### Make bucket and upload file

To start a connection and upload a file, please check [minioUpload directory](examples/minioUpload)

### Get the object from the bucket

To start a connection and get a file from specific bucket, please check [minioGetObject directory](examples/minioGetObject)

### Get the presigned URL

To start a connection and get the presigned URL of the file you uploaded, please check [minioGetPresignedURL directory](examples/minioGetPresignedURL)
