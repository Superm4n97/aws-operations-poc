package aws

import (
	_credentials "github.com/Superm4n97/aws-operations-poc/utils/credentials"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func newSession() (*session.Session, error) {
	region := _credentials.Region
	awsCfg := &aws.Config{
		Region:      &region,
		Credentials: credentials.NewStaticCredentials(_credentials.AWS_ACCESS_KEY_ID, _credentials.AWS_SECRET_ACCESS_KEY, ""),
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
