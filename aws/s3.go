package aws

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/glacier"
	"github.com/aws/aws-sdk-go/service/s3"
	"net/http"
	"os"
	"path"
)

var s3Cli *s3Client

type s3Client struct {
	client *s3.S3
}

func GetS3Client() *s3Client {
	return s3Cli
}

func (a *Aws) NewS3Client() (sc s3Client, err error) {
	creds := credentials.NewStaticCredentials(a.AccessKey, a.SecretKey, "")
	_, err = creds.Get()
	if err != nil {
		return sc, errors.New("New Static Credentials  error:" + err.Error())
	}
	cfg := aws.NewConfig().WithRegion(a.Region).WithCredentials(creds)
	sc.client = s3.New(session.New(), cfg)
	s3Cli = &sc
	return sc,nil
}

func (s *s3Client) Upload(bucket, prefix, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer func() {
		file.Close()
		//删除压缩包
		err := os.Remove(filePath)
		if err != nil {
			return
		}
	}()

	fileInfo, _ := file.Stat()
	size := fileInfo.Size()
	buffer := make([]byte, size)
	file.Read(buffer)
	s3key := fmt.Sprintf("%s/%s", prefix, path.Base(filePath))
	_, err = s.client.PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(bucket),
		Key:                  aws.String(s3key),
		ACL:                  aws.String(glacier.CannedACLPublicRead),
		Body:                 bytes.NewReader(buffer),
		ContentLength:        aws.Int64(size),
		ContentType:          aws.String(http.DetectContentType(buffer)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
	})

	if err != nil {
		return err
	}

	return nil
}
