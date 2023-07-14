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

func newKeyPair() error {
	ec2.CreateKeyPairInput{
		KeyFormat:         utils.StringP(keyFormat),
		KeyName:           utils.StringP(keyPairName),
		KeyType:           utils.StringP(keyType),
		TagSpecifications: nil,
	}

	return nil
}
