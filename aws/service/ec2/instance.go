package ec2

import (
	"errors"
	"github.com/Superm4n97/aws-operations-poc/utils"
	ec2 "github.com/aws/aws-sdk-go/service/ec2"
)

const (
	//t2.xlarge is a 4 vCPU and 16 GiB Memory
	instanceType = "t2.xlarge"

	//ami-024e6efaf93d85776 is Ubuntu 22.04LTS amd64
	awsImageID = "ami-024e6efaf93d85776"

	maxCount = 1
	minCount = 1
)

func network() []*ec2.InstanceNetworkInterfaceSpecification {
	var ret []*ec2.InstanceNetworkInterfaceSpecification
	ret = append(ret, &ec2.InstanceNetworkInterfaceSpecification{
		AssociatePublicIpAddress: utils.BoolP(true),
		DeleteOnTermination:      utils.BoolP(true),
		DeviceIndex:              utils.I64P(0),
	})
	return ret
}

func instanceSpecification(c *ec2.EC2, keypairName *string) *ec2.RunInstancesInput {
	return &ec2.RunInstancesInput{
		InstanceType:      utils.StringP(instanceType),
		ImageId:           utils.StringP(awsImageID),
		KeyName:           keypairName,
		NetworkInterfaces: network(),
		MinCount:          utils.I64P(minCount),
		MaxCount:          utils.I64P(maxCount),
	}
}

func NewInstance(c *ec2.EC2) (*ec2.Reservation, error) {
	kpout, err := newKeyPair(c)
	if err != nil {
		return nil, err
	}

	input := instanceSpecification(c, kpout.KeyName)
	if err != nil {
		return nil, err
	}
	reserv, err := c.RunInstances(input)
	if err != nil {
		keypairErr := deleteKeyPair(c, kpout.KeyName, kpout.KeyPairId)
		if err != nil {
			err = errors.Join(err, keypairErr)
		}
		return nil, err
	}

	return reserv, nil
}
