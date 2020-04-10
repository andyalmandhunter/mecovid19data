package config

import (
	"log"
	"os"
	"strconv"
	"time"
)

var AwsRegion = os.Getenv("AWS_REGION")
var S3Bucket = os.Getenv("AWS_S3_BUCKET")
var S3DataKey = os.Getenv("AWS_S3_DATA_KEY")
var PollingPeriod time.Duration

func init() {
	if AwsRegion == "" {
		log.Fatal("AWS_REGION undefined")
	}

	if S3Bucket == "" {
		log.Fatal("AWS_S3_BUCKET undefined")
	}

	if S3DataKey == "" {
		log.Fatal("AWS_S3_DATA_KEY undefined")
	}

	p, err := strconv.ParseUint(os.Getenv("POLLING_PERIOD_SECONDS"), 10, 64)
	if err != nil {
		log.Fatalf("Error parsing POLLING_PERIOD_SECONDS: %v\n", err)
	}
	PollingPeriod = time.Duration(p) * time.Second
}
