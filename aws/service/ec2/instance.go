package ec2

import (
	"github.com/Superm4n97/aws-operations-poc/aws"
	"github.com/Superm4n97/aws-operations-poc/utils"
	ec2 "github.com/aws/aws-sdk-go/service/ec2"
)

const (
	//t2.xlarge is a 4 vCPU and 16 GiB Memory
	instanceType = "t2.xlarge"

	//ami-024e6efaf93d85776 is Ubuntu 22.04LTS amd64
	awsImageID = "ami-024e6efaf93d85776"
)

func network() []*ec2.InstanceNetworkInterfaceSpecification {
	var ret []*ec2.InstanceNetworkInterfaceSpecification
	ret = append(ret, &ec2.InstanceNetworkInterfaceSpecification{
		AssociatePublicIpAddress: utils.BoolP(true),
		DeleteOnTermination:      utils.BoolP(true),
	})
	return ret
}

func instanceSpecification() *ec2.RunInstancesInput {
	return &ec2.RunInstancesInput{
		InstanceType:      utils.StringP(instanceType),
		ImageId:           utils.StringP(awsImageID),
		KeyName:           nil,
		NetworkInterfaces: network(),
	}
}

func NewInstance() error {
	c, err := aws.EC2Client()
	if err != nil {
		return err
	}

	c.RunInstances()

	return nil
}
