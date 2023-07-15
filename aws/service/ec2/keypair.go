package ec2

import (
	"github.com/Superm4n97/aws-operations-poc/utils"
	"github.com/aws/aws-sdk-go/service/ec2"
)

const (
	keyPairName = "rasel-keypair"
	keyType     = "rsa"
	keyFormat   = "pem"
)

func newKeyPair(client *ec2.EC2) (*ec2.CreateKeyPairOutput, error) {
	keypairInput := &ec2.CreateKeyPairInput{
		KeyFormat:         utils.StringP(keyFormat),
		KeyName:           utils.StringP(keyPairName),
		KeyType:           utils.StringP(keyType),
		TagSpecifications: nil,
	}
	return client.CreateKeyPair(keypairInput)
}

func deleteKeyPair(c *ec2.EC2, name, id *string) error {
	in := &ec2.DeleteKeyPairInput{
		KeyName:   name,
		KeyPairId: id,
	}
	_, err := c.DeleteKeyPair(in)
	return err
}
