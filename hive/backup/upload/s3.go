package upload

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"os"
)

var (
	bucket string
	key    string
)

// UploadS3
func UploadS3(database string, file string) {

	fileReader, err := os.Open(file)
	if err != nil {
		fmt.Printf("Unable to open file %q, %v", err)
	}
	defer fileReader.Close()

	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("eu-central-1"),
	}))

	uploader := s3manager.NewUploader(sess)

	bucket = "rds-pgdump-backups"
	key = file

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(database + "/" + file),

		Body: fileReader,
	})
	if err != nil {
		fmt.Printf("Unable to upload %q to %q, %v", file, bucket, err)
	}

	fmt.Printf("Successfully uploaded %q to %q\n", file, bucket)

}

