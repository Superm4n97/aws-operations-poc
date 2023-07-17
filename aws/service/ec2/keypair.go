package ec2

import (
	"errors"
	"fmt"
	"github.com/Superm4n97/aws-operations-poc/utils"
	"github.com/aws/aws-sdk-go/service/ec2"
)

const (
	keyPairName = "rasel-keypair"
	keyType     = "rsa"
	keyFormat   = "pem"
)

func getKeyPair(c *ec2.EC2, name string) (*ec2.KeyPairInfo, error) {
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

func newKeyPair(c *ec2.EC2) (*ec2.CreateKeyPairOutput, error) {
	keypairInput := &ec2.CreateKeyPairInput{
		KeyFormat:         utils.StringP(keyFormat),
		KeyName:           utils.StringP(keyPairName),
		KeyType:           utils.StringP(keyType),
		TagSpecifications: nil,
	}
	return c.CreateKeyPair(keypairInput)
}

func deleteKeyPair(c *ec2.EC2, name string) error {
	kp, err := getKeyPair(c, name)
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
