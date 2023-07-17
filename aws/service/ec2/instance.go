package ec2

import (
	"errors"
	"fmt"
	"github.com/Superm4n97/aws-operations-poc/aws"
	"github.com/Superm4n97/aws-operations-poc/utils"
	ec2 "github.com/aws/aws-sdk-go/service/ec2"
)

const (
	//t2.xlarge is a 4 vCPU and 16 GiB Memory
	instanceType = "t2.xlarge"

	maxCount = 1
	minCount = 1
)

// region wise image list for ubuntu 22.04LTS and t2.xlarge (4vCPU+16GB)
var imageID = map[string]string{
	"us-east-1":      "ami-053b0d53c279acc90",
	"us-east-2":      "ami-024e6efaf93d85776",
	"us-west-1":      "ami-0f8e81a3da6e2510a",
	"us-west-2":      "ami-03f65b8614a860c29",
	"ap-south-1":     "ami-0f5ee92e2d63afc18",
	"ap-northeast-3": "ami-0da13880f921c96a5",
	"ap-northeast-2": "ami-0c9c942bd7bf113a2",
	"ap-southeast-1": "ami-0df7a207adb9748c7",
	"ap-southeast-2": "ami-0310483fb2b488153",
	"ap-northeast-1": "ami-0d52744d6551d851e",
	"ca-central-1":   "ami-0ea18256de20ecdfc",
	"eu-central-1":   "ami-04e601abe3e1a910f",
	"eu-west-1":      "ami-01dd271720c1ba44f",
	"eu-west-2":      "ami-0eb260c4d5475b901",
	"eu-west-3":      "ami-05b5a865c3579bbc4",
	"eu-north-1":     "ami-0989fb15ce71ba39e",
	"sa-east-1":      "ami-0af6e9042ea5a4e3e",
}

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
		ImageId:           utils.StringP(imageID[aws.Region]),
		KeyName:           keypairName,
		NetworkInterfaces: network(),
		MinCount:          utils.I64P(minCount),
		MaxCount:          utils.I64P(maxCount),
	}
}

// NewInstance will create
// * KeyPair
// * Instance
func NewInstance(c *ec2.EC2) (*ec2.Reservation, error) {
	kpout, err := newKeyPair(c)
	if err != nil {
		return nil, err
	}
	fmt.Println("keypair created")

	input := instanceSpecification(c, kpout.KeyName)
	if err != nil {
		return nil, err
	}
	reserv, err := c.RunInstances(input)
	if err != nil {
		keypairErr := deleteKeyPair(c, *kpout.KeyName)
		if err != nil {
			err = errors.Join(err, keypairErr)
		} else {
			fmt.Println("keypair deleted")
		}
		return nil, err
	}

	return reserv, nil
}

func getInstace(c *ec2.EC2, id string) (*ec2.Instance, error) {
	out, err := c.DescribeInstances(&ec2.DescribeInstancesInput{
		InstanceIds: utils.StringPSlice([]string{id}),
	})
	if err != nil {
		return nil, err
	}
	for _, res := range out.Reservations {
		for _, ins := range res.Instances {
			return ins, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("no instance found with id: %s", id))
}

func DeleteInstances(c *ec2.EC2, instanceId string) error {
	ins, err := getInstace(c, instanceId)
	if err != nil {
		return err
	}

	kpErr := deleteKeyPair(c, *ins.KeyName)
	if kpErr != nil {
		return kpErr
	} else {
		fmt.Println("keypair deleted")
	}
	_, err = c.TerminateInstances(&ec2.TerminateInstancesInput{
		InstanceIds: utils.StringPSlice([]string{instanceId}),
	})
	return err
}
