package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	_ "github.com/aws/aws-sdk-go/service/ec2"
)

var (
	Region = "us-east-2"
)

func newSession() (*session.Session, error) {
	awsCfg := &aws.Config{
		Region:      &Region,
		Credentials: credentials.NewStaticCredentials(AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY, ""),
	}

	sess, err := session.NewSession(awsCfg)
	if err != nil {
		return nil, err
	}
	return sess, nil
}

func EC2Client() (*ec2.EC2, error) {
	sess, err := newSession()
	if err != nil {
		return nil, err
	}
	return ec2.New(sess), nil
}
