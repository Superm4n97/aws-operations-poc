package ec2

import (
	"errors"
	"fmt"
	"github.com/Superm4n97/aws-operations-poc/aws"
	"github.com/Superm4n97/aws-operations-poc/utils"
	ec2 "github.com/aws/aws-sdk-go/service/ec2"
	_ "k8s.io/apimachinery/pkg/util/wait"
	"time"
)

const (
	//t2.xlarge is a 4 vCPU and 16 GiB Memory
	instanceType             = "t2.xlarge"
	instanceStatusRunning    = "running"
	instanceStatusTerminated = "terminated"

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

func network(subnetID *string) []*ec2.InstanceNetworkInterfaceSpecification {
	var ret []*ec2.InstanceNetworkInterfaceSpecification
	ret = append(ret, &ec2.InstanceNetworkInterfaceSpecification{
		AssociatePublicIpAddress: utils.BoolP(true),
		DeleteOnTermination:      utils.BoolP(true),
		DeviceIndex:              utils.I64P(0),
		SubnetId:                 subnetID,
	})
	return ret
}

func instanceSpecification(keypairName, subnetId *string) *ec2.RunInstancesInput {
	return &ec2.RunInstancesInput{
		InstanceType:      utils.StringP(instanceType),
		ImageId:           utils.StringP(imageID[aws.Region]),
		KeyName:           keypairName,
		NetworkInterfaces: network(subnetId),
		MinCount:          utils.I64P(minCount),
		MaxCount:          utils.I64P(maxCount),
		UserData:          nil,
	}
}

func waitForInstanceStatus(c *ec2.EC2, instanceID, state *string) error {
	var interval time.Duration = 5 //interval in second
	retry := 15
	for i := 1; i <= retry; i++ {
		ins, err := GetInstance(c, instanceID)
		if err != nil {
			return err
		}
		fmt.Println(fmt.Sprintf("retry %d: instance condition %s", i, *ins.State.Name))
		if *ins.State.Name == *state {
			return nil
		}
		time.Sleep(interval * time.Second)
	}
	return errors.New("request timeout")
}

func NewInstance(c *ec2.EC2, keypairName, subnetId *string) (*ec2.Instance, error) {
	input := instanceSpecification(keypairName, subnetId)
	reserv, err := c.RunInstances(input)
	if err != nil {
		return nil, err
	}

	for _, ins := range reserv.Instances {
		err = waitForInstanceStatus(c, ins.InstanceId, utils.StringP(instanceStatusRunning))
		fmt.Println("------------------ instance -----------------------")
		fmt.Println(*ins)
		fmt.Println("---------------------------------------------------------")
		return ins, err
	}
	return nil, errors.New(fmt.Sprintf("failed to create instance"))
}

func GetInstance(c *ec2.EC2, id *string) (*ec2.Instance, error) {
	out, err := c.DescribeInstances(&ec2.DescribeInstancesInput{
		InstanceIds: utils.StringPSlice([]string{*id}),
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

func RemoveInstances(c *ec2.EC2, instanceId *string) error {
	_, err := c.TerminateInstances(&ec2.TerminateInstancesInput{
		InstanceIds: utils.StringPSlice([]string{*instanceId}),
	})
	if err != nil {
		return err
	}
	return waitForInstanceStatus(c, instanceId, utils.StringP(instanceStatusTerminated))
}
