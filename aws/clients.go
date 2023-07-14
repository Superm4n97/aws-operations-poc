package aws

import (
	"github.com/aws/aws-sdk-go/service/ec2"
	_ "github.com/aws/aws-sdk-go/service/ec2"
)

func EC2Client() (*ec2.EC2, error) {
	sess, err := newSession()
	if err != nil {
		return nil, err
	}
	return ec2.New(sess), nil
}
