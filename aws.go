package main

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func awsAccountCheck() bool {
	svc := s3.New(session.New())
	input := &s3.ListBucketsInput{}

	_, err := svc.ListBuckets(input)

	if err != nil {
		return false
	} else {
		return true
	}
}
