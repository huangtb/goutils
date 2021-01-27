package aws

import (
	"fmt"
	"testing"
)

func TestGetS3Object(t *testing.T) {
	aws := Aws{
		Region:    "",
		AccessKey: "",
		SecretKey: "",
	}
	err := aws.InitS3Client()
	if err != nil {
		panic(err)
	}
	s,err := GetS3Object("","")
	fmt.Println(s)

}

