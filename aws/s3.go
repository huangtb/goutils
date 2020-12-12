package aws

import (
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/glacier"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/pkg/errors"
	"os"
	"path"
)

var S3Cli *s3.S3

func (a *Aws) InitS3Client() error {

	cred := getCredentials(a.AccessKey,a.SecretKey)
	_, err := cred.Get()
	if err != nil {
		return errors.Errorf("New Static Credentials  error:" , err.Error())
	}

	cfg := aws.NewConfig().WithRegion(a.Region).WithCredentials(cred)
	s3 := s3.New(session.New(), cfg)
	S3Cli = s3
	return nil
}

func PutObjectToS3(bucket, s3Path string, b []byte) (*s3.PutObjectOutput,error) {
	out, err := S3Cli.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(s3Path),
		ACL:    aws.String(glacier.CannedACLPublicRead),
		Body:   bytes.NewReader(b),
	})
	return out,err
}


func PutFileToS3(bucket, prefix, filePath string) error {
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
	_, err = S3Cli.PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(bucket),
		Key:                  aws.String(s3key),
		ACL:                  aws.String(glacier.CannedACLPublicRead),
		Body:                 bytes.NewReader(buffer),
		ContentLength:        aws.Int64(size),
		//ContentType:          aws.String(http.DetectContentType(buffer)),
		//ContentDisposition:   aws.String("attachment"),	//压缩
		//ServerSideEncryption: aws.String("AES256"),
	})

	if err != nil {
		return err
	}

	return nil
}


