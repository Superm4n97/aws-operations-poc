package ec2

import (
	"errors"
	"fmt"
	"github.com/Superm4n97/aws-operations-poc/aws"
	"github.com/Superm4n97/aws-operations-poc/utils"
	"github.com/aws/aws-sdk-go/service/ec2"
	"time"
)

const (
	vpcCIDR = "10.1.0.0/16"
)

func NewVPC(c *ec2.EC2) (*ec2.Vpc, error) {
	out, err := c.CreateVpc(&ec2.CreateVpcInput{
		CidrBlock: utils.StringP(vpcCIDR),
	})
	if err != nil {
		return nil, err
	}

	return nil, aws.WaitForState(5*time.Second, 1*time.Minute, func() (bool, error) {
		vpc, err := getVPC(c, out.Vpc.VpcId)
		if err != nil {
			return false, err
		}

		fmt.Println(fmt.Sprintf("current status: %s", *vpc.State))
		if *vpc.State == "available" {
			return true, nil
		}
		return false, nil
	})
}

func RemoveVPC(c *ec2.EC2, id string) error {
	_, err := c.DeleteVpc(&ec2.DeleteVpcInput{
		VpcId: utils.StringP(id),
	})
	return err
}

func getVPC(c *ec2.EC2, vpcId *string) (*ec2.Vpc, error) {
	out, err := c.DescribeVpcs(&ec2.DescribeVpcsInput{
		VpcIds: utils.StringPSlice([]string{*vpcId}),
	})
	if err != nil {
		return nil, err
	}

	for _, vpc := range out.Vpcs {
		if *vpc.VpcId == *vpcId {
			return vpc, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("no vpc found with id %s", *vpcId))
}
