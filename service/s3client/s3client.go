package s3client

import (
	"bytes"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"io"
	"log"

	"mecovid19data/service/config"
)

var sess *session.Session
var downloader *s3manager.Downloader

func init() {
	sess = session.Must(session.NewSession(
		&aws.Config{Region: aws.String(config.AwsRegion)},
	))

	downloader = s3manager.NewDownloader(sess)
}

func Get() (io.Reader, error) {
	buf := aws.NewWriteAtBuffer([]byte{})

	n, err := downloader.Download(buf, &s3.GetObjectInput{
		Bucket: aws.String(config.S3Bucket),
		Key:    aws.String(config.S3DataKey),
	})
	if err != nil {
		return nil, err
	}

	log.Printf("Downloaded %d bytes\n", n)
	return bytes.NewBuffer(buf.Bytes()), nil
}
