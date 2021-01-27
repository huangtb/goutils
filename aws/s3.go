package aws

import (
	"bufio"
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

	sess, err := session.NewSession(a.GetConfig())
	if err != nil {
		return errors.Errorf("New S3 Session creation error: %v", err.Error())
	}
	S3Cli = s3.New(sess)
	return nil
}

func PutObjectToS3(bucket, s3Path string, b []byte) (*s3.PutObjectOutput, error) {
	out, err := S3Cli.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(s3Path),
		ACL:    aws.String(glacier.CannedACLPublicRead),
		Body:   bytes.NewReader(b),
	})
	return out, err
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
		Bucket:        aws.String(bucket),
		Key:           aws.String(s3key),
		ACL:           aws.String(glacier.CannedACLPublicRead),
		Body:          bytes.NewReader(buffer),
		ContentLength: aws.Int64(size),
		//ContentType:          aws.String(http.DetectContentType(buffer)),
		//ContentDisposition:   aws.String("attachment"),	//压缩
		//ServerSideEncryption: aws.String("AES256"),
	})

	if err != nil {
		return err
	}

	return nil
}


func GetS3Object(bucket, key string) (string, error) {
	out, err := S3Cli.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return "" ,err
	}
	defer out.Body.Close()
	br := bufio.NewReader(out.Body)
	return readLine(br)
}

func readLine(r *bufio.Reader) (string, error) {
	line, isprefix, err := r.ReadLine()
	for isprefix && err == nil {
		var bs []byte
		bs, isprefix, err = r.ReadLine()
		line = append(line, bs...)
	}
	return string(line), err
}
