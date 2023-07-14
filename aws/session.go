package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

var (
	region = "us-east-2"
)

func newSession() (*session.Session, error) {
	awsCfg := &aws.Config{
		Region:      &region,
		Credentials: credentials.NewStaticCredentials(AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY, ""),
	}

	sess, err := session.NewSession(awsCfg)
	if err != nil {
		return nil, err
	}
	return sess, nil
}
