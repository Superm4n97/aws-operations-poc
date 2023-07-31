package ec2

import (
	"errors"
	"fmt"
	"github.com/Superm4n97/aws-operations-poc/utils"
	"github.com/aws/aws-sdk-go/service/ec2"
)

const (
	keyType   = "rsa"
	keyFormat = "pem"
)

func GetKeyPair(c *ec2.EC2, name string) (*ec2.KeyPairInfo, error) {
	out, err := c.DescribeKeyPairs(&ec2.DescribeKeyPairsInput{
		KeyNames: utils.StringPSlice([]string{name}),
	})
	if err != nil {
		return nil, err
	}
	for _, kp := range out.KeyPairs {
		return kp, nil
	}
	return nil, errors.New(fmt.Sprintf("no valid keypair found with name: %s", name))
}

func NewKeyPair(c *ec2.EC2, keypairName string) (*ec2.CreateKeyPairOutput, error) {
	keypairInput := &ec2.CreateKeyPairInput{
		KeyFormat:         utils.StringP(keyFormat),
		KeyName:           &keypairName,
		KeyType:           utils.StringP(keyType),
		TagSpecifications: nil,
	}
	return c.CreateKeyPair(keypairInput)
}

func RemoveKeyPair(c *ec2.EC2, name string) error {
	kp, err := GetKeyPair(c, name)
	if err != nil {
		return err
	}
	in := &ec2.DeleteKeyPairInput{
		KeyName:   kp.KeyName,
		KeyPairId: kp.KeyPairId,
	}
	_, err = c.DeleteKeyPair(in)
	return err
}
