package ec2

import (
	"github.com/Superm4n97/aws-operations-poc/utils"
	"github.com/aws/aws-sdk-go/service/ec2"
)

const (
	vpcCIDR = "10.1.0.0/16"
)

func NewVPC(c *ec2.EC2) (*ec2.Vpc, error) {
	out, err := c.CreateVpc(&ec2.CreateVpcInput{
		CidrBlock: utils.StringP(vpcCIDR),
	})
	return out.Vpc, err
}

func RemoveVPC(c *ec2.EC2, id string) error {
	_, err := c.DeleteVpc(&ec2.DeleteVpcInput{
		VpcId: utils.StringP(id),
	})
	return err
}
